# Certificate Management System

Public-Key Infrastructure (PKI) is a technology which enables the secure communication over the trustless Internet. Its main function is attesting to the binding between the public key of an entity and its authentic identity (e.g., a domain name). The binding is established through a process of registration and issuance of certificates at and by a certificate authority (CA). Hyper Text Transfer Protocol over Secure Socket Layer (HTTPS) depending on PKI can effectively resist man-in-the-middle attack and has been adopted worldwide in various web applications such as e-mail, e-commerce and e-banking. For more fundamental details on the PKI see [Bulletproof SSL and TLS](https://www.feistyduck.com/books/bulletproof-ssl-and-tls/). 



We implement a  X.509 v3 certificate management system using the Golang official package [crypto](https://golang.org/pkg/crypto/). Our system has following capabilities:

* Certificate Sign Request (CSR) generation
  * Generating the CSR 

* [Key generation](certificateManagement/key-id.go) 
  * Generating the RSA based key pair 

* Certificate management

  *  Creating a self-signed certificate for the Root CA

  *  Issuing a signed certificate for the requested client

  *  Revoking an existing certificate

  *  Mudifying an existing certificate
  *  Verifying an existing certificate

* Test
  * Testing the secure connection between a TLS server and a TLS client



