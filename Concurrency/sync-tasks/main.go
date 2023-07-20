package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting executing main function")
	now := time.Now()
	go task1()  // its a fork
	go task2()  // its a fork
	task3()  // sync task
	task4()  // sync task
	fmt.Println("time elapsed in execution", time.Since(now))
	// not joining created fork
	fmt.Println("Exiting main function")
}

func task1() {
	time.Sleep(100*time.Millisecond)
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

func task4() {
	time.Sleep(100*time.Millisecond)
	fmt.Println("task4")
}