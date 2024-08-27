// Package cert
// Date       : 2024/8/13 18:09
// Author     : Amu
// Description:
package cert

import (
	"fmt"
	"time"
)

const (
	EC256   = "P256"
	EC384   = "P384"
	RSA2048 = "2048"
	RSA4096 = "4096"
	RSA8192 = "8192"
)

var SupportedKeyTypes = map[string]bool{
	EC256:   true,
	EC384:   true,
	RSA2048: true,
	RSA4096: true,
	RSA8192: true,
}

type Config struct {
	// RenewBefore renew the certificate X hours before it expires
	RenewBefore int

	// CheckInterval interval for checking if cert is closer to expiration than RenewBefore
	CheckInterval time.Duration

	// SSLEmail email for contact
	ContactEmail string

	// Domains for which to obtain the certificate
	Domains  []string
	CacheDir string
}

func CheckConfig(c *Config) error {
	if c.RenewBefore == 0 {
		return fmt.Errorf("RenewBefore must be greater than zero")
	}
	if c.CheckInterval == 0 {
		return fmt.Errorf("CheckInterval must be greater than zero")
	}
	if len(c.Domains) == 0 {
		return fmt.Errorf("at least one domain must be provided")
	}
	if c.CacheDir == "" {
		c.CacheDir = "certs"
	}
	return nil
}
