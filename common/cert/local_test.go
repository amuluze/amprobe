// Package cert
// Date       : 2024/8/14 16:22
// Author     : Amu
// Description:
package cert

import (
	"testing"
	"time"
)

func TestLocalCertificateGenerate(t *testing.T) {
	config := &Config{
		RenewBefore:   30 * 24,
		CheckInterval: 24 * time.Hour,
		ContactEmail:  "amuluze@163.com",
		CacheDir:      "/Users/amu/Desktop/github/cert/local",
		Domains:       []string{"amvector.amuluze.com"},
	}
	localCertificate := NewLocalCertificate(config)
	err := localCertificate.Generate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestLocalLoad(t *testing.T) {
	config := &Config{
		RenewBefore:   30 * 24,
		CheckInterval: 24 * time.Hour,
		ContactEmail:  "amuluze@163.com",
		CacheDir:      "/Users/amu/Desktop/github/cert/local",
		Domains:       []string{"amvector.amuluze.com"},
	}
	c := NewLocalCertificate(config)
	err := c.Load()
	if err != nil {
		t.Fatal(err)
	}
	expire, err := c.expires()
	if err != nil {
		t.Errorf("error getting expired certificate: %v", err)
	}
	t.Logf("expire: %d", expire)
}
