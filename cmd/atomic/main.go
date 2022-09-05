package main

import (
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// m := sync.Mutex{}
	// m.Lock()
	// value := 0
	// value++
	// fmt.Println(value)

	// value2 := int32(0)
	// atomic.AddInt32(&value2, 1)
	// fmt.Println(value2)

	// ctx := context.Background()
	// value := atomic.Value{}
	// value.Store(&ctx)
	// ct := value.Load().(context.Context)
	// fmt.Println(ct)
}

// func CalcSum(x, y int) int {
// 	return x + y
// }
func MutexCounter() int {
	goroutinesCount := 0
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			goroutinesCount++
			m.Unlock()
			time.Sleep(time.Microsecond)
			goroutinesCount--
			wg.Done()
		}()
	}
	wg.Wait()
	return goroutinesCount
}

func AtomicCounter() int32 {
	goroutinesCount := int32(0)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&goroutinesCount, 1)
			time.Sleep(time.Microsecond)
			atomic.AddInt32(&goroutinesCount, -1)
			wg.Done()
		}()
	}
	wg.Wait()
	return goroutinesCount
}
