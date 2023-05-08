package main

import (
	"fmt"
	"mygo/encryption"
)

func main() {
	key := []byte("abcdefgh12345678")
	plainText := []byte("hello world asdfgh4537837378343jkl;")
	fmt.Println("plainText:", string(plainText))
	cipherText := encryption.CTREnDecryption(key, plainText)
	fmt.Println("cipherText:", string(cipherText))
	plainText2 := encryption.CTREnDecryption(key, cipherText)
	fmt.Println("plainText2:", string(plainText2))
}
