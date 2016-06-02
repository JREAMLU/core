package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"io"
)

//HMacSha1 hmac-sha1加密
func HMacSha1(key string, src []byte) (string, error) {
	//hmac ,use sha1
	return _hmac(key, src, sha1.New)
}

//HMacMD5 hmac-md5加密
func HMacMD5(key string, src []byte) (string, error) {
	return _hmac(key, src, md5.New)
}

func _hmac(key string, src []byte, h func() hash.Hash) (string, error) {
	mac := hmac.New(h, []byte(key))
	if _, err := mac.Write(src); err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

//MD5 md5加密
func MD5(src string) (string, error) {
	return _hash(src, md5.New())
}

//Sha1 sha1加密
func Sha1(src string) (string, error) {
	return _hash(src, sha1.New())
}

func _hash(src string, h hash.Hash) (string, error) {
	if _, err := io.WriteString(h, src); err != nil {
		return "", err
	}
	sig := hex.EncodeToString(h.Sum(nil))
	return sig, nil
}
