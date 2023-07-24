package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Starting execution of main")

	var wg sync.WaitGroup
	wg.Add(2)
	go leak(&wg)
	wg.Wait()
}

func leak(wg *sync.WaitGroup){
	fmt.Println("Starting leak method")
	ch := make(chan int)
	go func() {
		val := <- ch
		fmt.Println("Var received is", val)
		wg.Done()
	}()
	wg.Done()
	fmt.Println("Exiting leak method")
}