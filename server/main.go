package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	addr = "localhost:8080"
)

func main() {
	var myCertFn, myKeyFn string
	flag.StringVar(&myCertFn, "crt", myCertFn, "Filename to containing my certificate in PEM format")
	flag.StringVar(&myKeyFn, "key", myKeyFn, "Filename to containing my private key in PEM format")
	flag.Parse()

	fmt.Printf("Listening on https://%v\n", addr)
	//
	http.HandleFunc("/", indexHandler)
	//func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
	err := http.ListenAndServeTLS(addr, myCertFn, myKeyFn, nil)

	if err != nil {
		logrus.Fatalf("unable to create server at %v:%v", addr, err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this a secure connection\n")
}
