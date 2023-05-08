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
	cipherText := encryption.RSAEncrypt("bublicKey.pem", plainText)
	fmt.Println("cipherText:", string(cipherText))
	plainText2 := encryption.RSADecrypt("privateKey.pem", cipherText)
	fmt.Println("plainText2:", string(plainText2))
}*/

func RSAGenerate(keysize int) {
	//保存私钥
	privatekey, err := rsa.GenerateKey(rand.Reader, keysize)
	ErrPanic(err)
	prikey := x509.MarshalPKCS1PrivateKey(privatekey)
	block := pem.Block{
		Type:  "RSA Private Key",
		Bytes: prikey,
	}
	f, err := os.Create("privateKey.pem")
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
	f, err = os.Create("bublicKey.pem")
	defer f.Close()
	ErrPanic(err)
	err = pem.Encode(f, &block)
	ErrPanic(err)
}

// RSA加密
func RSAEncrypt(pubfile, plainText string) (cipherText []byte) {
	f, err := os.Open(pubfile)
	ErrPanic(err)
	defer f.Close()
	stat, err := f.Stat()
	ErrPanic(err)
	buf := make([]byte, stat.Size())
	_, err = f.Read(buf)
	ErrPanic(err)

	block, _ := pem.Decode(buf)

	buplickey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	ErrPanic(err)
	cipherText, err = rsa.EncryptPKCS1v15(rand.Reader, buplickey, []byte(plainText))
	ErrPanic(err)
	return
}

// RSA解密
func RSADecrypt(prifile string, cipherText []byte) (plainText []byte) {
	f, err := os.Open(prifile)
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
	plainText, err = rsa.DecryptPKCS1v15(rand.Reader, privatekey, []byte(cipherText))
	ErrPanic(err)
	return
}
