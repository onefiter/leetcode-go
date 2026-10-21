package twitter

import (
	"reflect"
	"testing"
)

// twitterAPI 两种实现共用的接口，同一张操作表跑两遍。
type twitterAPI interface {
	PostTweet(userId, tweetId int)
	GetNewsFeed(userId int) []int
	Follow(followerId, followeeId int)
	Unfollow(followerId, followeeId int)
}

// op 一次操作：name 为 "post"、"feed"、"follow"、"unfollow"；
// post 时 a=userId b=tweetId；follow/unfollow 时 a=follower b=followee；feed 时 a=userId，want 为期望结果。
type op struct {
	name string
	a, b int
	want []int
}

func TestTwitter(t *testing.T) {
	impls := []struct {
		name string
		new  func() twitterAPI
	}{
		{"堆归并", func() twitterAPI { tw := Constructor(); return &tw }},
		{"全量排序", func() twitterAPI { tw := NewTwitterSort(); return &tw }},
	}
	// 用户 1 连发 12 条推文，用于测试只取最近 10 条。
	many := []op{}
	for i := 1; i <= 12; i++ {
		many = append(many, op{"post", 1, i, nil})
	}
	many = append(many, op{"feed", 1, 0, []int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3}})

	tests := []struct {
		name string
		ops  []op
	}{
		{"示例1", []op{
			{"post", 1, 5, nil}, {"feed", 1, 0, []int{5}},
			{"follow", 1, 2, nil}, {"post", 2, 6, nil}, {"feed", 1, 0, []int{6, 5}},
			{"unfollow", 1, 2, nil}, {"feed", 1, 0, []int{5}},
		}},
		{"空新闻流", []op{{"feed", 1, 0, []int{}}}},
		{"只取最近10条", many},
		{"多人交错归并", []op{
			{"post", 1, 1, nil}, {"post", 2, 2, nil}, {"post", 3, 3, nil},
			{"post", 1, 4, nil}, {"post", 2, 5, nil},
			{"follow", 1, 2, nil}, {"follow", 1, 3, nil},
			{"feed", 1, 0, []int{5, 4, 3, 2, 1}}, {"feed", 2, 0, []int{5, 2}},
		}},
		{"关注自己无效且不重复", []op{
			{"follow", 1, 1, nil}, {"post", 1, 7, nil}, {"feed", 1, 0, []int{7}},
			{"unfollow", 1, 1, nil}, {"feed", 1, 0, []int{7}},
		}},
		{"取关未关注的人", []op{
			{"post", 2, 9, nil}, {"unfollow", 1, 2, nil}, {"feed", 1, 0, []int{}},
			{"follow", 1, 2, nil}, {"follow", 1, 2, nil}, {"feed", 1, 0, []int{9}},
		}},
	}
	for _, impl := range impls {
		for _, tt := range tests {
			t.Run(impl.name+"/"+tt.name, func(t *testing.T) {
				tw := impl.new()
				for i, o := range tt.ops {
					switch o.name {
					case "post":
						tw.PostTweet(o.a, o.b)
					case "follow":
						tw.Follow(o.a, o.b)
					case "unfollow":
						tw.Unfollow(o.a, o.b)
					case "feed":
						if got := tw.GetNewsFeed(o.a); !reflect.DeepEqual(got, o.want) {
							t.Errorf("op#%d GetNewsFeed(%d) = %v, want %v", i, o.a, got, o.want)
						}
					}
				}
			})
		}
	}
}
