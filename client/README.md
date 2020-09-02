# TLS Client

This package implement a TLS client. It can establish the secure connection with the TLS server.

* Enter the client sub-folder to run (must ensure that the server is up before running the client):

  * to show the help: 

    `go run main.go -h`

    you will see the following hints:

    ```
    Usage of ...:
      -cacert string
            Filename containing the ca cert
      -k    Accept/Ingore all server SSL certificates
    ```

    

  * to run

    `go run main.go -cacert [cacertFile] [url]` 

  

As we all known, our operating system has pre-installed some certificates of the trusted CAs. 

If you want to connect to the server with the certificate issued by these CAs, please use the code  which is commented out in default:

```
  //CAs from the system

  rootCAs, _ := x509.SystemCertPool()

  if rootCAs == nil {
  //customer CA
  rootCAs = x509.NewCertPool()
  }
```

Otherwise, we can choice the customer certificate to trust:

```
  //customer CA

  rootCAs := x509.NewCertPool()
```