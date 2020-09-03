package main

import (
	"fmt"

	"github.com/zhaizhonghao/CertificateManagementSystem/server/subscriber/csr"
)

func main() {
	cn := "localhost"
	err := csr.GenerateCSR(cn)
	if err != nil {
		fmt.Printf("Unable to generate the csr:%v\n", err)
	}
}
