// Package cert
// Date       : 2024/8/13 19:44
// Author     : Amu
// Description:
package cert

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

var _ ICertificate = (*LocalCertificate)(nil)

type LocalCertificate struct {
	Domains         []string `json:"domains"`
	PrivateKey      []byte   `json:"private_key"`
	Certificate     []byte   `json:"certificate"`
	PrivateKeyPath  string   `json:"private_key_path"`
	CertificatePath string   `json:"certificate_path"`
	RenewBefore     int
}

func NewLocalCertificate(config *Config) *LocalCertificate {
	cert := &LocalCertificate{
		PrivateKeyPath:  filepath.Join(config.CacheDir, LocalPrivateKeyFileName),
		CertificatePath: filepath.Join(config.CacheDir, LocalCertificateFileName),
		Domains:         config.Domains,
		RenewBefore:     config.RenewBefore,
	}
	return cert
}

func (l *LocalCertificate) Generate() error {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	// 创建证书模板
	tpl := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Country:            []string{"CN"},
			Province:           []string{"BJ"},
			Locality:           []string{"BJ"},
			Organization:       []string{"Amuluze"},
			OrganizationalUnit: []string{"AmuluzeProxy"},
			CommonName:         "Amuluze Root CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		BasicConstraintsValid: true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	// 创建证书
	derCert, err := x509.CreateCertificate(rand.Reader, &tpl, &tpl, &privateKey.PublicKey, privateKey)
	if err != nil {
		return fmt.Errorf("x509.CreateCertificate error: %s", err)
	}

	// 证书编码
	buf := &bytes.Buffer{}
	err = pem.Encode(buf, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derCert,
	})
	if err != nil {
		return fmt.Errorf("certificate pem.Encode error: %s", err)
	}

	pemCert := buf.Bytes()
	l.Certificate = pemCert
	// 私钥编码
	buf = &bytes.Buffer{}
	err = pem.Encode(buf, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err != nil {
		return fmt.Errorf("key pem.Encode error: %s", err)
	}
	pemKey := buf.Bytes()
	l.PrivateKey = pemKey

	if err = l.save(); err != nil {
		return err
	}
	return nil
}

func (l *LocalCertificate) Load() error {
	if _, err := FileExists(l.PrivateKeyPath); err != nil {
		return err
	}
	keyData, err := os.ReadFile(l.PrivateKeyPath)
	if err != nil {
		return err
	}
	l.PrivateKey = keyData

	if _, err = FileExists(l.CertificatePath); err != nil {
		return err
	}
	certData, err := os.ReadFile(l.CertificatePath)
	if err != nil {
		return err
	}
	l.Certificate = certData
	return nil
}

func (l *LocalCertificate) save() error {
	if err := os.WriteFile(l.PrivateKeyPath, l.PrivateKey, os.ModePerm); err != nil {
		return fmt.Errorf("error creating private key: %s", err)
	}
	if err := os.WriteFile(l.CertificatePath, l.Certificate, os.ModePerm); err != nil {
		return fmt.Errorf("error creating certificate: %s", err)
	}
	return nil
}

func (l *LocalCertificate) Renew() error {
	expire, err := l.expires()
	if err != nil {
		return err
	}
	if l.RenewBefore > expire {
		return l.Generate()
	}
	return nil
}

type LocalCertificateStatus struct {
	Expires int `json:"expires"`
}

func (l *LocalCertificate) expires() (int, error) {
	var certDERBlock *pem.Block
	certDERBlock, _ = pem.Decode(l.Certificate)
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
