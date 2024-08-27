// Package cert
// Date       : 2024/8/13 10:27
// Author     : Amu
// Description:
package cert

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/providers/dns/tencentcloud"
	"github.com/go-acme/lego/v4/registration"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var _ ICertificate = (*Certificate)(nil)

type ICertificate interface {
	Generate() error
	Renew() error
}

type Certificate struct {
	Domains               []string `json:"domains"`
	Domain                string   `json:"domain"`
	PrivateKey            []byte   `json:"private_key"`
	Certificate           []byte   `json:"certificate"`
	IssuerCertificate     []byte   `json:"issuer_certificate"`
	CSR                   []byte   `json:"csr"`
	DomainPath            string   `json:"domain_path"`
	PrivateKeyPath        string   `json:"private_key_path"`
	CertificatePath       string   `json:"certificate_path"`
	IssuerCertificatePath string   `json:"issuer_certificate_path"`
	CSRPath               string   `json:"csr_path"`
	UserPath              string   `json:"user_path"`
	RenewBefore           int
}

func NewCertificate(config *Config) *Certificate {
	cert := &Certificate{
		PrivateKeyPath:        filepath.Join(config.CacheDir, PrivateKeyFileName),
		CertificatePath:       filepath.Join(config.CacheDir, CertificateFileName),
		UserPath:              filepath.Join(config.CacheDir, UserFileName),
		DomainPath:            filepath.Join(config.CacheDir, DomainFileName),
		IssuerCertificatePath: filepath.Join(config.CacheDir, IssuerCertificateFileName),
		CSRPath:               filepath.Join(config.CacheDir, CSRFileName),
		Domains:               config.Domains,
		RenewBefore:           config.RenewBefore,
	}
	return cert
}

func (c *Certificate) Load() error {
	if _, err := FileExists(c.DomainPath); err != nil {
		return err
	}
	domainData, err := os.ReadFile(c.DomainPath)
	if err != nil {
		return err
	}
	c.Domain = strings.TrimSpace(string(domainData))

	if _, err := FileExists(c.PrivateKeyPath); err != nil {
		return err
	}
	keyData, err := os.ReadFile(c.PrivateKeyPath)
	if err != nil {
		return err
	}
	c.PrivateKey = keyData

	if _, err = FileExists(c.CertificatePath); err != nil {
		return err
	}
	certData, err := os.ReadFile(c.CertificatePath)
	if err != nil {
		return err
	}
	c.Certificate = certData

	if _, err := FileExists(c.IssuerCertificatePath); err != nil {
		return err
	}
	issuerData, err := os.ReadFile(c.IssuerCertificatePath)
	if err != nil {
		return err
	}
	c.IssuerCertificate = issuerData

	if _, err = FileExists(c.CSRPath); err != nil {
		return err
	}
	csrData, err := os.ReadFile(c.CSRPath)
	if err != nil {
		return err
	}
	c.CSR = csrData
	return nil
}

func (c *Certificate) Generate() error {
	u, err := c.getUser()
	if err != nil {
		return err
	}
	client, err := c.createClient(u)
	if err != nil {
		return err
	}
	request := certificate.ObtainRequest{
		Domains: c.Domains,
		Bundle:  true,
	}
	certificates, err := client.Certificate.Obtain(request)
	if err != nil {
		return fmt.Errorf("Obtain error: %v\n", err)
	}

	c.Domain = certificates.Domain
	c.Certificate = certificates.Certificate
	c.PrivateKey = certificates.PrivateKey
	c.CSR = certificates.CSR
	c.IssuerCertificate = certificates.IssuerCertificate

	// save to file
	if err := c.save(); err != nil {
		return err
	}
	return nil
}

