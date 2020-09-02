# Certificate Management
* Enter the certificateManagement folder to run following command:

  * To show the help :

     `go run main.go -h`

    you will get the follow hints:

    ```
    Usage of ...:
      -ca-filename string
            file to write CA cert and private key (default "myca")
      -client-filename string
            file to write client cert and private key (default "localhost")
    ```

    ***Notice***: If the `client-filename` is not localhost, you must ensure that the assigned `client-filename` is in your **hosts** file (e.g., `/etc/hosts` in linux).

  * To initial the root CA and issue a certificate for a client:

    `go run main.go -ca-filename [ca-filename] -client-filename [client-filename]` 

    or using the default setting:

    `go run main.go`

    

* To check the certificate in DER (ensuring that the openssl has been installed)
`openssl x509 -inform DER -in [certFile] -text`