# interview

面试常见业务代码手撸合集。全部只用 Go 标准库，每个子包都可以独立 `go test`。

| 目录 | 主题 | 面试考点 |
| --- | --- | --- |
| `fileupload/` | 文件上传、大文件分片上传 | 分片、断点续传、合并、内存限制、命名安全 |
| `channel/` | channel 的用途场景 | worker pool、fan-out/fan-in、pipeline、done 取消 |
| `ratelimit/` | 限流 | 令牌桶、滑动窗口 |
| `lrucache/` | LRU 缓存 | 双向链表 + map，O(1) 读写 |

运行全部测试：

```bash
cd interview && go test ./...
```
