package main

import (
	"fmt"
	"time"
)

func task1() {
	intergerChan := make(chan int)
	go func(intChan chan int) {
		for i := 1; i < 11; i++ {
			intChan <- i
		}
		close(intChan)
	}(intergerChan)

	go func(intChan chan int) {
		for integer := range intChan {
			fmt.Println(integer)
		}
	}(intergerChan)
}

func task2() {
	intergerChan := make(chan int, 100)
	go func(intChan chan int) {
		for i := 1; i < 101; i++ {
			intChan <- i
		}
		close(intChan)
	}(intergerChan)

	go func(intChan chan int) {
		for integer := range intChan {
			fmt.Println(integer)
		}
	}(intergerChan)
}

func main() {
	task1()
	time.Sleep(1 * time.Second)
	task2()
	time.Sleep(1 * time.Second)
}
