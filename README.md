# leetcode-go
```text
leetcode-go/
├── README.md               # 刷题总记录：刷了哪些题、进度、错题标记
├── go.mod                  # go mod init leetcode-go 初始化
├── go.sum
├── template/               # ✅ 公共模板包（最重要！树状数组、二分、并查集都放这里）
│   ├── bit.go              # 树状数组BIT模板（2916要用）
│   ├── segment_tree.go     # 线段树
│   ├── union_find.go       # 并查集
│   ├── bfs.go              # BFS模板
│   ├── sort.go             # 自定义排序
│   └── test_helper.go      # 辅助函数：打印数组、链表打印
├── array/                  # 专题：数组
│   ├── 0001-two-sum/
│   │   ├── main.go         # 代码主体
│   │   └── main_test.go    # 单元测试用例
│   ├── 0027-remove-element/
│   └── 0088-merge-sorted-array/
├── hash/                   # 哈希表
├── two_pointer/            # 双指针、滑动窗口
├── string/                 # 字符串
├── linkedlist/             # 链表
├── stack_queue/            # 栈、队列、单调栈
├── binary_search/          # 二分查找
├── greedy/                 # 贪心
├── tree/                   # 二叉树
├── dfs_bfs/                # DFS、回溯、岛屿问题
├── dp/                     # 动态规划
├── graph/                  # 图论、拓扑排序
├── heap/                   # 堆
├── trie/                   # 字典树
└── hard/                   # 单独放Hard难题，比如2916
    └── 2916-subarrays-distinct-element-squares-ii/
        ├── main.go
        └── main_test.go


```