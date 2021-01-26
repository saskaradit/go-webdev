package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	c := getCode("rad@example.com")
	fmt.Println(c)
	c = getCode("rad@exampl.com")
	fmt.Println(c)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("yangmulia"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
