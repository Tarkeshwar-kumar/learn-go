package main

import "github.com/Tarkeshwar-kumar/learn-go/Cryptit/encrypt"
import "github.com/Tarkeshwar-kumar/learn-go/Cryptit/decrypt"
import "fmt"


func main() {
	encrypted_text := encrypt.Encrypt("Tarkeshwar")
	fmt.Println(encrypted_text)
	fmt.Println(decrypt.Decrypt(encrypted_text))
	
}