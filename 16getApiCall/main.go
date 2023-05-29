package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var url = "https://reqres.in/api/users/2"

func main() {
	fmt.Println("Get api call")

	getCall()
}

func getCall(){
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(response.Status)

	content, _ := ioutil.ReadAll(response.Body)

	var responseString strings.Builder
	responseString.Write(content)

	fmt.Println(responseString.String()) 
	fmt.Println(string(content)) // another way

}