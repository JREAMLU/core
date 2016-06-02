package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
)

// CookieCrypto 用于Cookie加解密
type CookieCrypto struct {
	encryptionKey []byte
	validationKey []byte
}

const (
	hashSize int = 8
)

// UpdateKeys 设置密钥
func (k *CookieCrypto) UpdateKeys(ekey, vkey string) {
	encryptionKey, err := hex.DecodeString(ekey)
	if err != nil {
		panic(err)
	}
	validationKey, err := hex.DecodeString(vkey)
	if err != nil {
		panic(err)
	}

	k.encryptionKey = encryptionKey
	k.validationKey = validationKey
}

func addPCK5(s string) []byte {
	plaintext := []byte(s)

	l := aes.BlockSize - (len(plaintext) % aes.BlockSize)
	padding := make([]byte, l)
	for i := 0; i < l; i++ {
		padding[i] = byte(l)
	}
	plaintext = append(plaintext, padding...)

	return plaintext
}

func removePCK5(b []byte) ([]byte, error) {
	if len(b) == 0 {
		return b, errors.New("")
	}
	l := int(b[len(b)-1])
	if len(b)-l < 0 {
		return nil, errors.New("")
	}
	return b[:len(b)-l], nil
}

// Encrypt Cookie加密
func (k *CookieCrypto) Encrypt(clearData string) (string, error) {
	//padding pck5 string byte[]
	plaintext := addPCK5(clearData)

	//get iv
	iv := make([]byte, aes.BlockSize)
	io.ReadFull(rand.Reader, iv)

	//get AES cipher
	block, err := aes.NewCipher(k.encryptionKey)
	if err != nil {
		return "", err
	}

	//get CBC mode
	mode := cipher.NewCBCEncrypter(block, iv)

	//aes encrypt data
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)

	//get mac data
	var hashData = append(iv, ciphertext...)
	mac := hmac.New(sha256.New, k.validationKey)
	mac.Write(hashData)
	hash := mac.Sum(nil)
	hash = hash[:hashSize]

	var encryptData = hex.EncodeToString(iv) + hex.EncodeToString(ciphertext) + hex.EncodeToString(hash)

	return encryptData, nil
}

// Decrypt Cookie解密
func (k *CookieCrypto) Decrypt(s string) (string, error) {
	//check encryptData mod == 0
	if len(s) < 48 || len(s)&1 == 1 {
		return "", nil
	}

	encryptData, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}

	//validate encryption data
	hash := encryptData[len(encryptData)-hashSize:]
	needHashData := encryptData[:len(encryptData)-hashSize]
	mac := hmac.New(sha256.New, k.validationKey)
	mac.Write(needHashData)
	newhash := mac.Sum(nil)
	newhash = hash[:hashSize]

	if !bytes.Equal(hash, newhash) {
		return "", nil
	}

	iv := encryptData[0:aes.BlockSize]

	//get AES cipher
	block, err := aes.NewCipher(k.encryptionKey)
	if err != nil {
		return "", err
	}

	//get CBC mode
	mode := cipher.NewCBCDecrypter(block, iv)

	var ciphertext = encryptData[aes.BlockSize : len(encryptData)-hashSize]
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext, err = removePCK5(ciphertext)

	return string(ciphertext), err
}
