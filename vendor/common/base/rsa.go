// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package base

import (
	rsarand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

const DOC_URL = "https://github.com/gogits/go-gogs-client/wiki"

type (
	TplName string
)

func GetRSAPublickey() []byte {
	return []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC9HbrefkoBbiov70eFTpvc/2Zd
OtCQcQyGN4zneHXzDjAZxZpqWEJSSePs4ExUEORHRA/V513+qVryMKSMmYfxxKNi
z6ZQOk1d9eQA0HNtDYpeZGykXxMy+qdwYmDrohVKkrLaCE0m4o1eqsVESxrIYiig
+SJLQMB62TIo7Z70zwIDAQAB
-----END PUBLIC KEY-----
`)
}

func GetRSAPrivatekey() []byte {
	return []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQC9HbrefkoBbiov70eFTpvc/2ZdOtCQcQyGN4zneHXzDjAZxZpq
WEJSSePs4ExUEORHRA/V513+qVryMKSMmYfxxKNiz6ZQOk1d9eQA0HNtDYpeZGyk
XxMy+qdwYmDrohVKkrLaCE0m4o1eqsVESxrIYiig+SJLQMB62TIo7Z70zwIDAQAB
AoGAJrARTXjqbNZ7nOjHadcF/bTl3QauKg5mcIKmuaekAYASiQMqrry1fJ4PGaNd
GgmGmamrz6eQTAubRlZ7IyQ5cfqKFDFh/cq9AC0MkWG9UOcerxNJXg1Tpzpb0pLw
PQ+b63rskhJymmc3ByONcEg34wXTl5OtuJF+D8Ftq8WRnLECQQDDHuwzKhWZX6PA
Kbvqi0u+zl2M2FpvQtPxPMeIwGGrBnqD7bQpZsZ6F+d4CFOmw7FyG8SQk/W0YqRv
bCl/T7BHAkEA+B8w8hn8Bd26UtA9GgV8NUoQCrmXz1viWXW3cV/sUuI5sQPfmkLp
UbCh+ldItly3xrLrDHoDTeCXi3KvRIEjOQJATd+NjW4CaNAO3qbJZPZrKJ/cHlZK
4ZTeWa1URXPihwty4iyAdvWZySi5LOLF4AzCSTRj4v/qVC/6SK32ceUwCQJBAIH8
5UmQr2XrWZfVHI2rXf0VBf54aL3rp1Oyxh4RYN+zQQIpw1UvxMhVPybF34QaYvUn
+tgYe+6qwPn/ZS8AcJECQQC88Jp3iIWM/3aqwBJLB0ogIrUZR/NGC+Is40Q9ONFM
JAd3jqYMpx7l92Llb9qphX2AIcigBrsu+7Y2aP/lHXgM
-----END RSA PRIVATE KEY-----
`)
}

// RSA加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(GetRSAPublickey())
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rsarand.Reader, pub, origData)
}

// RSA解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(GetRSAPrivatekey())
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rsarand.Reader, priv, ciphertext)
}
