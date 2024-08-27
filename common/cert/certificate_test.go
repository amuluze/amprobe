// Package cert
// Date       : 2024/8/13 10:27
// Author     : Amu
// Description:
package cert

import (
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	config := &Config{
		RenewBefore:   30 * 24,
		CheckInterval: 24 * time.Hour,
		ContactEmail:  "amuluze@163.com",
		CacheDir:      "/Users/amu/Desktop/github/cert/certs",
		Domains:       []string{"amprobe.amuluze.com"},
	}
	c := NewCertificate(config)
	err := c.Generate()
	if err != nil {
		t.Errorf("Error generating certificate: %s", err)
	}
}

func TestLoad(t *testing.T) {
	config := &Config{
		RenewBefore:   30 * 24,
		CheckInterval: 24 * time.Hour,
		ContactEmail:  "amuluze@163.com",
		CacheDir:      "/Users/amu/Desktop/github/cert/certs",
		Domains:       []string{"amvector.amuluze.com"},
	}
	c := NewCertificate(config)
	expire, err := c.expires()
	if err != nil {
		t.Errorf("Error loading certificate: %s", err)
	}
	t.Logf("expire: %d", expire)
}
