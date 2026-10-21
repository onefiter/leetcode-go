package frontmiddleback

import "testing"

// op：Push* 操作带 val，Pop* 操作断言 want。
type op struct {
	name string
	val  int
	want int
}

func TestFrontMiddleBackQueue(t *testing.T) {
	tests := []struct {
		name string
		ops  []op
	}{
		{"示例1", []op{
			{"PushFront", 1, 0}, {"PushBack", 2, 0}, {"PushMiddle", 3, 0}, {"PushMiddle", 4, 0},
			{"PopFront", 0, 1}, {"PopMiddle", 0, 3}, {"PopMiddle", 0, 4}, {"PopBack", 0, 2}, {"PopFront", 0, -1},
		}},
		{"空队列", []op{
			{"PopFront", 0, -1}, {"PopMiddle", 0, -1}, {"PopBack", 0, -1},
		}},
		{"单元素", []op{
			{"PushMiddle", 7, 0}, {"PopFront", 0, 7}, {"PushFront", 8, 0}, {"PopBack", 0, 8},
			{"PushBack", 9, 0}, {"PopMiddle", 0, 9}, {"PopMiddle", 0, -1},
		}},
		{"偶数长度PushMiddle", []op{
			// [1,2,3,4] PushMiddle(9) -> [1,2,9,3,4]
			{"PushBack", 1, 0}, {"PushBack", 2, 0}, {"PushBack", 3, 0}, {"PushBack", 4, 0},
			{"PushMiddle", 9, 0}, {"PopMiddle", 0, 9}, {"PopMiddle", 0, 2}, {"PopMiddle", 0, 3},
			{"PopFront", 0, 1}, {"PopBack", 0, 4},
		}},
		{"奇数长度PushMiddle", []op{
			// [1,2,3,4,5] PushMiddle(6) -> [1,2,6,3,4,5]，PopMiddle 弹出 6
			{"PushBack", 1, 0}, {"PushBack", 2, 0}, {"PushBack", 3, 0}, {"PushBack", 4, 0},
			{"PushBack", 5, 0}, {"PushMiddle", 6, 0}, {"PopMiddle", 0, 6}, {"PopMiddle", 0, 3},
		}},
		{"只从头部操作", []op{
			{"PushFront", 1, 0}, {"PushFront", 2, 0}, {"PushFront", 3, 0}, // [3,2,1]
			{"PopBack", 0, 1}, {"PopFront", 0, 3}, {"PopFront", 0, 2}, {"PopBack", 0, -1},
		}},
		{"值全相同", []op{
			{"PushMiddle", 5, 0}, {"PushMiddle", 5, 0}, {"PushMiddle", 5, 0},
			{"PopMiddle", 0, 5}, {"PopFront", 0, 5}, {"PopBack", 0, 5}, {"PopMiddle", 0, -1},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Constructor()
			for i, o := range tt.ops {
				got, check := 0, true
				switch o.name {
				case "PushFront":
					q.PushFront(o.val)
					check = false
				case "PushMiddle":
					q.PushMiddle(o.val)
					check = false
				case "PushBack":
					q.PushBack(o.val)
					check = false
				case "PopFront":
					got = q.PopFront()
				case "PopMiddle":
					got = q.PopMiddle()
				case "PopBack":
					got = q.PopBack()
				}
				if check && got != o.want {
					t.Fatalf("第%d步 %s() = %d, want %d", i, o.name, got, o.want)
				}
			}
		})
	}
}

// TestFrontMiddleBackQueueAgainstSlice 用切片朴素模拟对拍。
func TestFrontMiddleBackQueueAgainstSlice(t *testing.T) {
	seed := 2024
	next := func() int {
		seed = (seed*1103515245 + 12345) % (1 << 31)
		return seed
	}
	q := Constructor()
	var s []int
	for step := 0; step < 2000; step++ {
		v := next() % 100
		var got, want int
		switch next() % 6 {
		case 0:
			q.PushFront(v)
			s = append([]int{v}, s...)
			continue
		case 1:
			q.PushMiddle(v)
			m := len(s) / 2
			s = append(s[:m], append([]int{v}, s[m:]...)...)
			continue
		case 2:
			q.PushBack(v)
			s = append(s, v)
			continue
		case 3:
			got, want = q.PopFront(), -1
			if len(s) > 0 {
				want, s = s[0], s[1:]
			}
		case 4:
			got, want = q.PopMiddle(), -1
			if len(s) > 0 {
				m := (len(s) - 1) / 2
				want = s[m]
				s = append(s[:m], s[m+1:]...)
			}
		case 5:
			got, want = q.PopBack(), -1
			if len(s) > 0 {
				want, s = s[len(s)-1], s[:len(s)-1]
			}
		}
		if got != want {
			t.Fatalf("第%d步 got %d, want %d", step, got, want)
		}
	}
}
