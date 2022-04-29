package encrypt

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

func LoadKey(file string) (interface{}, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	} else if block.Type == "RSA PRIVATE KEY" {
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	} else if block.Type == "RSA PUBLIC KEY" {
		return x509.ParsePKCS1PublicKey(block.Bytes)
	} else if block.Type == "PUBLIC KEY" {
		return x509.ParsePKIXPublicKey(block.Bytes)
	} else if block.Type == "PRIVATE KEY" {
		return x509.ParsePKCS8PrivateKey(block.Bytes)
	} else {
		return nil, errors.New("PEM format error, must be pkcs#1 pkcs#8 or x.509 pubkey")
	}
}
