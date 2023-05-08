package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

/*func main() {
	key := []byte("abcdefgh12345678")
	plainText := []byte("hello world asdfgh4537837378343jkl;")
	fmt.Println("plainText:", string(plainText))
	cipherText := encryption.CBCEncrption(key, plainText)
	fmt.Println("cipherText:", string(cipherText))
	plainText2 := encryption.CBCDecrption(key, cipherText)
	fmt.Println("plainText2:", string(plainText2))
}*/

var iv_16 = []byte("12345678abcdefgh")

// 填充明文
func fillPlainText(plainText []byte, blockSize int) (afterPlainText []byte) {
	n := blockSize - len(plainText)%blockSize
	b := []byte{byte(n)}
	bn := bytes.Repeat(b, n)
	afterPlainText = append(plainText, bn...)
	fmt.Println("填充明文：", len(afterPlainText), afterPlainText)
	return
}

// 删除明文的填充
func dropPlainText(plainText []byte) (afterPlainText []byte) {
	n := int(plainText[len(plainText)-1])
	afterPlainText = plainText[:len(plainText)-n]
	fmt.Println("删除明文的填充：", afterPlainText)
	return
}

// CBC密码块链分组模式加密
func CBCEncrption(key []byte, plainText []byte) (cipherText []byte) {
	block, err := aes.NewCipher(key)
	ErrPanic(err)

	afterPlainText := fillPlainText(plainText, block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block, iv_16)
	cipherText = make([]byte, len(afterPlainText))
	blockMode.CryptBlocks(cipherText, afterPlainText)
	return
}

// CBC密码块链分组模式解密
func CBCDecrption(key []byte, cipherText []byte) (plainText []byte) {
	block, err := aes.NewCipher(key)
	ErrPanic(err)

	blockMode := cipher.NewCBCDecrypter(block, iv_16)
	blockMode.CryptBlocks(cipherText, cipherText)

	plainText = dropPlainText(cipherText)
	return
}
