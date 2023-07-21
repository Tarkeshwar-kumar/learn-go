package main

import "fmt"

func main() {
	channel := make(chan string)


    //////////////////////////////////////////   checking if channel is open or close ////////////////////////////////////
	go function(channel)

	for {
		err, ok := <- channel
		if ok == true {
			fmt.Println("Channel is open ", err, ok)
		} else{
			fmt.Println("Channel is closed")
			break
		}
	}

    //////////////////////////////////////////////   checking if channel is open or close /////////////////////////////////

	//////////////////////////////////////////////   programme to learn for loop in channel  /////////////////////////////////

	channel2 := make(chan string)

	go func () {
		channel2 <- "My"
		channel2 <- "Name"
		channel2 <- "Is"
		channel2 <- "Tarkeshwar"
		channel2 <- "Kumar"
		close(channel2)
	}()

	for _ = range(channel2){
		fmt.Println(<-channel2)
	}
	//////////////////////////////////////////////   programme to learn for loop in channel  /////////////////////////////////
}

func function(channel chan string) {
	for i:=0; i<3; i++ {
		channel <- "I love Go"
	}
	close(channel)
}