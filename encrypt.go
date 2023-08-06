package main

import (
  "crypto/aes"
  "crypto/cipher"
  "encoding/base64"
  "github.com/joho/godotenv"
  "log"
  "os"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func encrypt(toEncrypt string) (string, error){
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

  plainText := []byte(toEncrypt)
  cfb := cipher.NewCFBEncrypter(block, bytes)
  cipherText := make([]byte, len(plainText))
  cfb.XORKeyStream(cipherText, plainText)
  return encode(cipherText), nil
}

func encode(b []byte) string {
  return base64.StdEncoding.EncodeToString(b)
}
