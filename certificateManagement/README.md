# Certificate Management

This package is used to initial the root CA and issue a certificate for a client based on the CSR. 

In the future, we plan to add the automatic certificate revocation and certificate modification capabilities.



* Enter the certificateManagement folder to run following command:

  * To show the help

    `go run main.go -h`

    you will get the following hints:

    ```
    Usage of ...:
      -CSR-filename string
            file to read CSR (default "localhost")
      -ca-filename string
            file to write CA cert and private key (default "myca")
    ```

  * To run

    `go run main.go -ca-filename [ca-filename] -CSR-filename [csr-filename]`

    ***Notice***: If the `csr-filename` is not the localhost, please ensure that the assigned `csr-filename` is in the **hosts** file (e.g., `/etc/hosts` in the linux). Besides, you must ensure that the CSR file has generated and been moved under the **certificateManagement** sub-folder.

    or running with the default setting

    `go run main.go`
    
    

* To check the certificate in DER
  `openssl x509 -inform DER -in ca_cert.der -text`

* To check the certificate in PEM

  `openssl x509 -inform PEM -in ca_cert.pem -text`