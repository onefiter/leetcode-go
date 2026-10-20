// Package channel 演示 channel 的几种典型用途场景。
// 面试里 channel 相关高频考点基本就这几类：
//   - worker pool：固定数量的 worker 并发消费任务，控制并发度；
//   - fan-out/fan-in：一个输入分发给多个 worker，再汇聚结果；
//   - pipeline：多个阶段串联，前一阶段的输出是后一阶段的输入；
//   - done 取消：用关闭 channel 做广播，通知所有协程退出。
package channel

import (
	"sync"
	"time"
)

// Work 一个任务。
type Work struct {
	ID  int
	Val int
}

// Result 一个任务的结果。
type Result struct {
	ID     int
	Double int
}

// WorkerPool 固定并发度的 worker 池。
// 思路：jobs 是任务队列，results 是结果队列；启动 n 个 worker，每个 worker
// 从 jobs 里取任务、处理、把结果写进 results。全部任务处理完后关闭 results，
// 调用方可以用 for range 消费结果。
func WorkerPool(jobs []Work, n int, process func(Work) int) []Result {
	jobsCh := make(chan Work, len(jobs))
	resultsCh := make(chan Result, len(jobs))

	// 启动 n 个 worker。
	for i := 0; i < n; i++ {
		go func() {
			for job := range jobsCh { // jobsCh 关闭后 for range 自动结束
				resultsCh <- Result{ID: job.ID, Double: process(job)}
			}
		}()
	}

	// 投递任务，投完关闭 jobs，通知 worker 没有新任务了。
	for _, j := range jobs {
		jobsCh <- j
	}
	close(jobsCh)

	// 收集结果。
	out := make([]Result, 0, len(jobs))
	for range jobs {
		out = append(out, <-resultsCh)
	}
	return out
}

// FanOutFanIn 扇出扇入。
// 思路：gen 生成数据流 -> 多个 worker（扇出）并发处理 -> 通过 merge 汇聚（扇入）
// 到同一个通道，调用方只看到一个通道。
func FanOutFanIn(nums []int, workers int, process func(int) int) []int {
	in := gen(nums...)

	// 扇出：每个 worker 都从同一个 in 里取数据，各自处理，输出到自己的通道。
	outs := make([]<-chan int, workers)
	for i := 0; i < workers; i++ {
		outs[i] = mapChan(in, process)
	}
	// 扇入：merge 把所有 worker 的输出合并成一个通道。
	merged := merge(outs...)

	var res []int
	for v := range merged {
		res = append(res, v)
	}
	return res
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func mapChan(in <-chan int, process func(int) int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- process(n)
		}
		close(out)
	}()
	return out
}

func merge(chans ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, c := range chans {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				out <- n
			}
		}(c)
	}
	// 等所有上游通道都关闭后，再关闭 out，这样调用方的 for range 能正常结束。
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Pipeline 多阶段管道：stage1 -> stage2 -> stage3。
// 每一阶段是「读上游通道 -> 处理 -> 写下游通道」的协程，天然串成流水线。
func Pipeline(nums []int) []int {
	src := gen(nums...)
	square := mapChan(src, func(x int) int { return x * x })
	addOne := mapChan(square, func(x int) int { return x + 1 })
	var res []int
	for v := range addOne {
		res = append(res, v)
	}
	return res
}

// SlowWorker 一个慢任务，用于演示用 done 提前取消。
// 思路：通过 select 同时监听 done 和 time.After。done 被关闭时，所有等待它的协程
// 会立即收到零值，从而退出——关闭 channel 是最常用的「广播取消」手段。
func SlowWorker(done <-chan struct{}) bool {
	select {
	case <-done:
		return false // 被取消
	case <-time.After(100 * time.Millisecond):
		return true // 正常完成
	}
}
