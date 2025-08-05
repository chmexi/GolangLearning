package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func two_coroutine(wg *sync.WaitGroup) {
	wg.Add(2) // 增加两个等待任务
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				fmt.Println(i)
				time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println(i)
				time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			}
		}
	}()
}

func printOdd() error {
	for i := 2; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Println(i)
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		}
	}
	return nil
}

// 任务
type Task struct {
	Name string
	Func func() error
}

// 任务执行结果
type TaskResult struct {
	Name     string
	Duration time.Duration
	Error    error
}

// 任务分发器
type TaskScheduler struct {
	tasks []Task
}

func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{}
}

func (s *TaskScheduler) AddTask(name string, f func() error) {
	s.tasks = append(s.tasks, Task{
		Name: name,
		Func: f,
	})
}

func (s *TaskScheduler) ExecuteTasks() []TaskResult {
	results := make([]TaskResult, 0)
	resultChan := make(chan TaskResult, len(s.tasks)) // 手机执行结果
	wg := sync.WaitGroup{}
	wg.Add(len(s.tasks))
	for _, task := range s.tasks {
		go func() {
			defer wg.Done()
			startTime := time.Now()
			var err error = task.Func()
			endTime := time.Now()
			resultChan <- TaskResult{
				Name:     task.Name,
				Duration: endTime.Sub(startTime),
				Error:    err,
			}
		}()
	}

	wg.Wait()
	close(resultChan)
	for result := range resultChan {
		results = append(results, result)
	}
	return results
}

func main() {
	fmt.Println("nihao")
	var wg sync.WaitGroup
	two_coroutine(&wg)
	wg.Wait()

	scheduler := NewTaskScheduler()
	for i := 0; i < 10; i++ {
		scheduler.AddTask("打印奇数_"+strconv.Itoa(i), printOdd)
	}
	results := scheduler.ExecuteTasks()
	for _, result := range results {
		fmt.Println("name: ", result.Name, " Duration: ", result.Duration)
	}
}
