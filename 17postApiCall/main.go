package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var url = "https://reqres.in/api/users"

func main() {
	// fake json payload

	requestBody := strings.NewReader(`
		{
			"name" : "tarak",
			"job" : "go developer"
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		fmt.Println("Error in post call")
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