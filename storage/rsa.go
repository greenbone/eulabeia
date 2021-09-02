package storage

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/greenbone/eulabeia/config"
)

type Crypt interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}

// RSA is used to en- and decrypt using RSA.
type RSA struct {
	prvKey *rsa.PrivateKey
}

func (r RSA) Encrypt(b []byte) ([]byte, error) {
	return rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&r.prvKey.PublicKey,
		b,
		nil)
}

func (r RSA) Decrypt(b []byte) ([]byte, error) {
	return r.prvKey.Decrypt(nil, b, &rsa.OAEPOptions{Hash: crypto.SHA256})
}

func exportRSAPrivateKey(privkey *rsa.PrivateKey) []byte {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return privkey_pem
}

func parseRSAPrivateKey(b []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(b)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func NewRSACrypt(c config.Configuration) (Crypt, error) {
	if c.Director.KeyFile == "" {
		return nil, nil
	}
	k, err := ioutil.ReadFile(c.Director.KeyFile)
	var prvKey *rsa.PrivateKey
	if err != nil {
		dir := filepath.Dir(c.Director.KeyFile)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0700)
			if err != nil {
				return nil, err
			}
		} else if err != nil {
			return nil, err
		}
		prvKey, err = rsa.GenerateKey(rand.Reader, 4096)
		if err != nil {
			return nil, err
		}
		k = exportRSAPrivateKey(prvKey)
		if err := ioutil.WriteFile(c.Director.KeyFile, k, 0600); err != nil {
			return nil, err
		}
	}
	if prvKey, err = parseRSAPrivateKey(k); err != nil {
		return nil, err
	}

	return RSA{prvKey: prvKey}, nil
}
