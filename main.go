package main

import (
	"fmt"
	"goproject3/encryption"
)

func main() {
	//encryption.EcdsaGenerate("privateKey_ecdsa.pem", "publicKey_ecdsa.pem")
	msg := []byte("fasfhafahsdfl1f4a56f4sd56f4asdf4631563")
	_, hashed := encryption.Sha256Sum(msg)
	rbyte, sbyte := encryption.EcdsaSign("privateKey_ecdsa.pem", hashed)
	fmt.Println(string(rbyte))
	fmt.Println(string(sbyte))
	b := encryption.EcdsaVerify("publicKey_ecdsa.pem", hashed, rbyte, sbyte)
	fmt.Println(b)
}
