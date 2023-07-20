package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Starting executing main function")
	var wg sync.WaitGroup
	wg.Add(4)
	now := time.Now()
	go task1(&wg)  
	go task2(&wg) 
	go task3(&wg)
	go task4(&wg)
	wg.Wait()
	fmt.Println("time elapsed in execution", time.Since(now))
	// not joining created fork
	fmt.Println("Exiting main function")
}

func task1(wg *sync.WaitGroup) {
	time.Sleep(100*time.Millisecond)
	wg.Done()
	fmt.Println("task1")
}

func task2(wg *sync.WaitGroup) {
	time.Sleep(300*time.Millisecond)
	wg.Done()
	fmt.Println("task2")
}

func task3(wg *sync.WaitGroup) {
	time.Sleep(200*time.Millisecond)
	wg.Done()
	fmt.Println("task3")
}

func task4(wg *sync.WaitGroup) {
	time.Sleep(100*time.Millisecond)
	wg.Done()
	fmt.Println("task4")
}