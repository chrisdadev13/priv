package main

import (
  "crypto/aes"
  "crypto/cipher"
  "encoding/base64"
  "github.com/joho/godotenv"
  "log"
  "os"
)

func decrypt(toDecrypt string) (string, error) {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatalf("Error loading environment variables file")
    return "", err     
  }   
  
 KEY := os.Getenv("KEY")
 
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
