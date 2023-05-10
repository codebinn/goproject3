package encryption

import (
	"crypto/hmac"
	"crypto/sha256"
)

/*func main() {
	msg := []byte("fasfhafahsdfl131563")
	key := []byte("123456789")
	hashText := encryption.HmacEncrypt(msg, key)
	b := encryption.HmacEqual(msg, key, hashText)
	fmt.Println(b)
}*/

func HmacEncrypt(msg, key []byte) (hashText []byte) {
	h := hmac.New(sha256.New, key)
	h.Write(msg)
	hashText = h.Sum(nil)
	return
}

func HmacEqual(msg, key, hashText []byte) (b bool) {
	h := hmac.New(sha256.New, key)
	h.Write(msg)
	hText := h.Sum(nil)
	b = hmac.Equal(hText, hashText)
	return
}
