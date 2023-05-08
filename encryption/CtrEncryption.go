package encryption

import (
	"crypto/aes"
	"crypto/cipher"
)

/*func main() {
	key := []byte("abcdefgh12345678")
	plainText := []byte("hello world asdfgh4537837378343jkl;")
	fmt.Println("plainText:", string(plainText))
	cipherText := encryption.CTREnDecryption(key, plainText)
	fmt.Println("cipherText:", string(cipherText))
	plainText2 := encryption.CTREnDecryption(key, cipherText)
	fmt.Println("plainText2:", string(plainText2))
}*/

// CTR分组模式加密解密
func CTREnDecryption(key []byte, plainText []byte) (cipherText []byte) {
	block, err := aes.NewCipher(key)
	ErrPanic(err)
	Stream := cipher.NewCTR(block, iv_16)
	cipherText = make([]byte, len(plainText))
	Stream.XORKeyStream(cipherText, plainText)
	return
}
