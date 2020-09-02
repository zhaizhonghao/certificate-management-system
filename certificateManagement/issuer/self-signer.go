package issuer

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"

	"github.com/sirupsen/logrus"
)

//Reader(var Reader io.Reader) is a global, shared instance of a cryptographically secure random number generator.
var (
	rng           = rand.Reader
	issuerPrivKey = createKey()
)

//create the self-signed certificate
func NewSelfSignedCert() ([]byte, error) {
	issuer := &x509.Certificate{}
	//the serialNumber is required
	issuer.SerialNumber = big.NewInt(time.Now().Unix())

	now := time.Now()

	//To set the valid period
	issuer.NotBefore = now
	//AddDate(years int,months int,days int)
	issuer.NotAfter = now.AddDate(1, 0, 0)

	//To set the information of  parent *Certificate(https://golang.org/pkg/crypto/x509/#Certificate)
	issuer.Subject = issuerInfo()

	issuer.DNSNames = []string{"zhaizhonghao.njupt.edu.cn"}
	issuer.EmailAddresses = []string{"1016041104@njupt.edu.cn"}

	//func CreateCertificate(rand io.Reader, template, parent *Certificate, pub, priv interface{}) (cert []byte, err error)
	//CreateCertificate creates a new X.509v3 certificate based on a template
	//since it is a self-signed certificate, the subject is the issuer
	cert, err := x509.CreateCertificate(rng, issuer, issuer, &issuerPrivKey.PublicKey, issuerPrivKey)

	return cert, err
}

func issuerInfo() pkix.Name {
	name := pkix.Name{
		CommonName:         "Zhai Trust",
		Country:            []string{"CN"},
		Organization:       []string{"Nanjing University of Posts and Telecommunications"},
		OrganizationalUnit: []string{"IT"},
		Province:           []string{"Jiangsu"},
		Locality:           []string{"Nanjing"},
		StreetAddress:      []string{"Xin Mo Fan road NO.66"},
	}
	return name
}

//create a rsa key pair
func createKey() *rsa.PrivateKey {
	const keySize = 1024 * 2
	//GenerateKey generates an RSA keypair of the given bit size using the random source random (for example, crypto/rand.Reader).
	k, err := rsa.GenerateKey(rng, keySize)
	if err != nil {
		logrus.Errorf("Unable to create host Key:%v", err)
	}
	return k
}
