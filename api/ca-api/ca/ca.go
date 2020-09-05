package ca

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type certInfo struct {
	key  *rsa.PrivateKey
	cert *x509.Certificate
}

var (
	rootCAInfo *certInfo
	rng        = rand.Reader
	issuerName = "Zhai Trust"
)

func InitRootCA(fn string) error {
	rootCAInfo = &certInfo{}
	err := rootCAInfo.initKey(fn)
	if err != nil {
		return fmt.Errorf("unable to init root CA private key:%v", err)
	}

	err = rootCAInfo.initRootCert(fn)
	if err != nil {
		return fmt.Errorf("unable to init root CA certificate:%v", err)
	}
	return nil
}

func (ci *certInfo) initKey(fn string) error {
	fn += "-key.pem"
	//try to read private key for Root CA from file (PEM format)
	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		logrus.Infof("No existing key found in %v for root ca, creating", fn)
		//if we can't read the file, assume first run and generate a new private key
		ci.key, err = createKey(fn)
		return err
	}

	//we have a key file, so decode PEM bytes to rsa.PrivateKey
	b, _ := pem.Decode(buf)

	ci.key, err = x509.ParsePKCS1PrivateKey(b.Bytes)
	return err
}

func (ci *certInfo) initRootCert(fn string) error {
	fn += ".pem"

	var buf []byte
	var err error
	var der []byte

	buf, err = ioutil.ReadFile(fn)
	if err != nil {
		logrus.Infof("no cert found in %v for root CA, creating", fn)

		der, err = createSelfSignedCert(fn, ci.key)
		if err != nil {
			return fmt.Errorf("unable to create root CA cert:%v", err)
		}
	} else {
		b, _ := pem.Decode(buf)
		der = b.Bytes
	}
	ci.cert, err = x509.ParseCertificate(der)
	return err
}

func createSelfSignedCert(fn string, key *rsa.PrivateKey) ([]byte, error) {
	issuer := &x509.Certificate{}
	//the serialNumber is required
	issuer.SerialNumber = big.NewInt(time.Now().Unix())

	now := time.Now()

	//To set the valid period
	issuer.NotBefore = now
	//AddDate(years int,months int,days int)
	issuer.NotAfter = now.AddDate(2, 0, 0)

	//To set the information of  parent *Certificate(https://golang.org/pkg/crypto/x509/#Certificate)
	issuer.Subject = issuerInfo()
	//SHA-1 checksum of the its public key
	issuer.SubjectKeyId = getKeyID(&key.PublicKey)
	//SHA-1 checksum of its Issuer's public key
	issuer.AuthorityKeyId = getKeyID(&key.PublicKey)
	//The intent for the certificate
	issuer.KeyUsage = x509.KeyUsageCertSign
	//BasicConstrainsValid indicates whether IsCA, MaxPathLen, and MaxPathLenZero are valid
	issuer.BasicConstraintsValid = true
	issuer.IsCA = true
	issuer.MaxPathLenZero = true

	issuer.DNSNames = []string{"zhaizhonghao.njupt.edu.cn"}
	issuer.EmailAddresses = []string{"1016041104@njupt.edu.cn"}

	//func CreateCertificate(rand io.Reader, template, parent *Certificate, pub, priv interface{}) (cert []byte, err error)
	//CreateCertificate creates a new X.509v3 certificate(DER) based on a template
	//since it is a self-signed certificate, the subject is the issuer
	der, err := x509.CreateCertificate(rng, issuer, issuer, &key.PublicKey, key)
	if err != nil {
		return nil, fmt.Errorf("unable to create cert:%v", err)
	}
	err = exportCert(fn, der)
	if err != nil {
		return nil, err
	}

	return der, err
}

func exportCert(fn string, der []byte) error {
	b := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: der,
	}
	file, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("unable to create file %v:%v", fn, err)
	}
	defer file.Close()

	err = pem.Encode(file, b)
	if err != nil {
		return fmt.Errorf("unable to encode certifiate in PEM format:%v", err)
	}
	return nil
}

func issuerInfo() pkix.Name {
	name := pkix.Name{
		CommonName:         issuerName,
		Country:            []string{"CN"},
		Organization:       []string{"Nanjing University of Posts and Telecommunications"},
		OrganizationalUnit: []string{"IT"},
		Province:           []string{"Jiangsu"},
		Locality:           []string{"Nanjing"},
		StreetAddress:      []string{"Xin Mo Fan road NO.66"},
	}
	return name
}

//createKey creates a rsa.PrivatKey and saves it to a file in before returning the key
func createKey(fn string) (*rsa.PrivateKey, error) {
	const keySize = 1024 * 2
	k, err := rsa.GenerateKey(rng, keySize)
	if err != nil {
		return nil, fmt.Errorf("unable to create private key for %v:%v", fn, err)
	}
	//Translate the key to pem's block
	b := &pem.Block{
		Bytes: x509.MarshalPKCS1PrivateKey(k),
		Type:  "RSA PRIVATE KEY",
	}
	buf := &bytes.Buffer{}
	err = pem.Encode(buf, b)
	if err != nil {
		return nil, fmt.Errorf("Unable to encode private key to PEM:%v", err)
	}
	err = ioutil.WriteFile(fn, buf.Bytes(), 0600)
	return k, err
}
