# LeetCode 300 题 · 一周面试冲刺计划

> **适用人群**：已有 200+ 刷题基础，目标大厂 / 外企算法面试，一周集中冲刺。
> **核心策略**：高频考点快速回归（70%）+ 薄弱点定点攻坚（20%）+ Hard 题战术突破（10%）
> **题目来源**：LeetCode 热题 100 / 剑指 Offer / CodeTop 面试高频榜
> **完成状态**：300 题已全部完成，题单「代码」列直接链接到对应解法。代码按专题目录存放（不按天），同一题在计划中出现多次时指向同一份代码；目录说明与整体解题思路见 [README](README.md)。

---

## 一、计划总览

- **总题量**：300 题（🟢 Easy 66 题 / 🟡 Medium 196 题 / 🔴 Hard 38 题）
- **周期**：7 天（Day 1–Day 7）
- **每日节奏**：上午 9:00–12:00（攻坚）｜下午 14:00–18:00（新知）｜晚上 19:30–22:30（复习+复盘）
- **每日题量**：45 / 45 / 50 / 49 / 45 / 34 / 32 = **300 题**

### 7 天日程总表

| 天数 | 主题 | 题量 | 上午 | 下午 | 晚上 |
| --- | --- | --- | --- | --- | --- |
| **Day 1** | 数组 · 哈希 · 字符串 | 45 | 数组/前缀和/双指针 (20) | 哈希表 (15) | 字符串 (10) + 复盘 |
| **Day 2** | 双指针 · 滑窗 · 二分 | 45 | 双指针 (20) | 滑动窗口 (15) | 二分查找 (10) + 复盘 |
| **Day 3** | 链表 · 栈 · 队列 | 50 | 链表 (28) | 单调栈 (14) | 队列/设计 (8) + 复盘 |
| **Day 4** | 二叉树全专题 | 49 | 遍历 (20) | BST (15) | 路径/树形 (14) + 复盘 |
| **Day 5** | 回溯 · 图 · BFS | 45 | 回溯 (15) | 图论 (15) | BFS 网格 (15) + 复盘 |
| **Day 6** | 动态规划(上) · 贪心 | 34 | 线性 DP (12) | 背包 (8) | 贪心 (14) + 复盘 |
| **Day 7** | 动态规划(下) · 模考 | 32 | 进阶 DP (14) | 全真模考 (5) | Hard+设计 (13) + 总复习 |

### 专题权重分布

| 专题分类 | 题量 | 优先级 | 面试频率 | 分布天数 |
| --- | --- | --- | --- | --- |
| 数组 & 哈希表 | 35 | ★★★ | 极高 | Day 1 |
| 字符串 | 10 | ★★★ | 极高 | Day 1 |
| 双指针 & 滑动窗口 | 35 | ★★★ | 极高 | Day 2 |
| 二分查找 | 10 | ★★★ | 高 | Day 2 |
| 链表 | 28 | ★★★ | 高 | Day 3 |
| 栈 & 队列 | 22 | ★★★ | 高 | Day 3 |
| 二叉树 & BST | 49 | ★★★ | 极高 | Day 4 |
| 回溯 & 图 & BFS | 45 | ★★☆ | 中高 | Day 5 |
| 动态规划 | 34 | ★★★ | 极高 | Day 6-7 |
| 贪心 | 14 | ★★☆ | 中 | Day 6 |
| 模考 & Hard & 设计 | 18 | ★★☆ | 中高 | Day 7 |

---

## 二、每日题单

### 📅 Day 1 · 数组 / 哈希 / 字符串（45 题）

