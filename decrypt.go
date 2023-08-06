package main

import (
  "crypto/aes"
  "crypto/cipher"
  "encoding/base64"
)

func decrypt(toDecrypt string) (string, error) {
 KEY := "abc&1*~#^2^#s0^=)^^7%b34"
 
 block, err := aes.NewCipher([]byte(KEY))
 if err != nil {
  return "", err
 }

 cipherText := decode(toDecrypt)
 cfb := cipher.NewCFBDecrypter(block, bytes)
 plainText := make([]byte, len(cipherText))
 cfb.XORKeyStream(plainText, cipherText)

 return string(plainText), nil
}

func decode(s string) []byte {
 data, err := base64.StdEncoding.DecodeString(s)
 if err != nil {
  panic(err)
 }
 return data
}
