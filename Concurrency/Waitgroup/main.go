package main

import (
	"fmt"
	"sync"
	"time"
)
// different ways to implement waitgroups and gorourtines in go
func main() {
	fmt.Println("Starting executing main function")
	var wg sync.WaitGroup
	now := time.Now()
	wg.Add(4)
	go task1(&wg)  
	go func(){
		defer wg.Done()
		task2()
	}()
	go task3()
	go task4(&wg)
	defer wg.Wait()
	defer fmt.Println("time elapsed in execution", time.Since(now))
	defer wg.Done()
	// not joining created fork
	fmt.Println("Exiting main function")
}

func task1(wg *sync.WaitGroup) {
	time.Sleep(100*time.Millisecond)
	wg.Done()
	fmt.Println("task1")
}

func task2() {
	time.Sleep(300*time.Millisecond)
	fmt.Println("task2")
}

func task3() {
	time.Sleep(200*time.Millisecond)
	fmt.Println("task3")
}

func task4(wg *sync.WaitGroup) {
	time.Sleep(100*time.Millisecond)
	wg.Done()
	fmt.Println("task4")
}