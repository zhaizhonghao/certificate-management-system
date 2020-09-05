package main

import (
	"flag"

	"github.com/sirupsen/logrus"
	"github.com/zhaizhonghao/certificateDemo/ca"
)

func main() {

	caFile := "myca"
	//clientFile := "localhost"
	csrFile := "localhost"

	flag.StringVar(&caFile, "ca-filename", caFile, "file to write CA cert and private key")
	//flag.StringVar(&clientFile, "client-filename", clientFile, "file to write client cert and private key")
	flag.StringVar(&csrFile, "CSR-filename", csrFile, "file to read CSR")
	flag.Parse()
	//Create a self-sign certificate to initial the Root CA
	err := ca.InitRootCA(caFile)
	if err != nil {
		logrus.Fatalf("unable to init Root CA info:%v", err)
	}
	//Issue a certificate for a client
	err = ca.CreateClientCertFromCSR(csrFile)
	if err != nil {
		logrus.Fatalf("unable to create client certificate:%v", err)
	}

	

}
