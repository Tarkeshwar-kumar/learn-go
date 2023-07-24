package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go send(ch)
	select {
	case msg := <- ch :
		fmt.Println(msg)
	case <- time.After(time.Second) :
		fmt.Println("Time out")
	}
}

func send(ch chan int) {
	time.Sleep(3*time.Second)
	ch <- 10
}