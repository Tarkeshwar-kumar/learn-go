package main

import (
	"fmt"
	"net/url"
)

var news_url = "https://newsapi.org/v2/everything?q=tesla&from=2023-04-26&sortBy=publishedAt&apiKey=API_KEY"

func main(){
	res, _ := url.Parse(news_url)
	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)
	fmt.Println(res.Query())

	constructUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}
	fmt.Printf("%T", constructUrl)
}