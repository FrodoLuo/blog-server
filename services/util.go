package services

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
)

/*
IsPathExist check whether the path is exist
*/
func IsPathExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

/*
Decipher return the origin data of the ciphered password
*/
func Decipher(ciphered string, key string) (string, error) {
	cipheredBytes, _ := base64.StdEncoding.DecodeString(ciphered)
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()
	iv := keyBytes[:blockSize]

	des := make([]byte, len(cipheredBytes))

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(des, cipheredBytes)
	//解填充
	return string(des), nil
}
