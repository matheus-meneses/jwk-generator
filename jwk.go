package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"gopkg.in/square/go-jose.v2"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := ioutil.ReadFile("/Users/bi001189/Downloads/brseal.key")
	decode, _ := pem.Decode(file)

	cert, err := x509.ParsePKCS8PrivateKey(decode.Bytes)
	if err != nil {
		println(err.Error())
		os.Exit(0)
	}
	key := jose.JSONWebKey{
		Key:       cert.(*rsa.PrivateKey),
		Algorithm: "PS256",
	}
	json, err := key.MarshalJSON()
	if err != nil {
		println(err.Error())
		os.Exit(0)
	}
	println(string(json))
}
