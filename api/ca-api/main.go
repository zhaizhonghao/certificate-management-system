package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/zhaizhonghao/CertificateManagementSystem/api/ca-api/ca"
)

//The success information for return
type Success struct {
	Payload string `json:"Payload"`
	Message string `json:"message"`
}

//The error information for the return
type Error struct {
	Message string `json:"success"`
}

func main() {
	caFile := "myca"

	//Create a self-sign certificate to initial the Root CA
	err := ca.InitRootCA(caFile)
	if err != nil {
		logrus.Fatalf("unable to init Root CA info:%v", err)
	}
	router := mux.NewRouter()

	router.HandleFunc("/certificate", requestCertificate).Methods("POST")

	router.HandleFunc("/certificate", revokeCertificate).Methods("DELETE")

	router.HandleFunc("/certificate", updateCertificate).Methods("PUT")

	router.HandleFunc("/certificate", patchCertificate).Methods("PATCH")

	http.ListenAndServe(":5000", router)

}

func requestCertificate(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		failure := Error{
			Message: fmt.Sprintf("Unable to parse the form:%v", err),
		}
		json.NewEncoder(w).Encode(failure)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-csr-files", "upload-*.csr")

	if err != nil {
		failure := Error{
			Message: fmt.Sprintf("Unable to write the csr file:%v", err),
		}
		json.NewEncoder(w).Encode(failure)
		return
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		failure := Error{
			Message: fmt.Sprintf("Unable to read the csr file:%v", err),
		}
		json.NewEncoder(w).Encode(failure)
		return
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	//Generate the certificate based on the CSR

	csrFile := strings.Split(tempFile.Name(), ".")
	//csrn := strings.Split(csrFile[0], "\\")

	fmt.Println("create the certificate!", csrFile[0])
	fn := ""
	fn, err = ca.CreateClientCertFromCSR(csrFile[0])
	if err != nil {
		failure := Error{
			Message: fmt.Sprintf("Fail to create the certificate:%v", err),
		}
		json.NewEncoder(w).Encode(failure)
		return
	}
	cert := []byte{}
	cert, err = ioutil.ReadFile(fn)
	if err != nil {
		failure := Error{
			Message: fmt.Sprintf("Unable to read the cert file:%v", err),
		}
		json.NewEncoder(w).Encode(failure)
		return
	}

	success := Success{
		Payload: string(cert),
		Message: "Apply the certificate successfully",
	}

	json.NewEncoder(w).Encode(success)

}

func revokeCertificate(w http.ResponseWriter, r *http.Request) {

}

func updateCertificate(w http.ResponseWriter, r *http.Request) {

}

func patchCertificate(w http.ResponseWriter, r *http.Request) {

}
