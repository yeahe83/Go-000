package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Bucket 桶
type Bucket struct {
	Value int64
}

// SlidingWindow 滑动窗口
type SlidingWindow struct {
	Buckets map[int64]*Bucket // map[timestamp] = &Bucket{Value=0}
	//BucketRing *ring.Ring
	Mutex *sync.RWMutex // 读写锁
}

// NewSlidingWindow 构造函数
func NewSlidingWindow() *SlidingWindow {
	w := &SlidingWindow{
		Buckets: make(map[int64]*Bucket),
		//BucketRing: ring.New(10),
		Mutex: &sync.RWMutex{},
	}

	return w
}

// 创建或获取当前时间戳的桶
func (w *SlidingWindow) getCurrentBucket() *Bucket {
	now := time.Now().Unix()
	var bucket *Bucket
	var ok bool

	if bucket, ok = w.Buckets[now]; !ok { // 没有此key
		bucket = &Bucket{} // 初始化新桶
		w.Buckets[now] = bucket
	}

	return bucket
}

// 删除旧桶（超过了10s）
func (w *SlidingWindow) removeOldBuckets() {
	now := time.Now().Unix() - 10

	for timestamp := range w.Buckets {
		if timestamp <= now {
			delete(w.Buckets, timestamp)
		}
	}
}

// Increment 将当前(时间戳)桶的值+i
func (w *SlidingWindow) Increment(i int64) {
	if i == 0 {
		return
	}

	w.Mutex.Lock()
	defer w.Mutex.Unlock()

	b := w.getCurrentBucket()
	b.Value += i
	w.removeOldBuckets()
}

// UpdateMax 将当前(时间戳)桶的值改为n
func (w *SlidingWindow) UpdateMax(n int64) {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()

	b := w.getCurrentBucket()
	if n > b.Value {
		b.Value = n
	}
	w.removeOldBuckets()
}

// Sum 返回now之前10s内所有桶的值之和
func (w *SlidingWindow) Sum(now time.Time) int64 {
	sum := int64(0)

	w.Mutex.RLock()
	defer w.Mutex.RUnlock()

	for timestamp, bucket := range w.Buckets {
		// TODO: configurable rolling window
		if timestamp >= now.Unix()-10 {
			sum += bucket.Value
		}
	}

	return sum
}

// Max 返回now之前10s内的最大值
func (w *SlidingWindow) Max(now time.Time) int64 {
	var max int64

	w.Mutex.RLock()
	defer w.Mutex.RUnlock()

	for timestamp, bucket := range w.Buckets {
		if timestamp >= now.Unix()-10 {
			if bucket.Value > max {
				max = bucket.Value
			}
		}
	}

	return max
}

// Avg 返回now的10s内的平均值
func (w *SlidingWindow) Avg(now time.Time) int64 {
	return w.Sum(now) / 10.0
}

func main() {

	win := NewSlidingWindow()

	go func() {
		for {
			fmt.Printf("max=%d\n", win.Max(time.Now()))
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		win.Increment(rand.Int63n(10))
	}
}

//
// ref:
//   https://github.com/afex/hystrix-go/blob/master/hystrix/rolling/rolling.go
// todo：
//   use ring
//
