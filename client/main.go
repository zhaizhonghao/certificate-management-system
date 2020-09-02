package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	var caCertFilename string
	flag.StringVar(&caCertFilename, "cacert", caCertFilename, "Filename containing the ca cert")
	var insecure bool
	flag.BoolVar(&insecure, "k", insecure, "Accept/Ingore all server SSL certificates")
	flag.Parse()

	url := flag.Arg(0)
	// //CAs from the system
	// rootCAs, _ := x509.SystemCertPool()
	// if rootCAs == nil {
	// 	//customer CA
	// 	rootCAs = x509.NewCertPool()
	// }
	//customer CA
	rootCAs := x509.NewCertPool()

	caCertPem, err := ioutil.ReadFile(caCertFilename)
	if err != nil {
		logrus.Fatalf("Failed to append %q to RootCAs:%v", caCertFilename, err)
	}

	//append the cert to the cert pool
	if ok := rootCAs.AppendCertsFromPEM(caCertPem); !ok {
		logrus.Fatalf("unable to append customer CA cert to pool:%v", err)
	}
	//set the configuration for the transport
	config := &tls.Config{
		InsecureSkipVerify: insecure,
		RootCAs:            rootCAs,
	}
	tr := &http.Transport{
		TLSClientConfig: config,
	}
	client := &http.Client{
		Transport: tr,
	}
	res, err := client.Get(url)

	if err != nil {
		logrus.Fatalf("unable to connect to server %v:%v", url, err)
	}

	body := res.Body
	defer body.Close()

	io.Copy(os.Stdout, res.Body)
	fmt.Println()

}
