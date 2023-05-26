package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var url = "http://www.google.com"
func main() {
	response, err := http.Get(url)
	nilCheck(err)
	
	byteContent, err := ioutil.ReadAll(response.Body)
	nilCheck(err)

	createdFile, err_Creating_File := os.Create("./responseFile")
	nilCheck(err_Creating_File)

	_, err_writting := io.WriteString(createdFile, string(byteContent))
	nilCheck(err_writting)

	response.Body.Close()

}

func nilCheck(err error){
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	} 
}