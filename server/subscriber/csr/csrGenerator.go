package csr

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

var rng = rand.Reader

func GenerateCSR(cn string) error {

	fn := cn + ".pem"
	//Generate the rsa key pair
	key, err := createKey(fn)
	if err != nil {
		return fmt.Errorf("unable to create private key for client %v", err)
	}
	template := &x509.CertificateRequest{
		SignatureAlgorithm: x509.SHA256WithRSA,
		PublicKeyAlgorithm: x509.RSA,
		PublicKey:          &key.PublicKey,
		Subject:            clientInfo(cn),
		DNSNames:           []string{cn},
		EmailAddresses:     []string{"1016041104@njupt.edu.cn"},
	}

	csrDER, err := x509.CreateCertificateRequest(rng, template, key)
	if err != nil {

		return fmt.Errorf("unable to generate the CSR:%v", err)
	}
	pemEncode := func(b []byte, t string) []byte {
		return pem.EncodeToMemory(&pem.Block{Bytes: b, Type: t})
	}

	csrPEM := pemEncode(csrDER, "CERTIFICATE REQUEST")

	csrFile := cn + ".csr"
	if err = ioutil.WriteFile(csrFile, csrPEM, 0644); err != nil {
		log.Fatal(err)
		return fmt.Errorf("unable to generate the CSR:%v", err)
	}

	return nil
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