#### 数组 · 前缀和 · 排序双指针（20 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 1 | [704. 二分查找](https://leetcode.cn/problems/binary-search/) | 🟢 Easy | ✅ [binary_search/0704](binary_search/0704-binary-search/main.go) |
| 2 | [35. 搜索插入位置](https://leetcode.cn/problems/search-insert-position/) | 🟢 Easy | ✅ [binary_search/0035](binary_search/0035-search-insert-position/main.go) |
| 3 | [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/) | 🟡 Medium | ✅ [binary_search/0034](binary_search/0034-find-first-and-last-position/main.go) |
| 4 | [33. 搜索旋转排序数组](https://leetcode.cn/problems/search-in-rotated-sorted-array/) | 🟡 Medium | ✅ [binary_search/0033](binary_search/0033-search-in-rotated-sorted-array/main.go) |
| 5 | [81. 搜索旋转排序数组 II](https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/) | 🟡 Medium | ✅ [binary_search/0081](binary_search/0081-search-in-rotated-sorted-array-ii/main.go) |
| 6 | [153. 寻找旋转排序数组中的最小值](https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/) | 🟡 Medium | ✅ [binary_search/0153](binary_search/0153-find-minimum-in-rotated-sorted-array/main.go) |
| 7 | [1. 两数之和](https://leetcode.cn/problems/two-sum/) | 🟢 Easy | ✅ [array/0001](array/0001-two-sum/main.go) |
| 8 | [167. 两数之和 II - 输入有序数组](https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/) | 🟡 Medium | ✅ [two_pointer/0167](two_pointer/0167-two-sum-ii-input-array-is-sorted/main.go) |
| 9 | [15. 三数之和](https://leetcode.cn/problems/3sum/) | 🟡 Medium | ✅ [two_pointer/0015](two_pointer/0015-3sum/main.go) |
| 10 | [16. 最接近的三数之和](https://leetcode.cn/problems/3sum-closest/) | 🟡 Medium | ✅ [two_pointer/0016](two_pointer/0016-3sum-closest/main.go) |
| 11 | [18. 四数之和](https://leetcode.cn/problems/4sum/) | 🟡 Medium | ✅ [two_pointer/0018](two_pointer/0018-4sum/main.go) |
| 12 | [11. 盛最多水的容器](https://leetcode.cn/problems/container-with-most-water/) | 🟡 Medium | ✅ [two_pointer/0011](two_pointer/0011-container-with-most-water/main.go) |
| 13 | [26. 删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/) | 🟢 Easy | ✅ [array/0026](array/0026-remove-duplicates-from-sorted-array/main.go) |
| 14 | [27. 移除元素](https://leetcode.cn/problems/remove-element/) | 🟢 Easy | ✅ [array/0027](array/0027-remove-element/main.go) |
| 15 | [283. 移动零](https://leetcode.cn/problems/move-zeroes/) | 🟢 Easy | ✅ [array/0283](array/0283-move-zeroes/main.go) |
| 16 | [56. 合并区间](https://leetcode.cn/problems/merge-intervals/) | 🟡 Medium | ✅ [array/0056](array/0056-merge-intervals/main.go) |
| 17 | [57. 插入区间](https://leetcode.cn/problems/insert-interval/) | 🟡 Medium | ✅ [array/0057](array/0057-insert-interval/main.go) |
| 18 | [238. 除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/) | 🟡 Medium | ✅ [array/0238](array/0238-product-of-array-except-self/main.go) |
| 19 | [303. 区域和检索 - 数组不可变](https://leetcode.cn/problems/range-sum-query-immutable/) | 🟢 Easy | ✅ [array/0303](array/0303-range-sum-query-immutable/main.go) |
| 20 | [304. 二维区域和检索 - 矩阵不可变](https://leetcode.cn/problems/range-sum-query-2d-immutable/) | 🟡 Medium | ✅ [array/0304](array/0304-range-sum-query-2d-immutable/main.go) |

#### 哈希表 · 前缀和思想（15 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 21 | [49. 字母异位词分组](https://leetcode.cn/problems/group-anagrams/) | 🟡 Medium | ✅ [hash/0049](hash/0049-group-anagrams/main.go) |
| 22 | [128. 最长连续序列](https://leetcode.cn/problems/longest-consecutive-sequence/) | 🟡 Medium | ✅ [hash/0128](hash/0128-longest-consecutive-sequence/main.go) |
| 23 | [217. 存在重复元素](https://leetcode.cn/problems/contains-duplicate/) | 🟢 Easy | ✅ [hash/0217](hash/0217-contains-duplicate/main.go) |
| 24 | [219. 存在重复元素 II](https://leetcode.cn/problems/contains-duplicate-ii/) | 🟢 Easy | ✅ [hash/0219](hash/0219-contains-duplicate-ii/main.go) |
| 25 | [220. 存在重复元素 III](https://leetcode.cn/problems/contains-duplicate-iii/) | 🔴 Hard | ✅ [hash/0220](hash/0220-contains-duplicate-iii/main.go) |
| 26 | [242. 有效的字母异位词](https://leetcode.cn/problems/valid-anagram/) | 🟢 Easy | ✅ [hash/0242](hash/0242-valid-anagram/main.go) |
| 27 | [349. 两个数组的交集](https://leetcode.cn/problems/intersection-of-two-arrays/) | 🟢 Easy | ✅ [hash/0349](hash/0349-intersection-of-two-arrays/main.go) |
| 28 | [350. 两个数组的交集 II](https://leetcode.cn/problems/intersection-of-two-arrays-ii/) | 🟢 Easy | ✅ [hash/0350](hash/0350-intersection-of-two-arrays-ii/main.go) |
| 29 | [454. 四数相加 II](https://leetcode.cn/problems/4sum-ii/) | 🟡 Medium | ✅ [hash/0454](hash/0454-4sum-ii/main.go) |
| 30 | [560. 和为 K 的子数组](https://leetcode.cn/problems/subarray-sum-equals-k/) | 🟡 Medium | ✅ [hash/0560](hash/0560-subarray-sum-equals-k/main.go) |
| 31 | [974. 和可被 K 整除的子数组](https://leetcode.cn/problems/subarray-sums-divisible-by-k/) | 🟡 Medium | ✅ [hash/0974](hash/0974-subarray-sums-divisible-by-k/main.go) |
| 32 | [76. 最小覆盖子串](https://leetcode.cn/problems/minimum-window-substring/) | 🔴 Hard | ✅ [two_pointer/0076](two_pointer/0076-minimum-window-substring/main.go) |
| 33 | [438. 找到字符串中所有字母异位词](https://leetcode.cn/problems/find-all-anagrams-in-a-string/) | 🟡 Medium | ✅ [two_pointer/0438](two_pointer/0438-find-all-anagrams-in-a-string/main.go) |
| 34 | [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/) | 🟡 Medium | ✅ [two_pointer/0003](two_pointer/0003-longest-substring-without-repeating-characters/main.go) |
| 35 | [387. 字符串中的第一个唯一字符](https://leetcode.cn/problems/first-unique-character-in-a-string/) | 🟢 Easy | ✅ [hash/0387](hash/0387-first-unique-character-in-a-string/main.go) |

#### 字符串 · 模拟（10 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 36 | [5. 最长回文子串](https://leetcode.cn/problems/longest-palindromic-substring/) | 🟡 Medium | ✅ [string/0005](string/0005-longest-palindromic-substring/main.go) |
| 37 | [125. 验证回文串](https://leetcode.cn/problems/valid-palindrome/) | 🟢 Easy | ✅ [string/0125](string/0125-valid-palindrome/main.go) |
| 38 | [151. 反转字符串中的单词](https://leetcode.cn/problems/reverse-words-in-a-string/) | 🟡 Medium | ✅ [string/0151](string/0151-reverse-words-in-a-string/main.go) |
| 39 | [344. 反转字符串](https://leetcode.cn/problems/reverse-string/) | 🟢 Easy | ✅ [string/0344](string/0344-reverse-string/main.go) |
| 40 | [14. 最长公共前缀](https://leetcode.cn/problems/longest-common-prefix/) | 🟢 Easy | ✅ [string/0014](string/0014-longest-common-prefix/main.go) |
| 41 | [43. 字符串相乘](https://leetcode.cn/problems/multiply-strings/) | 🟡 Medium | ✅ [string/0043](string/0043-multiply-strings/main.go) |
| 42 | [8. 字符串转换整数 (atoi)](https://leetcode.cn/problems/string-to-integer-atoi/) | 🟡 Medium | ✅ [string/0008](string/0008-string-to-integer-atoi/main.go) |
| 43 | [165. 比较版本号](https://leetcode.cn/problems/compare-version-numbers/) | 🟡 Medium | ✅ [string/0165](string/0165-compare-version-numbers/main.go) |
| 44 | [6. Z 字形变换](https://leetcode.cn/problems/zigzag-conversion/) | 🟡 Medium | ✅ [string/0006](string/0006-zigzag-conversion/main.go) |
| 45 | [28. 找出字符串中第一个匹配项的下标](https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/) | 🟢 Easy | ✅ [string/0028](string/0028-find-the-index-of-the-first-occurrence-in-a-string/main.go) |

### 📅 Day 2 · 双指针 / 滑动窗口 / 二分（45 题）

#### 双指针（对撞 / 快慢 / 分离）（20 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 46 | [19. 删除链表的倒数第 N 个结点](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/) | 🟡 Medium | ✅ [linkedlist/0019](linkedlist/0019-remove-nth-node-from-end-of-list/main.go) |
| 47 | [21. 合并两个有序链表](https://leetcode.cn/problems/merge-two-sorted-lists/) | 🟢 Easy | ✅ [linkedlist/0021](linkedlist/0021-merge-two-sorted-lists/main.go) |
| 48 | [42. 接雨水](https://leetcode.cn/problems/trapping-rain-water/) | 🔴 Hard | ✅ [two_pointer/0042](two_pointer/0042-trapping-rain-water/main.go) |
| 49 | [75. 颜色分类](https://leetcode.cn/problems/sort-colors/) | 🟡 Medium | ✅ [two_pointer/0075](two_pointer/0075-sort-colors/main.go) |
| 50 | [80. 删除有序数组中的重复项 II](https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/) | 🟡 Medium | ✅ [two_pointer/0080](two_pointer/0080-remove-duplicates-from-sorted-array-ii/main.go) |
| 51 | [88. 合并两个有序数组](https://leetcode.cn/problems/merge-sorted-array/) | 🟢 Easy | ✅ [array/0088](array/0088-merge-sorted-array/main.go) |
| 52 | [141. 环形链表](https://leetcode.cn/problems/linked-list-cycle/) | 🟢 Easy | ✅ [linkedlist/0141](linkedlist/0141-linked-list-cycle/main.go) |
| 53 | [142. 环形链表 II](https://leetcode.cn/problems/linked-list-cycle-ii/) | 🟡 Medium | ✅ [linkedlist/0142](linkedlist/0142-linked-list-cycle-ii/main.go) |
| 54 | [143. 重排链表](https://leetcode.cn/problems/reorder-list/) | 🟡 Medium | ✅ [linkedlist/0143](linkedlist/0143-reorder-list/main.go) |
| 55 | [234. 回文链表](https://leetcode.cn/problems/palindrome-linked-list/) | 🟢 Easy | ✅ [linkedlist/0234](linkedlist/0234-palindrome-linked-list/main.go) |
| 56 | [287. 寻找重复数](https://leetcode.cn/problems/find-the-duplicate-number/) | 🟡 Medium | ✅ [two_pointer/0287](two_pointer/0287-find-the-duplicate-number/main.go) |
| 57 | [345. 反转字符串中的元音字母](https://leetcode.cn/problems/reverse-vowels-of-a-string/) | 🟢 Easy | ✅ [two_pointer/0345](two_pointer/0345-reverse-vowels-of-a-string/main.go) |
| 58 | [392. 判断子序列](https://leetcode.cn/problems/is-subsequence/) | 🟢 Easy | ✅ [two_pointer/0392](two_pointer/0392-is-subsequence/main.go) |
| 59 | [876. 链表的中间结点](https://leetcode.cn/problems/middle-of-the-linked-list/) | 🟢 Easy | ✅ [linkedlist/0876](linkedlist/0876-middle-of-the-linked-list/main.go) |
| 60 | [977. 有序数组的平方](https://leetcode.cn/problems/squares-of-a-sorted-array/) | 🟢 Easy | ✅ [two_pointer/0977](two_pointer/0977-squares-of-a-sorted-array/main.go) |
| 61 | [209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/) | 🟡 Medium | ✅ [two_pointer/0209](two_pointer/0209-minimum-size-subarray-sum/main.go) |
| 62 | [845. 数组中的最长山脉](https://leetcode.cn/problems/longest-mountain-in-array/) | 🟡 Medium | ✅ [two_pointer/0845](two_pointer/0845-longest-mountain-in-array/main.go) |
| 63 | [986. 区间列表的交集](https://leetcode.cn/problems/interval-list-intersections/) | 🟡 Medium | ✅ [two_pointer/0986](two_pointer/0986-interval-list-intersections/main.go) |
| 64 | [1021. 删除最外层的括号](https://leetcode.cn/problems/remove-outermost-parentheses/) | 🟢 Easy | ✅ [two_pointer/1021](two_pointer/1021-remove-outermost-parentheses/main.go) |
| 65 | [1099. 小于 K 的两数之和](https://leetcode.cn/problems/two-sum-less-than-k/) | 🟢 Easy | ✅ [two_pointer/1099](two_pointer/1099-two-sum-less-than-k/main.go) |

#### 滑动窗口（定长 / 不定长）（15 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 66 | [30. 串联所有单词的子串](https://leetcode.cn/problems/substring-with-concatenation-of-all-words/) | 🔴 Hard | ✅ [two_pointer/0030](two_pointer/0030-substring-with-concatenation-of-all-words/main.go) |
| 67 | [159. 至多包含两个不同字符的最长子串](https://leetcode.cn/problems/longest-substring-with-at-most-two-distinct-characters/) | 🟡 Medium | ✅ [two_pointer/0159](two_pointer/0159-longest-substring-with-at-most-two-distinct-characters/main.go) |
| 68 | [187. 重复的 DNA 序列](https://leetcode.cn/problems/repeated-dna-sequences/) | 🟡 Medium | ✅ [two_pointer/0187](two_pointer/0187-repeated-dna-sequences/main.go) |
| 69 | [239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/) | 🔴 Hard | ✅ [two_pointer/0239](two_pointer/0239-sliding-window-maximum/main.go) |
| 70 | [340. 至多包含 K 个不同字符的最长子串](https://leetcode.cn/problems/longest-substring-with-at-most-k-distinct-characters/) | 🟡 Medium | ✅ [two_pointer/0340](two_pointer/0340-longest-substring-with-at-most-k-distinct-characters/main.go) |
| 71 | [424. 替换后的最长重复字符](https://leetcode.cn/problems/longest-repeating-character-replacement/) | 🟡 Medium | ✅ [two_pointer/0424](two_pointer/0424-longest-repeating-character-replacement/main.go) |
| 72 | [487. 最大连续 1 的个数 II](https://leetcode.cn/problems/max-consecutive-ones-ii/) | 🟡 Medium | ✅ [two_pointer/0487](two_pointer/0487-max-consecutive-ones-ii/main.go) |
| 73 | [567. 字符串的排列](https://leetcode.cn/problems/permutation-in-string/) | 🟡 Medium | ✅ [two_pointer/0567](two_pointer/0567-permutation-in-string/main.go) |
| 74 | [643. 子数组最大平均数 I](https://leetcode.cn/problems/maximum-average-subarray-i/) | 🟢 Easy | ✅ [two_pointer/0643](two_pointer/0643-maximum-average-subarray-i/main.go) |
| 75 | [713. 乘积小于 K 的子数组](https://leetcode.cn/problems/subarray-product-less-than-k/) | 🟡 Medium | ✅ [two_pointer/0713](two_pointer/0713-subarray-product-less-than-k/main.go) |
| 76 | [904. 水果成篮](https://leetcode.cn/problems/fruit-into-baskets/) | 🟡 Medium | ✅ [two_pointer/0904](two_pointer/0904-fruit-into-baskets/main.go) |
| 77 | [1004. 最大连续 1 的个数 III](https://leetcode.cn/problems/max-consecutive-ones-iii/) | 🟡 Medium | ✅ [two_pointer/1004](two_pointer/1004-max-consecutive-ones-iii/main.go) |
| 78 | [1052. 爱生气的书店老板](https://leetcode.cn/problems/grumpy-bookstore-owner/) | 🟡 Medium | ✅ [two_pointer/1052](two_pointer/1052-grumpy-bookstore-owner/main.go) |
| 79 | [1151. 最少交换次数来组合所有的 1](https://leetcode.cn/problems/minimum-swaps-to-group-all-1s-together/) | 🟡 Medium | ✅ [two_pointer/1151](two_pointer/1151-minimum-swaps-to-group-all-1s-together/main.go) |
| 80 | [1493. 删掉一个元素以后全为 1 的最长子数组](https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/) | 🟡 Medium | ✅ [two_pointer/1493](two_pointer/1493-longest-subarray-of-1s-after-deleting-one-element/main.go) |

#### 二分查找（基础 / 边界 / 二分答案）（10 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 81 | [4. 寻找两个正序数组的中位数](https://leetcode.cn/problems/median-of-two-sorted-arrays/) | 🔴 Hard | ✅ [binary_search/0004](binary_search/0004-median-of-two-sorted-arrays/main.go) |
| 82 | [69. x 的平方根](https://leetcode.cn/problems/sqrtx/) | 🟢 Easy | ✅ [binary_search/0069](binary_search/0069-sqrtx/main.go) |
| 83 | [74. 搜索二维矩阵](https://leetcode.cn/problems/search-a-2d-matrix/) | 🟡 Medium | ✅ [binary_search/0074](binary_search/0074-search-a-2d-matrix/main.go) |
| 84 | [162. 寻找峰值](https://leetcode.cn/problems/find-peak-element/) | 🟡 Medium | ✅ [binary_search/0162](binary_search/0162-find-peak-element/main.go) |
| 85 | [278. 第一个错误的版本](https://leetcode.cn/problems/first-bad-version/) | 🟢 Easy | ✅ [binary_search/0278](binary_search/0278-first-bad-version/main.go) |
| 86 | [300. 最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence/) | 🟡 Medium | ✅ [binary_search/0300](binary_search/0300-longest-increasing-subsequence/main.go) |
| 87 | [374. 猜数字大小](https://leetcode.cn/problems/guess-number-higher-or-lower/) | 🟢 Easy | ✅ [binary_search/0374](binary_search/0374-guess-number-higher-or-lower/main.go) |
| 88 | [540. 有序数组中的单一元素](https://leetcode.cn/problems/single-element-in-a-sorted-array/) | 🟡 Medium | ✅ [binary_search/0540](binary_search/0540-single-element-in-a-sorted-array/main.go) |
| 89 | [658. 找到 K 个最接近的元素](https://leetcode.cn/problems/find-k-closest-elements/) | 🟡 Medium | ✅ [binary_search/0658](binary_search/0658-find-k-closest-elements/main.go) |
| 90 | [875. 爱吃香蕉的珂珂](https://leetcode.cn/problems/koko-eating-bananas/) | 🟡 Medium | ✅ [binary_search/0875](binary_search/0875-koko-eating-bananas/main.go) |

### 📅 Day 3 · 链表 / 栈 / 队列（50 题）

#### 链表（反转 / 合并 / 环 / 设计）（25 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 91 | [2. 两数相加](https://leetcode.cn/problems/add-two-numbers/) | 🟡 Medium | ✅ [linkedlist/0002](linkedlist/0002-add-two-numbers/main.go) |
| 92 | [23. 合并 K 个升序链表](https://leetcode.cn/problems/merge-k-sorted-lists/) | 🔴 Hard | ✅ [linkedlist/0023](linkedlist/0023-merge-k-sorted-lists/main.go) |
| 93 | [24. 两两交换链表中的节点](https://leetcode.cn/problems/swap-nodes-in-pairs/) | 🟡 Medium | ✅ [linkedlist/0024](linkedlist/0024-swap-nodes-in-pairs/main.go) |
| 94 | [25. K 个一组翻转链表](https://leetcode.cn/problems/reverse-nodes-in-k-group/) | 🔴 Hard | ✅ [linkedlist/0025](linkedlist/0025-reverse-nodes-in-k-group/main.go) |
| 95 | [61. 旋转链表](https://leetcode.cn/problems/rotate-list/) | 🟡 Medium | ✅ [linkedlist/0061](linkedlist/0061-rotate-list/main.go) |
| 96 | [82. 删除排序链表中的重复元素 II](https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/) | 🟡 Medium | ✅ [linkedlist/0082](linkedlist/0082-remove-duplicates-from-sorted-list-ii/main.go) |
| 97 | [83. 删除排序链表中的重复元素](https://leetcode.cn/problems/remove-duplicates-from-sorted-list/) | 🟢 Easy | ✅ [linkedlist/0083](linkedlist/0083-remove-duplicates-from-sorted-list/main.go) |
| 98 | [86. 分隔链表](https://leetcode.cn/problems/partition-list/) | 🟡 Medium | ✅ [linkedlist/0086](linkedlist/0086-partition-list/main.go) |
| 99 | [92. 反转链表 II](https://leetcode.cn/problems/reverse-linked-list-ii/) | 🟡 Medium | ✅ [linkedlist/0092](linkedlist/0092-reverse-linked-list-ii/main.go) |
| 100 | [138. 随机链表的复制](https://leetcode.cn/problems/copy-list-with-random-pointer/) | 🟡 Medium | ✅ [linkedlist/0138](linkedlist/0138-copy-list-with-random-pointer/main.go) |
| 101 | [146. LRU 缓存](https://leetcode.cn/problems/lru-cache/) | 🟡 Medium | ✅ [linkedlist/0146](linkedlist/0146-lru-cache/main.go) |
| 102 | [148. 排序链表](https://leetcode.cn/problems/sort-list/) | 🟡 Medium | ✅ [linkedlist/0148](linkedlist/0148-sort-list/main.go) |
| 103 | [160. 相交链表](https://leetcode.cn/problems/intersection-of-two-linked-lists/) | 🟢 Easy | ✅ [linkedlist/0160](linkedlist/0160-intersection-of-two-linked-lists/main.go) |
| 104 | [203. 移除链表元素](https://leetcode.cn/problems/remove-linked-list-elements/) | 🟢 Easy | ✅ [linkedlist/0203](linkedlist/0203-remove-linked-list-elements/main.go) |
| 105 | [206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/) | 🟢 Easy | ✅ [linkedlist/0206](linkedlist/0206-reverse-linked-list/main.go) |
| 106 | [328. 奇偶链表](https://leetcode.cn/problems/odd-even-linked-list/) | 🟡 Medium | ✅ [linkedlist/0328](linkedlist/0328-odd-even-linked-list/main.go) |
| 107 | [445. 两数相加 II](https://leetcode.cn/problems/add-two-numbers-ii/) | 🟡 Medium | ✅ [linkedlist/0445](linkedlist/0445-add-two-numbers-ii/main.go) |
| 108 | [460. LFU 缓存](https://leetcode.cn/problems/lfu-cache/) | 🔴 Hard | ✅ [linkedlist/0460](linkedlist/0460-lfu-cache/main.go) |
| 109 | [707. 设计链表](https://leetcode.cn/problems/design-linked-list/) | 🟡 Medium | ✅ [linkedlist/0707](linkedlist/0707-design-linked-list/main.go) |
| 110 | [725. 分隔链表 II](https://leetcode.cn/problems/split-linked-list-in-parts/) | 🟡 Medium | ✅ [linkedlist/0725](linkedlist/0725-split-linked-list-in-parts/main.go) |
| 111 | [1382. 将二叉搜索树变平衡](https://leetcode.cn/problems/balance-a-binary-search-tree/) | 🟡 Medium | ✅ [tree/1382](tree/1382-balance-a-binary-search-tree/main.go) |
| 112 | [1669. 合并两个链表](https://leetcode.cn/problems/merge-in-between-linked-lists/) | 🟡 Medium | ✅ [linkedlist/1669](linkedlist/1669-merge-in-between-linked-lists/main.go) |
| 113 | [2095. 删除链表的中间节点](https://leetcode.cn/problems/delete-the-middle-node-of-a-linked-list/) | 🟡 Medium | ✅ [linkedlist/2095](linkedlist/2095-delete-the-middle-node-of-a-linked-list/main.go) |
| 114 | [2130. 链表最大孪生和](https://leetcode.cn/problems/maximum-twin-sum-of-a-linked-list/) | 🟡 Medium | ✅ [linkedlist/2130](linkedlist/2130-maximum-twin-sum-of-a-linked-list/main.go) |
| 115 | [2487. 从链表中移除节点](https://leetcode.cn/problems/remove-nodes-from-linked-list/) | 🟡 Medium | ✅ [linkedlist/2487](linkedlist/2487-remove-nodes-from-linked-list/main.go) |

#### 链表 · 补充练习（3 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 116 | [1290. 二进制链表转整数](https://leetcode.cn/problems/convert-binary-number-in-a-linked-list-to-integer/) | 🟢 Easy | ✅ [linkedlist/1290](linkedlist/1290-convert-binary-number-in-a-linked-list-to-integer/main.go) |
| 117 | [1171. 从链表中删去总和值为零的连续节点](https://leetcode.cn/problems/remove-zero-sum-consecutive-nodes-from-linked-list/) | 🟡 Medium | ✅ [linkedlist/1171](linkedlist/1171-remove-zero-sum-consecutive-nodes-from-linked-list/main.go) |
| 118 | [2058. 找出临界点之间的最小和最大距离](https://leetcode.cn/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/) | 🟡 Medium | ✅ [linkedlist/2058](linkedlist/2058-find-the-minimum-and-maximum-number-of-nodes-between-critical-points/main.go) |

#### 栈 & 单调栈（14 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 119 | [20. 有效的括号](https://leetcode.cn/problems/valid-parentheses/) | 🟢 Easy | ✅ [stack_queue/0020](stack_queue/0020-valid-parentheses/main.go) |
| 120 | [32. 最长有效括号](https://leetcode.cn/problems/longest-valid-parentheses/) | 🔴 Hard | ✅ [stack_queue/0032](stack_queue/0032-longest-valid-parentheses/main.go) |
| 121 | [84. 柱状图中最大的矩形](https://leetcode.cn/problems/largest-rectangle-in-histogram/) | 🔴 Hard | ✅ [stack_queue/0084](stack_queue/0084-largest-rectangle-in-histogram/main.go) |
| 122 | [150. 逆波兰表达式求值](https://leetcode.cn/problems/evaluate-reverse-polish-notation/) | 🟡 Medium | ✅ [stack_queue/0150](stack_queue/0150-evaluate-reverse-polish-notation/main.go) |
| 123 | [155. 最小栈](https://leetcode.cn/problems/min-stack/) | 🟡 Medium | ✅ [stack_queue/0155](stack_queue/0155-min-stack/main.go) |
| 124 | [224. 基本计算器](https://leetcode.cn/problems/basic-calculator/) | 🔴 Hard | ✅ [stack_queue/0224](stack_queue/0224-basic-calculator/main.go) |
| 125 | [225. 用队列实现栈](https://leetcode.cn/problems/implement-stack-using-queues/) | 🟢 Easy | ✅ [stack_queue/0225](stack_queue/0225-implement-stack-using-queues/main.go) |
| 126 | [232. 用栈实现队列](https://leetcode.cn/problems/implement-queue-using-stacks/) | 🟢 Easy | ✅ [stack_queue/0232](stack_queue/0232-implement-queue-using-stacks/main.go) |
| 127 | [316. 去除重复字母](https://leetcode.cn/problems/remove-duplicate-letters/) | 🟡 Medium | ✅ [stack_queue/0316](stack_queue/0316-remove-duplicate-letters/main.go) |
| 128 | [394. 字符串解码](https://leetcode.cn/problems/decode-string/) | 🟡 Medium | ✅ [stack_queue/0394](stack_queue/0394-decode-string/main.go) |
| 129 | [496. 下一个更大元素 I](https://leetcode.cn/problems/next-greater-element-i/) | 🟢 Easy | ✅ [stack_queue/0496](stack_queue/0496-next-greater-element-i/main.go) |
| 130 | [503. 下一个更大元素 II](https://leetcode.cn/problems/next-greater-element-ii/) | 🟡 Medium | ✅ [stack_queue/0503](stack_queue/0503-next-greater-element-ii/main.go) |
| 131 | [739. 每日温度](https://leetcode.cn/problems/daily-temperatures/) | 🟡 Medium | ✅ [stack_queue/0739](stack_queue/0739-daily-temperatures/main.go) |
| 132 | [907. 子数组的最小值之和](https://leetcode.cn/problems/sum-of-subarray-minimums/) | 🟡 Medium | ✅ [stack_queue/0907](stack_queue/0907-sum-of-subarray-minimums/main.go) |

#### 队列 & 设计题（8 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 133 | [622. 设计循环队列](https://leetcode.cn/problems/design-circular-queue/) | 🟡 Medium | ✅ [stack_queue/0622](stack_queue/0622-design-circular-queue/main.go) |
| 134 | [641. 设计循环双端队列](https://leetcode.cn/problems/design-circular-deque/) | 🟡 Medium | ✅ [stack_queue/0641](stack_queue/0641-design-circular-deque/main.go) |
| 135 | [862. 和至少为 K 的最短子数组](https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/) | 🔴 Hard | ✅ [stack_queue/0862](stack_queue/0862-shortest-subarray-with-sum-at-least-k/main.go) |
| 136 | [933. 最近的请求次数](https://leetcode.cn/problems/number-of-recent-calls/) | 🟢 Easy | ✅ [stack_queue/0933](stack_queue/0933-number-of-recent-calls/main.go) |
| 137 | [1146. 快照数组](https://leetcode.cn/problems/snapshot-array/) | 🟡 Medium | ✅ [stack_queue/1146](stack_queue/1146-snapshot-array/main.go) |
| 138 | [1429. 第一个唯一数字 (建表)](https://leetcode.cn/problems/first-unique-number/) | 🟡 Medium | ✅ [stack_queue/1429](stack_queue/1429-first-unique-number/main.go) |
| 139 | [1670. 设计前中后队列](https://leetcode.cn/problems/design-front-middle-back-queue/) | 🟡 Medium | ✅ [stack_queue/1670](stack_queue/1670-design-front-middle-back-queue/main.go) |
| 140 | [1823. 找出游戏的获胜者](https://leetcode.cn/problems/find-the-winner-of-the-circular-game/) | 🟡 Medium | ✅ [stack_queue/1823](stack_queue/1823-find-the-winner-of-the-circular-game/main.go) |

### 📅 Day 4 · 二叉树全专题（49 题）

#### 遍历（前中后序 / 层序 / 构造）（20 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 141 | [94. 二叉树的中序遍历](https://leetcode.cn/problems/binary-tree-inorder-traversal/) | 🟢 Easy | ✅ [tree/0094](tree/0094-binary-tree-inorder-traversal/main.go) |
| 142 | [98. 验证二叉搜索树](https://leetcode.cn/problems/validate-binary-search-tree/) | 🟡 Medium | ✅ [tree/0098](tree/0098-validate-binary-search-tree/main.go) |
| 143 | [99. 恢复二叉搜索树](https://leetcode.cn/problems/recover-binary-search-tree/) | 🟡 Medium | ✅ [tree/0099](tree/0099-recover-binary-search-tree/main.go) |
| 144 | [100. 相同的树](https://leetcode.cn/problems/same-tree/) | 🟢 Easy | ✅ [tree/0100](tree/0100-same-tree/main.go) |
| 145 | [101. 对称二叉树](https://leetcode.cn/problems/symmetric-tree/) | 🟢 Easy | ✅ [tree/0101](tree/0101-symmetric-tree/main.go) |
| 146 | [102. 二叉树的层序遍历](https://leetcode.cn/problems/binary-tree-level-order-traversal/) | 🟡 Medium | ✅ [tree/0102](tree/0102-binary-tree-level-order-traversal/main.go) |
| 147 | [103. 二叉树的锯齿形层序遍历](https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/) | 🟡 Medium | ✅ [tree/0103](tree/0103-binary-tree-zigzag-level-order-traversal/main.go) |
| 148 | [104. 二叉树的最大深度](https://leetcode.cn/problems/maximum-depth-of-binary-tree/) | 🟢 Easy | ✅ [tree/0104](tree/0104-maximum-depth-of-binary-tree/main.go) |
| 149 | [105. 从前序与中序遍历序列构造二叉树](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/) | 🟡 Medium | ✅ [tree/0105](tree/0105-construct-binary-tree-from-preorder-and-inorder-traversal/main.go) |
| 150 | [106. 从中序与后序遍历序列构造二叉树](https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/) | 🟡 Medium | ✅ [tree/0106](tree/0106-construct-binary-tree-from-inorder-and-postorder-traversal/main.go) |
| 151 | [107. 二叉树的层序遍历 II](https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/) | 🟡 Medium | ✅ [tree/0107](tree/0107-binary-tree-level-order-traversal-ii/main.go) |
| 152 | [108. 将有序数组转换为二叉搜索树](https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/) | 🟢 Easy | ✅ [tree/0108](tree/0108-convert-sorted-array-to-binary-search-tree/main.go) |
| 153 | [110. 平衡二叉树](https://leetcode.cn/problems/balanced-binary-tree/) | 🟢 Easy | ✅ [tree/0110](tree/0110-balanced-binary-tree/main.go) |
| 154 | [111. 二叉树的最小深度](https://leetcode.cn/problems/minimum-depth-of-binary-tree/) | 🟢 Easy | ✅ [tree/0111](tree/0111-minimum-depth-of-binary-tree/main.go) |
| 155 | [112. 路径总和](https://leetcode.cn/problems/path-sum/) | 🟢 Easy | ✅ [tree/0112](tree/0112-path-sum/main.go) |
| 156 | [113. 路径总和 II](https://leetcode.cn/problems/path-sum-ii/) | 🟡 Medium | ✅ [tree/0113](tree/0113-path-sum-ii/main.go) |
| 157 | [114. 二叉树展开为链表](https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/) | 🟡 Medium | ✅ [tree/0114](tree/0114-flatten-binary-tree-to-linked-list/main.go) |
| 158 | [116. 填充每个节点的下一个右侧节点指针](https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/) | 🟡 Medium | ✅ [tree/0116](tree/0116-populating-next-right-pointers-in-each-node/main.go) |
| 159 | [117. 填充每个节点的下一个右侧节点指针 II](https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/) | 🟡 Medium | ✅ [tree/0117](tree/0117-populating-next-right-pointers-in-each-node-ii/main.go) |
| 160 | [144. 二叉树的前序遍历](https://leetcode.cn/problems/binary-tree-preorder-traversal/) | 🟢 Easy | ✅ [tree/0144](tree/0144-binary-tree-preorder-traversal/main.go) |

#### BST & 最近公共祖先（15 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 161 | [145. 二叉树的后序遍历](https://leetcode.cn/problems/binary-tree-postorder-traversal/) | 🟢 Easy | ✅ [tree/0145](tree/0145-binary-tree-postorder-traversal/main.go) |
| 162 | [173. 二叉搜索树迭代器](https://leetcode.cn/problems/binary-search-tree-iterator/) | 🟡 Medium | ✅ [tree/0173](tree/0173-binary-search-tree-iterator/main.go) |
| 163 | [199. 二叉树的右视图](https://leetcode.cn/problems/binary-tree-right-side-view/) | 🟡 Medium | ✅ [tree/0199](tree/0199-binary-tree-right-side-view/main.go) |
| 164 | [226. 翻转二叉树](https://leetcode.cn/problems/invert-binary-tree/) | 🟢 Easy | ✅ [tree/0226](tree/0226-invert-binary-tree/main.go) |
| 165 | [230. 二叉搜索树中第 K 小的元素](https://leetcode.cn/problems/kth-smallest-element-in-a-bst/) | 🟡 Medium | ✅ [tree/0230](tree/0230-kth-smallest-element-in-a-bst/main.go) |
| 166 | [235. 二叉搜索树的最近公共祖先](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/) | 🟡 Medium | ✅ [tree/0235](tree/0235-lowest-common-ancestor-of-a-binary-search-tree/main.go) |
| 167 | [236. 二叉树的最近公共祖先](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/) | 🟡 Medium | ✅ [tree/0236](tree/0236-lowest-common-ancestor-of-a-binary-tree/main.go) |
| 168 | [285. 二叉搜索树中的中序后继](https://leetcode.cn/problems/inorder-successor-in-bst/) | 🟡 Medium | ✅ [tree/0285](tree/0285-inorder-successor-in-bst/main.go) |
| 169 | [297. 二叉树的序列化与反序列化](https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/) | 🔴 Hard | ✅ [tree/0297](tree/0297-serialize-and-deserialize-binary-tree/main.go) |
| 170 | [449. 序列化和反序列化二叉搜索树](https://leetcode.cn/problems/serialize-and-deserialize-bst/) | 🟡 Medium | ✅ [tree/0449](tree/0449-serialize-and-deserialize-bst/main.go) |
| 171 | [450. 删除二叉搜索树中的节点](https://leetcode.cn/problems/delete-node-in-a-bst/) | 🟡 Medium | ✅ [tree/0450](tree/0450-delete-node-in-a-bst/main.go) |
| 172 | [501. 二叉搜索树中的众数](https://leetcode.cn/problems/find-mode-in-binary-search-tree/) | 🟢 Easy | ✅ [tree/0501](tree/0501-find-mode-in-binary-search-tree/main.go) |
| 173 | [530. 二叉搜索树的最小绝对差](https://leetcode.cn/problems/minimum-absolute-difference-in-bst/) | 🟢 Easy | ✅ [tree/0530](tree/0530-minimum-absolute-difference-in-bst/main.go) |
| 174 | [538. 把二叉搜索树转换为累加树](https://leetcode.cn/problems/convert-bst-to-greater-tree/) | 🟡 Medium | ✅ [tree/0538](tree/0538-convert-bst-to-greater-tree/main.go) |
| 175 | [543. 二叉树的直径](https://leetcode.cn/problems/diameter-of-binary-tree/) | 🟢 Easy | ✅ [tree/0543](tree/0543-diameter-of-binary-tree/main.go) |

#### 路径 / 深度 / 直径 / 序列化（14 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 176 | [572. 另一棵树的子树](https://leetcode.cn/problems/subtree-of-another-tree/) | 🟢 Easy | ✅ [tree/0572](tree/0572-subtree-of-another-tree/main.go) |
| 177 | [617. 合并二叉树](https://leetcode.cn/problems/merge-two-binary-trees/) | 🟢 Easy | ✅ [tree/0617](tree/0617-merge-two-binary-trees/main.go) |
| 178 | [652. 寻找重复的子树](https://leetcode.cn/problems/find-duplicate-subtrees/) | 🟡 Medium | ✅ [tree/0652](tree/0652-find-duplicate-subtrees/main.go) |
| 179 | [654. 最大二叉树](https://leetcode.cn/problems/maximum-binary-tree/) | 🟡 Medium | ✅ [tree/0654](tree/0654-maximum-binary-tree/main.go) |
| 180 | [669. 修剪二叉搜索树](https://leetcode.cn/problems/trim-a-binary-search-tree/) | 🟡 Medium | ✅ [tree/0669](tree/0669-trim-a-binary-search-tree/main.go) |
| 181 | [700. 二叉搜索树中的搜索](https://leetcode.cn/problems/search-in-a-binary-search-tree/) | 🟢 Easy | ✅ [tree/0700](tree/0700-search-in-a-binary-search-tree/main.go) |
| 182 | [701. 二叉搜索树中的插入操作](https://leetcode.cn/problems/insert-into-a-binary-search-tree/) | 🟡 Medium | ✅ [tree/0701](tree/0701-insert-into-a-binary-search-tree/main.go) |
| 183 | [814. 二叉树剪枝](https://leetcode.cn/problems/binary-tree-pruning/) | 🟡 Medium | ✅ [tree/0814](tree/0814-binary-tree-pruning/main.go) |
| 184 | [863. 二叉树中所有距离为 K 的结点](https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree/) | 🟡 Medium | ✅ [tree/0863](tree/0863-all-nodes-distance-k-in-binary-tree/main.go) |
| 185 | [889. 根据前序和后序遍历构造二叉树](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/) | 🟡 Medium | ✅ [tree/0889](tree/0889-construct-binary-tree-from-preorder-and-postorder-traversal/main.go) |
| 186 | [958. 二叉树的完全性检验](https://leetcode.cn/problems/check-completeness-of-a-binary-tree/) | 🟡 Medium | ✅ [tree/0958](tree/0958-check-completeness-of-a-binary-tree/main.go) |
| 187 | [1110. 删点成林](https://leetcode.cn/problems/delete-nodes-and-return-forest/) | 🟡 Medium | ✅ [tree/1110](tree/1110-delete-nodes-and-return-forest/main.go) |
| 188 | [1245. 树的直径](https://leetcode.cn/problems/tree-diameter/) | 🟡 Medium | ✅ [tree/1245](tree/1245-tree-diameter/main.go) |
| 189 | [1448. 统计二叉树中好节点的数目](https://leetcode.cn/problems/count-good-nodes-in-binary-tree/) | 🟡 Medium | ✅ [tree/1448](tree/1448-count-good-nodes-in-binary-tree/main.go) |

### 📅 Day 5 · 回溯 / 图 / BFS（45 题）

#### 回溯（组合 / 排列 / 子集 / 棋盘）（15 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 190 | [17. 电话号码的字母组合](https://leetcode.cn/problems/letter-combinations-of-a-phone-number/) | 🟡 Medium | ✅ [dfs_bfs/0017](dfs_bfs/0017-letter-combinations-of-a-phone-number/main.go) |
| 191 | [39. 组合总和](https://leetcode.cn/problems/combination-sum/) | 🟡 Medium | ✅ [dfs_bfs/0039](dfs_bfs/0039-combination-sum/main.go) |
| 192 | [40. 组合总和 II](https://leetcode.cn/problems/combination-sum-ii/) | 🟡 Medium | ✅ [dfs_bfs/0040](dfs_bfs/0040-combination-sum-ii/main.go) |
| 193 | [46. 全排列](https://leetcode.cn/problems/permutations/) | 🟡 Medium | ✅ [dfs_bfs/0046](dfs_bfs/0046-permutations/main.go) |
| 194 | [47. 全排列 II](https://leetcode.cn/problems/permutations-ii/) | 🟡 Medium | ✅ [dfs_bfs/0047](dfs_bfs/0047-permutations-ii/main.go) |
| 195 | [51. N 皇后](https://leetcode.cn/problems/n-queens/) | 🔴 Hard | ✅ [dfs_bfs/0051](dfs_bfs/0051-n-queens/main.go) |
| 196 | [52. N 皇后 II](https://leetcode.cn/problems/n-queens-ii/) | 🔴 Hard | ✅ [dfs_bfs/0052](dfs_bfs/0052-n-queens-ii/main.go) |
| 197 | [77. 组合](https://leetcode.cn/problems/combinations/) | 🟡 Medium | ✅ [dfs_bfs/0077](dfs_bfs/0077-combinations/main.go) |
| 198 | [78. 子集](https://leetcode.cn/problems/subsets/) | 🟡 Medium | ✅ [dfs_bfs/0078](dfs_bfs/0078-subsets/main.go) |
| 199 | [79. 单词搜索](https://leetcode.cn/problems/word-search/) | 🟡 Medium | ✅ [dfs_bfs/0079](dfs_bfs/0079-word-search/main.go) |
| 200 | [90. 子集 II](https://leetcode.cn/problems/subsets-ii/) | 🟡 Medium | ✅ [dfs_bfs/0090](dfs_bfs/0090-subsets-ii/main.go) |
| 201 | [93. 复原 IP 地址](https://leetcode.cn/problems/restore-ip-addresses/) | 🟡 Medium | ✅ [dfs_bfs/0093](dfs_bfs/0093-restore-ip-addresses/main.go) |
| 202 | [131. 分割回文串](https://leetcode.cn/problems/palindrome-partitioning/) | 🟡 Medium | ✅ [dfs_bfs/0131](dfs_bfs/0131-palindrome-partitioning/main.go) |
| 203 | [216. 组合总和 III](https://leetcode.cn/problems/combination-sum-iii/) | 🟡 Medium | ✅ [dfs_bfs/0216](dfs_bfs/0216-combination-sum-iii/main.go) |
| 204 | [301. 删除无效的括号](https://leetcode.cn/problems/remove-invalid-parentheses/) | 🔴 Hard | ✅ [dfs_bfs/0301](dfs_bfs/0301-remove-invalid-parentheses/main.go) |

#### 图论（拓扑排序 / 并查集 / 最短路）（15 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 205 | [133. 克隆图](https://leetcode.cn/problems/clone-graph/) | 🟡 Medium | ✅ [graph/0133](graph/0133-clone-graph/main.go) |
| 206 | [200. 岛屿数量](https://leetcode.cn/problems/number-of-islands/) | 🟡 Medium | ✅ [graph/0200](graph/0200-number-of-islands/main.go) |
| 207 | [207. 课程表](https://leetcode.cn/problems/course-schedule/) | 🟡 Medium | ✅ [graph/0207](graph/0207-course-schedule/main.go) |
| 208 | [210. 课程表 II](https://leetcode.cn/problems/course-schedule-ii/) | 🟡 Medium | ✅ [graph/0210](graph/0210-course-schedule-ii/main.go) |
| 209 | [261. 以图判树](https://leetcode.cn/problems/graph-valid-tree/) | 🟡 Medium | ✅ [graph/0261](graph/0261-graph-valid-tree/main.go) |
| 210 | [269. 火星词典](https://leetcode.cn/problems/alien-dictionary/) | 🔴 Hard | ✅ [graph/0269](graph/0269-alien-dictionary/main.go) |
| 211 | [323. 无向图中连通分量的数目](https://leetcode.cn/problems/number-of-connected-components-in-an-undirected-graph/) | 🟡 Medium | ✅ [graph/0323](graph/0323-number-of-connected-components-in-an-undirected-graph/main.go) |
| 212 | [399. 除法求值](https://leetcode.cn/problems/evaluate-division/) | 🟡 Medium | ✅ [graph/0399](graph/0399-evaluate-division/main.go) |
| 213 | [547. 省份数量](https://leetcode.cn/problems/number-of-provinces/) | 🟡 Medium | ✅ [graph/0547](graph/0547-number-of-provinces/main.go) |
| 214 | [684. 冗余连接](https://leetcode.cn/problems/redundant-connection/) | 🟡 Medium | ✅ [graph/0684](graph/0684-redundant-connection/main.go) |
| 215 | [695. 岛屿的最大面积](https://leetcode.cn/problems/max-area-of-island/) | 🟡 Medium | ✅ [graph/0695](graph/0695-max-area-of-island/main.go) |
| 216 | [743. 网络延迟时间](https://leetcode.cn/problems/network-delay-time/) | 🟡 Medium | ✅ [graph/0743](graph/0743-network-delay-time/main.go) |
| 217 | [785. 判断二分图](https://leetcode.cn/problems/is-graph-bipartite/) | 🟡 Medium | ✅ [graph/0785](graph/0785-is-graph-bipartite/main.go) |
| 218 | [787. K 站中转内最便宜的航班](https://leetcode.cn/problems/cheapest-flights-within-k-stops/) | 🟡 Medium | ✅ [graph/0787](graph/0787-cheapest-flights-within-k-stops/main.go) |
| 219 | [997. 找到小镇的法官](https://leetcode.cn/problems/find-the-town-judge/) | 🟢 Easy | ✅ [graph/0997](graph/0997-find-the-town-judge/main.go) |

#### BFS（网格 / 多源 / 状态搜索）（15 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 220 | [127. 单词接龙](https://leetcode.cn/problems/word-ladder/) | 🔴 Hard | ✅ [dfs_bfs/0127](dfs_bfs/0127-word-ladder/main.go) |
| 221 | [130. 被围绕的区域](https://leetcode.cn/problems/surrounded-regions/) | 🟡 Medium | ✅ [dfs_bfs/0130](dfs_bfs/0130-surrounded-regions/main.go) |
| 222 | [286. 墙与门](https://leetcode.cn/problems/walls-and-gates/) | 🟡 Medium | ✅ [dfs_bfs/0286](dfs_bfs/0286-walls-and-gates/main.go) |
| 223 | [317. 离建筑物最近的距离](https://leetcode.cn/problems/shortest-distance-from-all-buildings/) | 🔴 Hard | ✅ [dfs_bfs/0317](dfs_bfs/0317-shortest-distance-from-all-buildings/main.go) |
| 224 | [417. 太平洋大西洋水流问题](https://leetcode.cn/problems/pacific-atlantic-water-flow/) | 🟡 Medium | ✅ [dfs_bfs/0417](dfs_bfs/0417-pacific-atlantic-water-flow/main.go) |
| 225 | [433. 最小基因变化](https://leetcode.cn/problems/minimum-genetic-mutation/) | 🟡 Medium | ✅ [dfs_bfs/0433](dfs_bfs/0433-minimum-genetic-mutation/main.go) |
| 226 | [542. 01 矩阵](https://leetcode.cn/problems/01-matrix/) | 🟡 Medium | ✅ [dfs_bfs/0542](dfs_bfs/0542-01-matrix/main.go) |
| 227 | [752. 打开转盘锁](https://leetcode.cn/problems/open-the-lock/) | 🟡 Medium | ✅ [dfs_bfs/0752](dfs_bfs/0752-open-the-lock/main.go) |
| 228 | [934. 最短的桥](https://leetcode.cn/problems/shortest-bridge/) | 🟡 Medium | ✅ [dfs_bfs/0934](dfs_bfs/0934-shortest-bridge/main.go) |
| 229 | [994. 腐烂的橘子](https://leetcode.cn/problems/rotting-oranges/) | 🟡 Medium | ✅ [dfs_bfs/0994](dfs_bfs/0994-rotting-oranges/main.go) |
| 230 | [1091. 二进制矩阵中的最短路径](https://leetcode.cn/problems/shortest-path-in-binary-matrix/) | 🟡 Medium | ✅ [dfs_bfs/1091](dfs_bfs/1091-shortest-path-in-binary-matrix/main.go) |
| 231 | [1162. 地图分析](https://leetcode.cn/problems/as-far-from-land-as-possible/) | 🟡 Medium | ✅ [dfs_bfs/1162](dfs_bfs/1162-as-far-from-land-as-possible/main.go) |
| 232 | [1293. 网格中的最短路径](https://leetcode.cn/problems/shortest-path-in-a-grid-with-obstacles-elimination/) | 🔴 Hard | ✅ [dfs_bfs/1293](dfs_bfs/1293-shortest-path-in-a-grid-with-obstacles-elimination/main.go) |
| 233 | [1730. 获取食物的最短路径](https://leetcode.cn/problems/shortest-path-to-get-food/) | 🟡 Medium | ✅ [dfs_bfs/1730](dfs_bfs/1730-shortest-path-to-get-food/main.go) |
| 234 | [1926. 迷宫中离入口最近的出口](https://leetcode.cn/problems/nearest-exit-from-entrance-in-maze/) | 🟡 Medium | ✅ [dfs_bfs/1926](dfs_bfs/1926-nearest-exit-from-entrance-in-maze/main.go) |

### 📅 Day 6 · 动态规划(上) / 贪心（34 题）

#### 线性 DP（子序列 / 股票 / 打家劫舍）（12 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 235 | [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/) | 🟡 Medium | ✅ [dp/0053](dp/0053-maximum-subarray/main.go) |
| 236 | [70. 爬楼梯](https://leetcode.cn/problems/climbing-stairs/) | 🟢 Easy | ✅ [dp/0070](dp/0070-climbing-stairs/main.go) |
| 237 | [121. 买卖股票的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/) | 🟢 Easy | ✅ [dp/0121](dp/0121-best-time-to-buy-and-sell-stock/main.go) |
| 238 | [122. 买卖股票的最佳时机 II](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/) | 🟡 Medium | ✅ [dp/0122](dp/0122-best-time-to-buy-and-sell-stock-ii/main.go) |
| 239 | [123. 买卖股票的最佳时机 III](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/) | 🔴 Hard | ✅ [dp/0123](dp/0123-best-time-to-buy-and-sell-stock-iii/main.go) |
| 240 | [152. 乘积最大子数组](https://leetcode.cn/problems/maximum-product-subarray/) | 🟡 Medium | ✅ [dp/0152](dp/0152-maximum-product-subarray/main.go) |
| 241 | [198. 打家劫舍](https://leetcode.cn/problems/house-robber/) | 🟡 Medium | ✅ [dp/0198](dp/0198-house-robber/main.go) |
| 242 | [213. 打家劫舍 II](https://leetcode.cn/problems/house-robber-ii/) | 🟡 Medium | ✅ [dp/0213](dp/0213-house-robber-ii/main.go) |
| 243 | [300. 最长递增子序列 (DP版)](https://leetcode.cn/problems/longest-increasing-subsequence/) | 🟡 Medium | ✅ [dp/0300](dp/0300-longest-increasing-subsequence/main.go) |
| 244 | [322. 零钱兑换](https://leetcode.cn/problems/coin-change/) | 🟡 Medium | ✅ [dp/0322](dp/0322-coin-change/main.go) |
| 245 | [337. 打家劫舍 III](https://leetcode.cn/problems/house-robber-iii/) | 🟡 Medium | ✅ [dp/0337](dp/0337-house-robber-iii/main.go) |
| 246 | [368. 最大整除子集](https://leetcode.cn/problems/largest-divisible-subset/) | 🟡 Medium | ✅ [dp/0368](dp/0368-largest-divisible-subset/main.go) |

#### 背包问题（01 / 完全 / 多重）（8 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 247 | [279. 完全平方数](https://leetcode.cn/problems/perfect-squares/) | 🟡 Medium | ✅ [dp/0279](dp/0279-perfect-squares/main.go) |
| 248 | [416. 分割等和子集](https://leetcode.cn/problems/partition-equal-subset-sum/) | 🟡 Medium | ✅ [dp/0416](dp/0416-partition-equal-subset-sum/main.go) |
| 249 | [474. 一和零](https://leetcode.cn/problems/ones-and-zeroes/) | 🟡 Medium | ✅ [dp/0474](dp/0474-ones-and-zeroes/main.go) |
| 250 | [494. 目标和](https://leetcode.cn/problems/target-sum/) | 🟡 Medium | ✅ [dp/0494](dp/0494-target-sum/main.go) |
| 251 | [518. 零钱兑换 II](https://leetcode.cn/problems/coin-change-ii/) | 🟡 Medium | ✅ [dp/0518](dp/0518-coin-change-ii/main.go) |
| 252 | [1049. 最后一块石头的重量 II](https://leetcode.cn/problems/last-stone-weight-ii/) | 🟡 Medium | ✅ [dp/1049](dp/1049-last-stone-weight-ii/main.go) |
| 253 | [1155. 掷骰子等于目标和的方法数](https://leetcode.cn/problems/number-of-dice-rolls-with-target-sum/) | 🟡 Medium | ✅ [dp/1155](dp/1155-number-of-dice-rolls-with-target-sum/main.go) |
| 254 | [1449. 数位成本和为目标值的最大数字](https://leetcode.cn/problems/form-largest-integer-with-digits-that-sum-to-target/) | 🔴 Hard | ✅ [dp/1449](dp/1449-form-largest-integer-with-digits-that-sum-to-target/main.go) |

#### 贪心（区间 / 跳跃 / 排序贪心）（14 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 255 | [45. 跳跃游戏 II](https://leetcode.cn/problems/jump-game-ii/) | 🟡 Medium | ✅ [greedy/0045](greedy/0045-jump-game-ii/main.go) |
| 256 | [55. 跳跃游戏](https://leetcode.cn/problems/jump-game/) | 🟡 Medium | ✅ [greedy/0055](greedy/0055-jump-game/main.go) |
| 257 | [134. 加油站](https://leetcode.cn/problems/gas-station/) | 🟡 Medium | ✅ [greedy/0134](greedy/0134-gas-station/main.go) |
| 258 | [135. 分发糖果](https://leetcode.cn/problems/candy/) | 🔴 Hard | ✅ [greedy/0135](greedy/0135-candy/main.go) |
| 259 | [253. 会议室 II](https://leetcode.cn/problems/meeting-rooms-ii/) | 🟡 Medium | ✅ [greedy/0253](greedy/0253-meeting-rooms-ii/main.go) |
| 260 | [435. 无重叠区间](https://leetcode.cn/problems/non-overlapping-intervals/) | 🟡 Medium | ✅ [greedy/0435](greedy/0435-non-overlapping-intervals/main.go) |
| 261 | [452. 用最少数量的箭引爆气球](https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/) | 🟡 Medium | ✅ [greedy/0452](greedy/0452-minimum-number-of-arrows-to-burst-balloons/main.go) |
| 262 | [455. 分发饼干](https://leetcode.cn/problems/assign-cookies/) | 🟢 Easy | ✅ [greedy/0455](greedy/0455-assign-cookies/main.go) |
| 263 | [630. 课程表 III](https://leetcode.cn/problems/course-schedule-iii/) | 🔴 Hard | ✅ [greedy/0630](greedy/0630-course-schedule-iii/main.go) |
| 264 | [763. 划分字母区间](https://leetcode.cn/problems/partition-labels/) | 🟡 Medium | ✅ [greedy/0763](greedy/0763-partition-labels/main.go) |
| 265 | [1005. K 次取反后最大化的数组和](https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations/) | 🟢 Easy | ✅ [greedy/1005](greedy/1005-maximize-sum-of-array-after-k-negations/main.go) |
| 266 | [1221. 分割平衡字符串](https://leetcode.cn/problems/split-a-string-in-balanced-strings/) | 🟢 Easy | ✅ [greedy/1221](greedy/1221-split-a-string-in-balanced-strings/main.go) |
| 267 | [1323. 6 和 9 组成的最大数字](https://leetcode.cn/problems/maximum-69-number/) | 🟢 Easy | ✅ [greedy/1323](greedy/1323-maximum-69-number/main.go) |
| 268 | [1710. 卡车上的最大单元数](https://leetcode.cn/problems/maximum-units-on-a-truck/) | 🟢 Easy | ✅ [greedy/1710](greedy/1710-maximum-units-on-a-truck/main.go) |

### 📅 Day 7 · 动态规划(下) / 模考 / Hard 攻坚（32 题）

#### 进阶 DP（区间 / 路径 / 字符串）（14 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 269 | [72. 编辑距离](https://leetcode.cn/problems/edit-distance/) | 🟡 Medium | ✅ [dp/0072](dp/0072-edit-distance/main.go) |
| 270 | [91. 解码方法](https://leetcode.cn/problems/decode-ways/) | 🟡 Medium | ✅ [dp/0091](dp/0091-decode-ways/main.go) |
| 271 | [139. 单词拆分](https://leetcode.cn/problems/word-break/) | 🟡 Medium | ✅ [dp/0139](dp/0139-word-break/main.go) |
| 272 | [221. 最大正方形](https://leetcode.cn/problems/maximal-square/) | 🟡 Medium | ✅ [dp/0221](dp/0221-maximal-square/main.go) |
| 273 | [312. 戳气球](https://leetcode.cn/problems/burst-balloons/) | 🔴 Hard | ✅ [dp/0312](dp/0312-burst-balloons/main.go) |
| 274 | [647. 回文子串](https://leetcode.cn/problems/palindromic-substrings/) | 🟡 Medium | ✅ [dp/0647](dp/0647-palindromic-substrings/main.go) |
| 275 | [516. 最长回文子序列](https://leetcode.cn/problems/longest-palindromic-subsequence/) | 🟡 Medium | ✅ [dp/0516](dp/0516-longest-palindromic-subsequence/main.go) |
| 276 | [1039. 多边形三角剖分的最低得分](https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/) | 🟡 Medium | ✅ [dp/1039](dp/1039-minimum-score-triangulation-of-polygon/main.go) |
| 277 | [1143. 最长公共子序列](https://leetcode.cn/problems/longest-common-subsequence/) | 🟡 Medium | ✅ [dp/1143](dp/1143-longest-common-subsequence/main.go) |
| 278 | [1547. 切棍子的最小成本](https://leetcode.cn/problems/minimum-cost-to-cut-a-stick/) | 🔴 Hard | ✅ [dp/1547](dp/1547-minimum-cost-to-cut-a-stick/main.go) |
| 279 | [62. 不同路径](https://leetcode.cn/problems/unique-paths/) | 🟡 Medium | ✅ [dp/0062](dp/0062-unique-paths/main.go) |
| 280 | [64. 最小路径和](https://leetcode.cn/problems/minimum-path-sum/) | 🟡 Medium | ✅ [dp/0064](dp/0064-minimum-path-sum/main.go) |
| 281 | [120. 三角形最小路径和](https://leetcode.cn/problems/triangle/) | 🟡 Medium | ✅ [dp/0120](dp/0120-triangle/main.go) |
| 282 | [96. 不同的二叉搜索树](https://leetcode.cn/problems/unique-binary-search-trees/) | 🟡 Medium | ✅ [dp/0096](dp/0096-unique-binary-search-trees/main.go) |

#### 全真模考（限时闭卷）（5 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 283 | [146. LRU 缓存 (真题)](https://leetcode.cn/problems/lru-cache/) | 🟡 Medium | ✅ [linkedlist/0146](linkedlist/0146-lru-cache/main.go) |
| 284 | [215. 数组中的第 K 个最大元素](https://leetcode.cn/problems/kth-largest-element-in-an-array/) | 🟡 Medium | ✅ [heap/0215](heap/0215-kth-largest-element-in-an-array/main.go) |
| 285 | [236. 二叉树的最近公共祖先 (真题)](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/) | 🟡 Medium | ✅ [tree/0236](tree/0236-lowest-common-ancestor-of-a-binary-tree/main.go) |
| 286 | [4. 寻找两个正序数组的中位数 (真题)](https://leetcode.cn/problems/median-of-two-sorted-arrays/) | 🔴 Hard | ✅ [binary_search/0004](binary_search/0004-median-of-two-sorted-arrays/main.go) |
| 287 | [42. 接雨水 (真题)](https://leetcode.cn/problems/trapping-rain-water/) | 🔴 Hard | ✅ [two_pointer/0042](two_pointer/0042-trapping-rain-water/main.go) |

#### Hard 题攻坚（8 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 288 | [10. 正则表达式匹配](https://leetcode.cn/problems/regular-expression-matching/) | 🔴 Hard | ✅ [hard/0010](hard/0010-regular-expression-matching/main.go) |
| 289 | [44. 通配符匹配](https://leetcode.cn/problems/wildcard-matching/) | 🔴 Hard | ✅ [hard/0044](hard/0044-wildcard-matching/main.go) |
| 290 | [124. 二叉树中的最大路径和](https://leetcode.cn/problems/binary-tree-maximum-path-sum/) | 🔴 Hard | ✅ [hard/0124](hard/0124-binary-tree-maximum-path-sum/main.go) |
| 291 | [212. 单词搜索 II](https://leetcode.cn/problems/word-search-ii/) | 🔴 Hard | ✅ [trie/0212](trie/0212-word-search-ii/main.go) |
| 292 | [295. 数据流的中位数](https://leetcode.cn/problems/find-median-from-data-stream/) | 🔴 Hard | ✅ [heap/0295](heap/0295-find-median-from-data-stream/main.go) |
| 293 | [297. 二叉树的序列化与反序列化 (复习)](https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/) | 🔴 Hard | ✅ [tree/0297](tree/0297-serialize-and-deserialize-binary-tree/main.go) |
| 294 | [301. 删除无效的括号 (复习)](https://leetcode.cn/problems/remove-invalid-parentheses/) | 🔴 Hard | ✅ [dfs_bfs/0301](dfs_bfs/0301-remove-invalid-parentheses/main.go) |
| 295 | [329. 矩阵中的最长递增路径](https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/) | 🔴 Hard | ✅ [hard/0329](hard/0329-longest-increasing-path-in-a-matrix/main.go) |

#### 数据结构设计补充（5 题）

| # | 题目 | 难度 | 代码 |
| --- | --- | --- | --- |
| 296 | [208. 实现 Trie (前缀树)](https://leetcode.cn/problems/implement-trie-prefix-tree/) | 🟡 Medium | ✅ [trie/0208](trie/0208-implement-trie-prefix-tree/main.go) |
| 297 | [211. 添加与搜索单词 - 数据结构设计](https://leetcode.cn/problems/design-add-and-search-words-data-structure/) | 🟡 Medium | ✅ [trie/0211](trie/0211-design-add-and-search-words-data-structure/main.go) |
| 298 | [355. 设计推特](https://leetcode.cn/problems/design-twitter/) | 🟡 Medium | ✅ [heap/0355](heap/0355-design-twitter/main.go) |
| 299 | [380. O(1) 时间插入、删除和获取随机元素](https://leetcode.cn/problems/insert-delete-getrandom-o1/) | 🟡 Medium | ✅ [hash/0380](hash/0380-insert-delete-getrandom-o1/main.go) |
| 300 | [432. 全 O(1) 的数据结构](https://leetcode.cn/problems/all-oone-data-structure/) | 🔴 Hard | ✅ [hash/0432](hash/0432-all-oone-data-structure/main.go) |

---

## 三、面试专项加成动作

刷题量之外，这三件事决定你能否通过面试：

### 1. 高频题单优先

优先保证以下榜单题目全部覆盖（300 题中至少 180 题来自这里）：

- **LeetCode 热题 100**（[链接](https://leetcode.cn/problem-list/2cktkvj/)）
- **剑指 Offer**（[链接](https://leetcode.cn/problem-list/xb9nqhhg/)）
- **CodeTop 面试高频榜**（[链接](https://leetcode.cn/problem-list/2ckc81c/)）

### 2. 口述 + 手写双训练

每道 Medium / Hard 题完成后执行：

1. 合上代码，**用嘴讲清**：思路 → 边界 → 复杂度
2. 打开纯文本 / 白板，**默写代码**一遍
3. 标注卡壳点，次日早课重做

> 大厂面试淘汰多数发生在“讲不清”，而非“做不出”。

### 3. 限时模考（Day 7 必做）

- **45 分钟 / 题**，闭卷，纯文本写
- 模拟真实面试：先沟通思路 → 再写代码 → 最后自测边界
- 模考题目见 Day 7「全真模考」专区

---

## 四、每日复盘模板

建议每晚 22:00–22:30 填写，形成错题资产：

```markdown
## Day X 复盘
- 完成题数：__ / 当日题量
- 卡壳题号：____（原因：思路 / 边界 / 实现 / 复杂度）
- 今日最典型套路：______
- 明日需重做：____
- 一句话心得：______
```

---

## 五、执行提醒

1. **不要死磕单题**：一眼无思路的 Medium/Hard，20 分钟没进展就看题解，理解后**默写**，效率远高于硬耗。
2. **复习 > 新题**：已有 200+ 基础，回归高频题的手感比刷冷门新题更重要。
3. **按专题而非题号刷**：同一专题连刷能快速形成模式识别（Pattern Recognition），这是面试秒题的关键。
4. **优先保 Medium**：300 题中 Medium 占比最高，是面试主力。Hard 题保证“见过、能讲思路”即可。
5. **难度配比**：本计划 Easy 66 / Medium 196 / Hard 38，符合大厂面试实际考察分布。

> **祝冲刺顺利，offer 到手。**
