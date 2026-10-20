package fileupload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

// 大文件分片上传（chunked upload）。
// 客户端把大文件切成 N 片，每片携带 index（第几片，从 1 开始）和 total（总片数）上传；
// 服务端把每片先存成临时分片文件，等所有分片到齐后按 index 顺序合并成最终文件。
// 好处：
//   - 单片小、可重传（断点续传），网络抖动不用从头传；
//   - 服务端每片内存占用小；
//   - 天然支持进度、并发分片、秒传（按文件哈希判重）。

// ChunkUploadHandler 处理单个分片，全部到齐后自动合并。
func ChunkUploadHandler(dir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(32 << 20); err != nil { // 单片最多 32MB
			http.Error(w, "解析失败: "+err.Error(), http.StatusBadRequest)
			return
		}

		index, err := strconv.Atoi(r.FormValue("index"))
		if err != nil || index <= 0 {
			http.Error(w, "index 无效", http.StatusBadRequest)
			return
		}
		total, err := strconv.Atoi(r.FormValue("total"))
		if err != nil || total <= 0 || index > total {
			http.Error(w, "total 无效", http.StatusBadRequest)
			return
		}
		filename := r.FormValue("filename")
		if filename == "" {
			http.Error(w, "filename 不能为空", http.StatusBadRequest)
			return
		}

		file, _, err := r.FormFile("chunk")
		if err != nil {
			http.Error(w, "缺少 chunk 字段", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// 1. 保存当前分片（按 index 命名，方便后面排序合并）。
		partPath := chunkPartPath(dir, filename, index)
		if err := writeChunkPart(partPath, file); err != nil {
			http.Error(w, "保存分片失败", http.StatusInternalServerError)
			return
		}

		// 2. 检查是否所有分片都已到齐，是则合并。
		done, err := allChunksPresent(dir, filename, total)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !done {
			fmt.Fprintf(w, "分片 %d/%d 已接收\n", index, total)
			return
		}

		if err := mergeChunks(dir, filename, total); err != nil {
			http.Error(w, "合并失败", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "上传完成，已合并: %s\n", filename)
	}
}

func chunkPartPath(dir, filename string, index int) string {
	return filepath.Join(dir, fmt.Sprintf("%s.part.%d", filename, index))
}

func writeChunkPart(path string, src io.Reader) error {
	// 先写临时文件再 rename，保证单片落盘是原子的，避免写一半的文件被当成有效分片。
	tmp := path + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	if _, err := io.Copy(f, src); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func allChunksPresent(dir, filename string, total int) (bool, error) {
	for i := 1; i <= total; i++ {
		if _, err := os.Stat(chunkPartPath(dir, filename, i)); os.IsNotExist(err) {
			return false, nil
		} else if err != nil {
			return false, err
		}
	}
	return true, nil
}

// mergeChunks 按 index 顺序把分片拼接成最终文件，成功后删除分片。
func mergeChunks(dir, filename string, total int) error {
	dst, err := os.Create(filepath.Join(dir, filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	// 明确按 index 排序，不能依赖目录遍历顺序（os.ReadDir 按文件名字典序，恰好一致，但显式更稳）。
	indexes := make([]int, total)
	for i := range indexes {
		indexes[i] = i + 1
	}
	sort.Ints(indexes)

	for _, i := range indexes {
		partPath := chunkPartPath(dir, filename, i)
		f, err := os.Open(partPath)
		if err != nil {
			return err
		}
		if _, err := io.Copy(dst, f); err != nil {
			f.Close()
			return err
		}
		f.Close()
	}
	// 合并成功后再清理分片；失败则保留分片，便于重传缺失的那几片（断点续传）。
	for _, i := range indexes {
		os.Remove(chunkPartPath(dir, filename, i))
	}
	return nil
}
