package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"github.com/zhaizhonghao/CertificateManagementSystem/server/subscriber/csr"
)

func main() {
	cn := "localhost"
	err := csr.GenerateCSR(cn)
	if err != nil {
		fmt.Printf("Unable to generate the csr:%v\n", err)
	}

	csrFile := cn + ".csr"
	data, err := ioutil.ReadFile(csrFile)
	if err != nil {
		fmt.Println("unable to read the csr file")
	}
	b, _ := pem.Decode(data)
	var csr *x509.CertificateRequest
	if b == nil {
		csr, err = x509.ParseCertificateRequest(data)
	} else {
		csr, err = x509.ParseCertificateRequest(b.Bytes)
	}
	if err != nil {
		fmt.Println("unable to parse the csr file")
	}
	fmt.Println(csr.Subject.CommonName)
	fmt.Println(csr.Subject)
	fmt.Println(csr.PublicKey)
	err = csr.CheckSignature()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("the signature of the csr is valid")
	}
}
