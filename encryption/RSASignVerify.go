package encryption

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
)

/*func main() {
	msg := []byte("fasfhafahsdfl131563")
	_, hashed := encryption.Sha256Sum(msg)
	sig := encryption.RSASign("privateKey.pem", hashed)
	fmt.Println(string(sig))
	b := encryption.RSAVerify("publicKey.pem", hashed, sig)
	fmt.Println(b)
}*/

// RSA签名
func RSASign(pvtKeyFilename string, hashed []byte) []byte {
	pvtKey := GetPvtKeyFromFile(pvtKeyFilename)
	signed, err := rsa.SignPKCS1v15(rand.Reader, pvtKey, crypto.SHA256, hashed)
	ErrPanic(err)
	return signed
}

// RSA验证
func RSAVerify(pubKeyFilename string, hashed []byte, sig []byte) bool {
	pubKey := GetPubKeyFromFile(pubKeyFilename)
	err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed, sig)
	if err == nil {
		return true
	}
	return false
}
