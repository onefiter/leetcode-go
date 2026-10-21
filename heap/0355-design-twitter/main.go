package twitter

import (
	"container/heap"
	"sort"
)

// 355. 设计推特
// https://leetcode.cn/problems/design-twitter/
// 设计简化版推特：PostTweet(userId, tweetId) 发推；GetNewsFeed(userId) 返回该用户及其关注者
// 最近的 10 条推文 id（从新到旧）；Follow / Unfollow 关注与取关（不能关注自己）。
//
// 核心思路：每个用户维护自己的推文列表（按时间追加，天然有序），用全局递增时间戳标记先后；
// 拉取新闻流就是对「自己 + 关注的人」这 k 个有序列表做 k 路归并取前 10，
// 用大顶堆存每个列表当前最新的一条，每弹出一条就把该列表的上一条压入堆。

type tweet struct {
	id, time int
}

// cursor 堆元素：指向某个用户推文列表中的一个位置。
type cursor struct {
	tweets []tweet
	idx    int // 当前指向的下标，从后往前（新到旧）移动
}

// maxHeap 按 cursor 指向推文的时间戳排序的大顶堆。
type maxHeap []cursor

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	return h[i].tweets[h[i].idx].time > h[j].tweets[h[j].idx].time
}
func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)   { *h = append(*h, x.(cursor)) }
func (h *maxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

const feedSize = 10

// Twitter 解法一：每用户推文列表 + 大顶堆 k 路归并 ⭐ 最优方案
// PostTweet / Follow / Unfollow 时间复杂度 O(1)；
// GetNewsFeed 时间复杂度 O(k + 10 log k)，k 为自己加关注的人数。空间复杂度 O(用户数 + 推文数 + 关注关系数)。
type Twitter struct {
	time      int
	tweets    map[int][]tweet
	following map[int]map[int]bool
}

func Constructor() Twitter {
	return Twitter{
		tweets:    map[int][]tweet{},
		following: map[int]map[int]bool{},
	}
}

func (this *Twitter) PostTweet(userId, tweetId int) {
	this.time++
	this.tweets[userId] = append(this.tweets[userId], tweet{id: tweetId, time: this.time})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	h := &maxHeap{}
	push := func(uid int) {
		if ts := this.tweets[uid]; len(ts) > 0 {
			*h = append(*h, cursor{tweets: ts, idx: len(ts) - 1})
		}
	}
	push(userId)
	for followee := range this.following[userId] {
		push(followee)
	}
	heap.Init(h) // O(k) 建堆
	feed := make([]int, 0, feedSize)
	for h.Len() > 0 && len(feed) < feedSize {
		top := (*h)[0]
		feed = append(feed, top.tweets[top.idx].id)
		if top.idx > 0 {
			(*h)[0].idx-- // 该列表还有更早的推文，替换堆顶后下沉
			heap.Fix(h, 0)
		} else {
			heap.Pop(h)
		}
	}
	return feed
}

func (this *Twitter) Follow(followerId, followeeId int) {
	if followerId == followeeId {
		return
	}
	if this.following[followerId] == nil {
		this.following[followerId] = map[int]bool{}
	}
	this.following[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId, followeeId int) {
	delete(this.following[followerId], followeeId)
}

// TwitterSort 解法二：收集全部相关推文后排序（对照用）
// 新闻流时把自己和关注者的所有推文收集起来按时间倒序排序，取前 10。
// GetNewsFeed 时间复杂度 O(N log N)，N 为相关推文总数；其他操作 O(1)。
type TwitterSort struct {
	time      int
	tweets    map[int][]tweet
	following map[int]map[int]bool
}

func NewTwitterSort() TwitterSort {
	return TwitterSort{
		tweets:    map[int][]tweet{},
		following: map[int]map[int]bool{},
	}
}

func (this *TwitterSort) PostTweet(userId, tweetId int) {
	this.time++
	this.tweets[userId] = append(this.tweets[userId], tweet{id: tweetId, time: this.time})
}

func (this *TwitterSort) GetNewsFeed(userId int) []int {
	all := append([]tweet{}, this.tweets[userId]...)
	for followee := range this.following[userId] {
		all = append(all, this.tweets[followee]...)
	}
	sort.Slice(all, func(i, j int) bool { return all[i].time > all[j].time })
	feed := make([]int, 0, feedSize)
	for i := 0; i < len(all) && i < feedSize; i++ {
		feed = append(feed, all[i].id)
	}
	return feed
}

func (this *TwitterSort) Follow(followerId, followeeId int) {
	if followerId == followeeId {
		return
	}
	if this.following[followerId] == nil {
		this.following[followerId] = map[int]bool{}
	}
	this.following[followerId][followeeId] = true
}

func (this *TwitterSort) Unfollow(followerId, followeeId int) {
	delete(this.following[followerId], followeeId)
}
