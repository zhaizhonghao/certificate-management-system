package main

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/zhaizhonghao/certificateDemo/ca"
	"github.com/zhaizhonghao/certificateDemo/issuer"
)

var (
	caCertOutFileName string
	formatPem         = false
	caCertInFileName  string
)

func main() {
	// //StringVar defines a string flag with specified name, default value, and usage string. The argument p points to a string variable in which to store the value of the flag.
	// flag.StringVar(&caCertOutFileName, "out", caCertOutFileName, "Out cert filename")
	// //BoolVar defines a bool flag with specified name, default value, and usage string. The argument p points to a bool variable in which to store the value of the flag.
	// flag.BoolVar(&formatPem, "pem", formatPem, "Write cert in PEM(text) format, default is DER(binary)")
	// flag.StringVar(&caCertInFileName, "in", caCertInFileName, "In cert filename")
	// flag.Parse()

	// if len(caCertOutFileName) == 0 && len(caCertInFileName) == 0 {
	// 	logrus.Fatalf("At least 'in' or 'out' MUST be specified")
	// }

	// if len(caCertOutFileName) > 0 && len(caCertInFileName) > 0 {
	// 	logrus.Fatalf("At least 'in' or 'out', but not both")
	// }

	// if len(caCertOutFileName) > 0 {
	// 	createCert()
	// 	return
	// }

	// readCert()

	caFile := "myca"
	flag.StringVar(&caFile, "ca-filename", caFile, "file to write CA cert and private key")
	flag.Parse()

	err := ca.InitRootCA(caFile)
	if err != nil {
		logrus.Fatalf("unable to init Root CA info:%v", err)
	}

	err = ca.CreateClientCA("localhost")
	if err != nil {
		logrus.Fatalf("unable to create client certificate:%v", err)
	}

}

func readCert() {
	buf, err := ioutil.ReadFile(caCertInFileName)
	if err != nil {
		logrus.Fatalf("Unable to read file:%v", err)
	}
	//try to decode bytes as PEM
	block, _ := pem.Decode(buf)

	var cert *x509.Certificate
	if block != nil {
		cert, err = x509.ParseCertificate(block.Bytes)
	} else {
		cert, err = x509.ParseCertificate(buf)
	}

	if err != nil {
		logrus.Fatalf("Failed to parse certificate:%v", err)
	}

	fmt.Println("Successfully parsed certificate.\n", cert)

}

func createCert() {
	cert, err := issuer.NewSelfSignedCert()
	if err != nil {
		logrus.Fatalf("Unable to create certificate:%v", err)
	}

	var buf []byte
	if !formatPem {
		caCertOutFileName += ".der"
		buf = cert
	} else {
		caCertOutFileName += ".pem"
		//To translate the bytes into PEM format
		//https://golang.org/pkg/encoding/pem/
		//https://golang.org/pkg/encoding/pem/#Block
		block := &pem.Block{
			Type: "CERTIFICATE",
			Headers: map[string]string{
				"Created by": os.Args[0],
			},
			Bytes: cert,
		}
		//func Encode(out io.Writer, b *Block) error
		//Encode writes the PEM encoding of b to out.
		b := &bytes.Buffer{}
		err = pem.Encode(b, block)
		if err != nil {
			logrus.Fatalf("Unable to encode certificate in PEM:%v", err)
		}
		buf = b.Bytes()

	}
	//func WriteFile(filename string, data []byte, perm os.FileMode) error
	//WriteFile writes data to a file named by filename. If the file does not exist, WriteFile creates it with permissions perm (before umask); otherwise WriteFile truncates it before writing, without changing permissions.
	err = ioutil.WriteFile(caCertOutFileName, buf, 0644)
	if err != nil {
		logrus.Fatalf("Unable to write certificate:%v", err)
	}

}
