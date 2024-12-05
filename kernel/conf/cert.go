package conf

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"kernel/common"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

func InitCertificate() (string, string) {
	if Conf.CertDNSNames == nil || len(Conf.CertDNSNames) == 0 {
		Conf.CertDNSNames = []string{"localhost"}
	}
	certFileName := filepath.Join(CertDir, Conf.CertFileName)
	certKeyFileName := filepath.Join(CertDir, Conf.CertKeyFileName)
	if common.File.IsExist(certFileName) && common.File.IsExist(certKeyFileName) {
		return certFileName, certKeyFileName
	}

	GenerateCertificate(Conf.CertDNSNames, certFileName, certKeyFileName)
	common.Log.Info("Generated self-signed certificate and key [%s] [%s]", certFileName, certKeyFileName)
	return certFileName, certKeyFileName
}

// GenerateCertificate Generate self signed certificate
func GenerateCertificate(DNSNames []string, certFileName string, certKeyFileName string) {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		common.Log.Fatal(common.ExitCodeCertificateErr, "Failed to generate private key: %v", err)
	}

	// 创建证书模板
	n, _ := rand.Int(rand.Reader, big.NewInt(4096))
	template := x509.Certificate{
		SerialNumber: n,
		Subject: pkix.Name{
			Organization: []string{Name},
			CommonName:   DNSNames[0],
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // 有效期为10年
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              DNSNames, // 添加需要的域名
	}

	// 自签名证书
	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		common.Log.Fatal(common.ExitCodeCertificateErr, "Failed to create certificate: %v", err)
	}

	// 将证书保存到文件
	certOut, err := os.Create(certFileName)
	if err != nil {
		common.Log.Fatal(common.ExitCodeCertificateErr, "Failed to open cert.pem for writing: %v", err)
	}
	defer certOut.Close()

	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes}); err != nil {
		common.Log.Fatal(common.ExitCodeCertificateErr, "Failed to write data to cert.pem: %v", err)
	}

	// 将私钥保存到文件
	keyOut, err := os.OpenFile(certKeyFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		common.Log.Fatal(common.ExitCodeCertificateErr, "Failed to open key.pem for writing: %v", err)
	}
	defer keyOut.Close()

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	if err := pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKeyBytes}); err != nil {
		common.Log.Fatal(common.ExitCodeCertificateErr, "Failed to write data to key.pem: %v", err)
	}
}
