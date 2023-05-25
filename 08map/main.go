package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["Madhya Pradesh"] = "Hindi"
	languages["Karnataka"] = "kannada"
	languages["Punjab"] = "Punjabi"
	languages["Maharashtra"] = "Marathi"
	languages["Andhra Pradesh"] = "Telgu"
	languages["Tamil Nadu"] = "Tamil"

	fmt.Println(len(languages))

	for state, lang := range languages {
		fmt.Printf("%v is spoken in %v\n", lang, state)
	}
}