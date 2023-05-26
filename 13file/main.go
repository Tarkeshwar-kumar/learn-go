package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	toWrite := "This is a text file"
	createdFile, err_Creating_File := os.Create("./demoFile")
	nilCheck(err_Creating_File)

	length, err_writting := io.WriteString(createdFile, toWrite)
	nilCheck(err_writting)

	fmt.Println("length of file is ", length)
	content, err := ioutil.ReadFile("./demoFile")

	nilCheck(err)
	fmt.Println(string(content))
	defer createdFile.Close()
}

func nilCheck(err error){
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	} 
}