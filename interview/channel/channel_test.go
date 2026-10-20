package channel

import (
	"reflect"
	"slices"
	"testing"
)

func double(x int) int { return x * 2 }

func TestWorkerPool(t *testing.T) {
	jobs := []Work{{0, 1}, {1, 2}, {2, 3}, {3, 4}}
	got := WorkerPool(jobs, 2, func(w Work) int { return w.Val * 2 })

	want := []int{2, 4, 6, 8}
	vals := make([]int, 0, len(got))
	for _, r := range got {
		vals = append(vals, r.Double)
	}
	slices.Sort(vals)
	if !slices.Equal(vals, want) {
		t.Errorf("WorkerPool 结果 = %v, want %v", vals, want)
	}
}

func TestFanOutFanIn(t *testing.T) {
	got := FanOutFanIn([]int{1, 2, 3, 4, 5, 6}, 3, double)
	slices.Sort(got)
	want := []int{2, 4, 6, 8, 10, 12}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FanOutFanIn = %v, want %v", got, want)
	}
}

func TestPipeline(t *testing.T) {
	// (x^2)+1 对 [1,2,3] -> [2,5,10]
	got := Pipeline([]int{1, 2, 3})
	want := []int{2, 5, 10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Pipeline = %v, want %v", got, want)
	}
}

func TestSlowWorker(t *testing.T) {
	done := make(chan struct{})
	close(done) // 提前取消
	if SlowWorker(done) {
		t.Error("期望被取消返回 false，实际返回 true")
	}
	if !SlowWorker(make(chan struct{})) {
		t.Error("未取消时应正常完成返回 true，实际返回 false")
	}
}
