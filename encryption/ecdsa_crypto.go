package encryption

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"os"
)

/*func main() {
	//encryption.EcdsaGenerate("privateKey_ecdsa.pem", "publicKey_ecdsa.pem")
	msg := []byte("fasfhafahsdfl1f4a56f4sd56f4asdf4631563")
	_, hashed := encryption.Sha256Sum(msg)
	rbyte, sbyte := encryption.EcdsaSign("privateKey_ecdsa.pem", hashed)
	fmt.Println(string(rbyte))
	fmt.Println(string(sbyte))
	b := encryption.EcdsaVerify("publicKey_ecdsa.pem", hashed, rbyte, sbyte)
	fmt.Println(b)
}*/

// ECDSA密钥生成
func EcdsaGenerate(pvtKeyFileName, pubKeyFileName string) {
	privatekey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	ErrPanic(err)
	pvt, err := x509.MarshalECPrivateKey(privatekey)
	ErrPanic(err)
	block := pem.Block{
		Type:  "ECDSA Private Key",
		Bytes: pvt,
	}
	f, err := os.Create(pvtKeyFileName)
	defer f.Close()
	err = pem.Encode(f, &block)
	ErrPanic(err)

	publickey := privatekey.PublicKey
	pub, err := x509.MarshalPKIXPublicKey(&publickey)
	ErrPanic(err)
	block = pem.Block{
		Type:  "ECDSA Public Key",
		Bytes: pub,
	}
	f, err = os.Create(pubKeyFileName)
	defer f.Close()
	err = pem.Encode(f, &block)
	ErrPanic(err)
}

// 提取私钥
func GetPvtKeyFromFileEcdsa(filename string) *ecdsa.PrivateKey {
	f, err := os.Open(filename)
	ErrPanic(err)
	defer f.Close()
	stat, err := f.Stat()
	ErrPanic(err)
	buf := make([]byte, stat.Size())
	_, err = f.Read(buf)
	ErrPanic(err)

	block, _ := pem.Decode(buf)

	privatekey, err := x509.ParseECPrivateKey(block.Bytes)
	ErrPanic(err)
	return privatekey
}

// 提取公钥
func GetPubKeyFromFileEcdsa(filename string) *ecdsa.PublicKey {
	f, err := os.Open(filename)
	ErrPanic(err)
	defer f.Close()
	stat, err := f.Stat()
	ErrPanic(err)
	buf := make([]byte, stat.Size())
	_, err = f.Read(buf)
	ErrPanic(err)

	block, _ := pem.Decode(buf)

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	ErrPanic(err)
	publicKey := pubInterface.(*ecdsa.PublicKey)
	return publicKey
}

// ECDSA签名
func EcdsaSign(privfilename string, hash []byte) (rbyte, sbyte []byte) {
	priv := GetPvtKeyFromFileEcdsa(privfilename)
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash)
	ErrPanic(err)
	rbyte, err = r.MarshalText()
	ErrPanic(err)
	sbyte, err = s.MarshalText()
	ErrPanic(err)
	return rbyte, sbyte
}

// ECDSA验证
func EcdsaVerify(pubfilename string, hash, rbyte, sbyte []byte) bool {
	pub := GetPubKeyFromFileEcdsa(pubfilename)
	var r, s big.Int
	err := r.UnmarshalText(rbyte)
	ErrPanic(err)
	err = s.UnmarshalText(sbyte)
	ErrPanic(err)
	b := ecdsa.Verify(pub, hash, &r, &s)
	return b
}
