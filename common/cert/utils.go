// Package cert
// Date       : 2024/8/13 19:46
// Author     : Amu
// Description:
package cert

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"os"
)

func GetPrivateKey() crypto.PrivateKey {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return privateKey
}

func GetSecretID() string {
	return os.Getenv(SecretID)
}

func GetSecretKey() string {
	return os.Getenv(SecretKey)
}

func FileExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, err
		} else {
			return false, err
		}
	} else {
		if stat.IsDir() {
			return false, fmt.Errorf("[%s] is a directory", path)
		} else {
			return true, nil
		}
	}
}
