package main

import (
	"time"
	"fmt"
)

func main() {
	presentTime := time.Now()
	fmt.Println("Time is ", presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
}