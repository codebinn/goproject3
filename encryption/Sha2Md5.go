package encryption

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

/*func main() {
	msg := []byte("fasfhafahsdfl131563")
	f, err := os.Open("test.txt")
	ErrPanic(err)
	defer f.Close()

	str := encryption.Md5Sum(msg)
	fmt.Println(len(str), "Md5Sum:", str)
	str = encryption.Md5New(f)
	fmt.Println(len(str), "Md5New:", str)

	str = encryption.Sha256Sum(msg)
	fmt.Println(len(str), "Sha256Sum:", str)
	str = encryption.Sha256New(f)
	fmt.Println(len(str), "Sha256New:", str)
}*/

func Md5Sum(msg []byte) (res string) {
	h := md5.Sum(msg)
	hash := h[:]
	res = hex.EncodeToString(hash)
	return
}

func Md5New(f *os.File) (res string) {
	h := md5.New()
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if n == 0 {
			fmt.Println("file read ok")
			break
		}
		if err != nil && err != io.EOF {
			return
		}
		h.Write(buf[:n])
	}
	r := h.Sum(nil)
	res = hex.EncodeToString(r)
	return
}

func Sha256Sum(msg []byte) (hexRes string, hashed []byte) {
	h := sha256.Sum256(msg)
	hashed = h[:]
	hexRes = hex.EncodeToString(hashed)
	return
}

func Sha256New(f *os.File) (hexRes string, hashed []byte) {
	h := sha256.New()
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if n == 0 {
			fmt.Println("file read ok")
			break
		}
		if err != nil && err != io.EOF {
			return
		}
		h.Write(buf[:n])
	}
	hashed = h.Sum(nil)
	hexRes = hex.EncodeToString(hashed)
	return
}
