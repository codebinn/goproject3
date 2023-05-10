package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

/*func main() {
	encryption.RSAGenerate(1024)
	plainText := "123456789qwertyuiopasdfghjklzxcvbnm"
	fmt.Println("plainText:", plainText)
	cipherText := encryption.RSAEncrypt("publicKey.pem", plainText)
	fmt.Println("cipherText:", string(cipherText))
	plainText2 := encryption.RSADecrypt("privateKey.pem", cipherText)
	fmt.Println("plainText2:", string(plainText2))
}*/

// 生成密钥对
func RSAGenerate(keysize int, pvtFilename string, pubFilename string) {
	//保存私钥
	privatekey, err := rsa.GenerateKey(rand.Reader, keysize)
	ErrPanic(err)
	prikey := x509.MarshalPKCS1PrivateKey(privatekey)
	block := pem.Block{
		Type:  "RSA Private Key",
		Bytes: prikey,
	}
	f, err := os.Create(pvtFilename)
	defer f.Close()
	ErrPanic(err)
	err = pem.Encode(f, &block)
	ErrPanic(err)

	//保存公钥
	bublickey := x509.MarshalPKCS1PublicKey(&privatekey.PublicKey)
	block = pem.Block{
		Type:  "RSA Private Key",
		Bytes: bublickey,
	}
	f, err = os.Create(pubFilename)
	defer f.Close()
	ErrPanic(err)
	err = pem.Encode(f, &block)
	ErrPanic(err)
}

// 提取私钥
func GetPvtKeyFromFile(filename string) *rsa.PrivateKey {
	f, err := os.Open(filename)
	ErrPanic(err)
	defer f.Close()
	stat, err := f.Stat()
	ErrPanic(err)
	buf := make([]byte, stat.Size())
	_, err = f.Read(buf)
	ErrPanic(err)

	block, _ := pem.Decode(buf)

	privatekey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	ErrPanic(err)
	return privatekey
}

// 提取公钥
func GetPubKeyFromFile(filename string) *rsa.PublicKey {
	f, err := os.Open(filename)
	ErrPanic(err)
	defer f.Close()
	stat, err := f.Stat()
	ErrPanic(err)
	buf := make([]byte, stat.Size())
	_, err = f.Read(buf)
	ErrPanic(err)

	block, _ := pem.Decode(buf)

	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	ErrPanic(err)
	return publicKey
}

// RSA加密
func RSAEncrypt(pubfile, plainText string) (cipherText []byte) {
	publicKey := GetPubKeyFromFile(pubfile)
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plainText))
	ErrPanic(err)
	return
}

// RSA解密
func RSADecrypt(prifile string, cipherText []byte) string {
	privatekey := GetPvtKeyFromFile(prifile)
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privatekey, cipherText)
	ErrPanic(err)
	return string(plainText)
}
