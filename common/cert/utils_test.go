// Package cert
// Date       : 2024/8/26 17:39
// Author     : Amu
// Description:
package cert

import "testing"

func TestGetSecretID(t *testing.T) {
	secretID := GetSecretID()
	t.Logf("secret id: %s", secretID)
}

func TestGetSecretKey(t *testing.T) {
	secretKey := GetSecretKey()
	t.Logf("secret key: %s", secretKey)
}
