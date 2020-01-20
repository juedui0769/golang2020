package share_mem

import (
	"sync"
	"testing"
	"time"
)

// 线程不安全的
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// 线程安全的
func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var waitGroup sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		waitGroup.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			waitGroup.Done()
		}()
	}
	//time.Sleep(1 * time.Second)
	waitGroup.Wait()
	t.Logf("counter = %d", counter)
}
