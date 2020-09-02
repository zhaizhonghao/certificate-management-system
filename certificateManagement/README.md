# Certificate Management

This package is used to initial the root CA and issue a certificate for a client.



* Enter the certificateManagement folder to run following command:

  * To show the help

    `go run main.go -h`

    you will get the following hints:

    ```
    Usage of ...:
      -ca-filename string
            file to write CA cert and private key (default "myca")
      -client-filename string
            file to write client cert and private key (default "localhost")
    ```

  * To run

    `go run main.go -ca-filename [ca-filename] -client-filename [client-filename]`

    Notice: If the `client-filename` is not the localhost, please ensure that the assigned `client-filename` is in the **hosts** file (e.g., `/etc/hosts` in the linux).

    or running with the default setting

    `go run main.go`

* To check the certificate in DER
`openssl x509 -inform DER -in ca_cert.der -text`