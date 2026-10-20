package fileupload

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

// newUploadRequest 构造一个模拟 multipart 上传的 HTTP 请求。
func newUploadRequest(t *testing.T, field, filename string, content []byte, extra map[string]string) *http.Request {
	t.Helper()
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	fw, err := w.CreateFormFile(field, filename)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := fw.Write(content); err != nil {
		t.Fatal(err)
	}
	for k, v := range extra {
		if err := w.WriteField(k, v); err != nil {
			t.Fatal(err)
		}
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/", &buf)
	req.Header.Set("Content-Type", w.FormDataContentType())
	return req
}

func TestSimpleUploadHandler(t *testing.T) {
	dir := t.TempDir()
	h := SimpleUploadHandler(dir)
	req := newUploadRequest(t, "file", "hello.txt", []byte("hello world"), nil)
	rec := httptest.NewRecorder()
	h(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200, body = %s", rec.Code, rec.Body.String())
	}
	// 目录里应恰好多出一个文件，内容正确，且文件名不是客户端原始名。
	entries, _ := os.ReadDir(dir)
	if len(entries) != 1 {
		t.Fatalf("目录文件数 = %d, want 1", len(entries))
	}
	if entries[0].Name() == "hello.txt" {
		t.Errorf("直接使用了客户端文件名，存在安全风险")
	}
	data, _ := os.ReadFile(filepath.Join(dir, entries[0].Name()))
	if string(data) != "hello world" {
		t.Errorf("文件内容 = %q, want %q", data, "hello world")
	}
}

func TestChunkUploadHandler(t *testing.T) {
	dir := t.TempDir()
	h := ChunkUploadHandler(dir)

	// 把一段内容拆成 3 片，乱序上传，验证「按 index 合并」而非「按到达顺序」。
	chunks := [][]byte{[]byte("ab"), []byte("cde"), []byte("fg")}
	order := []int{2, 1, 3}
	for _, idx := range order {
		req := newUploadRequest(t, "chunk", "chunk", chunks[idx-1], map[string]string{
			"index":    strconv.Itoa(idx),
			"total":    "3",
			"filename": "big.bin",
		})
		rec := httptest.NewRecorder()
		h(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("分片 %d: status = %d", idx, rec.Code)
		}
	}

	data, err := os.ReadFile(filepath.Join(dir, "big.bin"))
	if err != nil {
		t.Fatalf("合并文件不存在: %v", err)
	}
	if string(data) != "abcdefg" {
		t.Errorf("合并结果 = %q, want %q", data, "abcdefg")
	}

	// 分片临时文件应已清理。
	if matches, _ := filepath.Glob(filepath.Join(dir, "*.part.*")); len(matches) != 0 {
		t.Errorf("分片临时文件未清理: %v", matches)
	}
}
