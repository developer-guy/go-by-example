package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := []byte("TOKEN_SECRET")
	sEnc := b64.StdEncoding.EncodeToString(data)

	fmt.Println("This is the encoded format of string:", sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)

	fmt.Println("This is the decoded format of string:", string(sDec))

}