func (c *Certificate) Renew() error {
	expire, err := c.expires()
	if err != nil {
		return err
	}
	if c.RenewBefore >= expire {
		u, err := c.getUser()
		if err != nil {
			return err
		}
		client, err := c.createClient(u)
		if err != nil {
			return err
		}
		resource, err := client.Certificate.RenewWithOptions(certificate.Resource{
			Domain:            c.Domain,
			CertURL:           lego.LEDirectoryProduction,
			CertStableURL:     lego.LEDirectoryStaging,
			PrivateKey:        c.PrivateKey,
			Certificate:       c.Certificate,
			IssuerCertificate: c.IssuerCertificate,
			CSR:               c.CSR,
		}, &certificate.RenewOptions{})
		if err != nil {
			return err
		}
		c.CSR = resource.CSR
		c.IssuerCertificate = resource.IssuerCertificate
		c.PrivateKey = resource.PrivateKey
		c.Certificate = resource.Certificate
		c.Domain = resource.Domain

		err = c.save()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Certificate) createClient(u *User) (lego.Client, error) {
	config := lego.NewConfig(u)
	config.CADirURL = lego.LEDirectoryStaging
	config.Certificate.KeyType = certcrypto.RSA2048

	client, err := lego.NewClient(config)
	if err != nil {
		return lego.Client{}, fmt.Errorf("error creating new client: %s", err)
	}

	cfg := tencentcloud.NewDefaultConfig()
	cfg.SecretID = GetSecretID()
	cfg.SecretKey = GetSecretKey()
	provider, err := tencentcloud.NewDNSProviderConfig(cfg)
	fmt.Printf("config: %#v\n", cfg)
	if err != nil {
		return lego.Client{}, fmt.Errorf("error creating DNS provider: %s", err)
	}
	err = client.Challenge.SetDNS01Provider(provider)
	if err != nil {
		return lego.Client{}, fmt.Errorf("error setting DNS provider: %s", err)
	}

	reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		return lego.Client{}, fmt.Errorf("registration error: %v\n", err)
	}

	u.Registration = reg
	c.saveUser(u)
	return *client, nil
}

func (c *Certificate) save() error {
	if err := os.WriteFile(c.DomainPath, []byte(c.Domain), 0600); err != nil {
		return fmt.Errorf("error writing domain file: %v\n", err)
	}
	if err := os.WriteFile(c.PrivateKeyPath, c.PrivateKey, os.ModePerm); err != nil {
		return fmt.Errorf("error creating private key: %s", err)
	}
	if err := os.WriteFile(c.CertificatePath, c.Certificate, os.ModePerm); err != nil {
		return fmt.Errorf("error creating certificate: %s", err)
	}
	if err := os.WriteFile(c.IssuerCertificatePath, c.IssuerCertificate, 0600); err != nil {
		return fmt.Errorf("error creating issuer certificate: %s", err)
	}
	if err := os.WriteFile(c.CSRPath, c.CSR, os.ModePerm); err != nil {
		return fmt.Errorf("error creating CSR: %s", err)
	}
	return nil
}

func (c *Certificate) expires() (int, error) {
	var certDERBlock *pem.Block
	certDERBlock, _ = pem.Decode(c.Certificate)
	if certDERBlock.Type == "CERTIFICATE" {
		cert, err := x509.ParseCertificate(certDERBlock.Bytes)
		if err != nil {
			return -1, err
		}
		timeLeft := cert.NotAfter.Sub(time.Now())
		return int(timeLeft.Hours()), nil
	}
	return -1, fmt.Errorf("invalid certificate")
}

func (c *Certificate) getUser() (*User, error) {
	var user User
	b, err := os.ReadFile(c.UserPath)
	if err == nil {
		err = json.Unmarshal(b, &user)
		if err != nil {
			return nil, err
		}
	} else {
		privateKey := GetPrivateKey()
		user.PrivateKey = privateKey
		user.Email = DefaultContactEmail
	}
	return &user, nil
}

func (c *Certificate) saveUser(user *User) {
	b, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Printf("error marshalling user: %v\n", err)
		return
	}
	err = os.WriteFile(c.UserPath, b, os.ModePerm)
	if err != nil {
		fmt.Printf("error saving user: %v\n", err)
	}
}
