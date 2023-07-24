package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg, wg2 sync.WaitGroup
	wg.Add(10)
	wg2.Add(10)
	for i:=1; i<=10; i++{
		go func ()  {
			fmt.Println(i)
			wg2.Done()	
		}()
	}
	for i:=1; i<=10; i++{
		go func (i int)  {
			fmt.Println(i)
			wg.Done()	
		}(i)
	}
	wg.Wait()
	wg2.Wait()
}