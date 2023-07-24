package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 1; i <= 1000; i++ {
		go calculateSquare(i, &wg)
	}
	wg.Wait()
	timeEclapsed := time.Since(now)
	fmt.Println("Time eclaspesed is ", timeEclapsed)
}

func calculateSquare(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println(i*i)
}