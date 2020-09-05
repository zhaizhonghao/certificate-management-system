package ca

import (
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"time"
)

func CreateClientCertFromCSR(csrn string) (string, error) {
	csrFile := csrn + ".csr"
	data, err := ioutil.ReadFile(csrFile)
	if err != nil {
		return "", fmt.Errorf("unable to read the csr file:%v", err)
	}
	b, _ := pem.Decode(data)
	var csr *x509.CertificateRequest
	if b == nil {
		csr, err = x509.ParseCertificateRequest(data)
	} else {
		csr, err = x509.ParseCertificateRequest(b.Bytes)
	}
	if err != nil {
		return "", fmt.Errorf("unable to parse the csr file:%v", err)
	}

	err = csr.CheckSignature()
	if err != nil {
		return "", fmt.Errorf("The signature of csr is invalid:%v", err)
	}

	client := &x509.Certificate{}
	//the serialNumber is required
	client.SerialNumber = big.NewInt(time.Now().Unix())

	now := time.Now()

	//To set the valid period
	client.NotBefore = now
	//AddDate(years int,months int,days int)
	client.NotAfter = now.AddDate(1, 0, 0)

	//To set the information of  parent *Certificate(https://golang.org/pkg/crypto/x509/#Certificate)
	client.Subject = csr.Subject

	client.SubjectKeyId = getKeyID(csr.PublicKey.(*rsa.PublicKey))
	//The intent for the cerificate
	client.KeyUsage = x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment
	client.ExtKeyUsage = []x509.ExtKeyUsage{
		x509.ExtKeyUsageServerAuth,
		x509.ExtKeyUsageAny,
		x509.ExtKeyUsageEmailProtection,
	}

	client.DNSNames = csr.DNSNames
	client.EmailAddresses = csr.EmailAddresses
	fmt.Println(client)
	//func CreateCertificate(rand io.Reader, template, parent *Certificate, pub, priv interface{}) (cert []byte, err error)
	//CreateCertificate creates a new X.509v3 certificate(DER) based on a template
	//since it is a self-signed certificate, the subject is the issuer
	der, err := x509.CreateCertificate(rng, client, rootCAInfo.cert, csr.PublicKey.(*rsa.PublicKey), rootCAInfo.key)
	if err != nil {
		return "", fmt.Errorf("unable to create cert:%v", err)
	}

	fn := csr.Subject.CommonName + ".pem"
	err = exportCert(fn, der)
	if err != nil {
		return "", err
	}

	return fn, nil
}

func initClientCert(cn string, key *rsa.PrivateKey) error {
	err := createSignedCert(cn, &key.PublicKey)
	if err != nil {
		return fmt.Errorf("unable to create cert for client %v:%v", cn, err)
	}

	return err
}

func createSignedCert(cn string, clientKey *rsa.PublicKey) error {
	client := &x509.Certificate{}
	//the serialNumber is required
	client.SerialNumber = big.NewInt(time.Now().Unix())

	now := time.Now()

	//To set the valid period
	client.NotBefore = now
	//AddDate(years int,months int,days int)
	client.NotAfter = now.AddDate(1, 0, 0)

	//To set the information of  parent *Certificate(https://golang.org/pkg/crypto/x509/#Certificate)
	client.Subject = clientInfo(cn)

	client.SubjectKeyId = getKeyID(clientKey)
	//The intent for the cerificate
	client.KeyUsage = x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment
	client.ExtKeyUsage = []x509.ExtKeyUsage{
		x509.ExtKeyUsageServerAuth,
		x509.ExtKeyUsageAny,
		x509.ExtKeyUsageEmailProtection,
	}

	client.DNSNames = []string{cn}
	client.EmailAddresses = []string{"1016041104@njupt.edu.cn"}

	//func CreateCertificate(rand io.Reader, template, parent *Certificate, pub, priv interface{}) (cert []byte, err error)
	//CreateCertificate creates a new X.509v3 certificate(DER) based on a template
	//since it is a self-signed certificate, the subject is the issuer
	der, err := x509.CreateCertificate(rng, client, rootCAInfo.cert, clientKey, rootCAInfo.key)
	if err != nil {
		return fmt.Errorf("unable to create cert:%v", err)
	}

	fn := cn + ".pem"
	err = exportCert(fn, der)
	if err != nil {
		return err
	}

	return err
}

func clientInfo(cn string) pkix.Name {
	name := pkix.Name{
		CommonName:         cn,
		Country:            []string{"CN"},
		Organization:       []string{"Nanjing University of Posts and Telecommunications"},
		OrganizationalUnit: []string{"IT"},
		Province:           []string{"Jiangsu"},
		Locality:           []string{"Nanjing"},
		StreetAddress:      []string{"Xin Mo Fan road NO.66"},
	}
	return name
}
