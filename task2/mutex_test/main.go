package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter struct {
	mutex sync.Mutex
	count int32
}

func (c *Counter) AddCount() {
	c.mutex.Lock()
	c.count++
	c.mutex.Unlock()
}

func (c *Counter) AddCountAtomic() {
	atomic.AddInt32(&c.count, 1)
}

func task1() {
	startTime := time.Now()
	counter := Counter{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(counter *Counter, wg *sync.WaitGroup) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				counter.AddCount()
			}
		}(&counter, &wg)
	}
	wg.Wait()
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println(counter.count)
	fmt.Println("task1用时: ", duration)
}

func task2() {
	startTime := time.Now()
	counter := Counter{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(counter *Counter, wg *sync.WaitGroup) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				counter.AddCountAtomic()
			}
		}(&counter, &wg)
	}
	wg.Wait()
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println(counter.count)
	fmt.Println("task2用时: ", duration)
}

func main() {
	task1()
	task2()
}
