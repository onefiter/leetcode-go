package fileupload

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// 本文件演示两种最常问的上传：
//   1. SimpleUploadHandler —— 普通小文件上传，关注「限制大小 + 不信任文件名 + 落盘」。
//   2. ChunkUploadHandler —— 大文件分片上传，关注「按序保存分片 + 全部到齐后合并 + 断点续传」。

const (
	// maxUploadSize 普通上传单文件上限，避免内存/磁盘被打爆。
	maxUploadSize = 10 << 20 // 10 MB
	// maxMemory multipart 解析时允许留在内存里的部分，超出部分写临时文件。
	maxMemory = 2 << 20 // 2 MB
)

// randomName 生成随机文件名（16 字节 -> 32 位十六进制），保留原扩展名。
// 关键点：绝不直接信任客户端传来的原始文件名，防止路径穿越和重名覆盖。
func randomName(ext string) (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b) + ext, nil
}

// SimpleUploadHandler 普通文件上传。
// 流程：限制请求体大小 -> 解析 multipart -> 校验大小 -> 生成安全文件名 -> 落盘。
func SimpleUploadHandler(dir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. 用 http.MaxBytesReader 限制整个请求体，防止恶意大请求。
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

		// 2. 解析 multipart 表单。超过 maxMemory 的部分会自动落到临时文件，不会占满内存。
		if err := r.ParseMultipartForm(maxMemory); err != nil {
			http.Error(w, "文件过大或格式错误: "+err.Error(), http.StatusBadRequest)
			return
		}

		// 3. 取出文件。
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "缺少 file 字段", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// 4. 生成随机文件名，只保留服务端自己提取的扩展名。
		name, err := randomName(filepath.Ext(header.Filename))
		if err != nil {
			http.Error(w, "生成文件名失败", http.StatusInternalServerError)
			return
		}

		// 5. 落盘。
		dst, err := os.Create(filepath.Join(dir, name))
		if err != nil {
			http.Error(w, "创建文件失败", http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "保存文件失败", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "上传成功: %s\n", name)
	}
}
