# TLS Server

We simply implement a TLS server to listen to the request from the TLS client.

* Enter the **server** sub-folder to

  *  To show the help

    `go run main.go -h`

    you will get the following hints

    ```
    Usage of ...:
      -crt string
            Filename to containing my certificate in PEM format
      -key string
            Filename to containing my private key in PEM format
    ```

    

  * To run 

    `go run main.go -crt [certFlie] -key [keyFile]`

    The server will be listening on the `localhost:8080` , you can modify the source code to change the ip and port.

