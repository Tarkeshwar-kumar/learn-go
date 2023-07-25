package decrypt


func Decrypt(str string) string {
	decrypt_string := ""
	for _, s := range(str) {
		asci := int(s)
		temp := string(asci- 3)
		decrypt_string += temp
	}
	return decrypt_string
}