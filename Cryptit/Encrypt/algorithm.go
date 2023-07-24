package encrypt


func Encrypt(str string) string {
	encrypt_string := ""
	for _, s := range(str) {
		asci := int(s)
		temp := string(asci+ 3)
		encrypt_string += temp
	}
	return encrypt_string
}