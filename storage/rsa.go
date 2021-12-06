// Copyright (C) 2021 Greenbone Networks GmbH
//
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
	hash := sha256.New()
	msgLen := len(b)
	step := r.prvKey.PublicKey.Size() - 2*hash.Size() - 2
	var encryptedBytes []byte

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}

		encrypedChunk, err := rsa.EncryptOAEP(
			hash,
			rand.Reader,
			&r.prvKey.PublicKey,
			b[start:finish],
			nil,
		)
		if err != nil {
			return nil, err
		}
		encryptedBytes = append(encryptedBytes, encrypedChunk...)
	}
	return encryptedBytes, nil
}

func (r RSA) Decrypt(b []byte) ([]byte, error) {
	msgLen := len(b)
	step := r.prvKey.PublicKey.Size()
	var decryptedBytes []byte

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}

		decryptedBlockBytes, err := r.prvKey.Decrypt(
			nil,
			b[start:finish],
			&rsa.OAEPOptions{Hash: crypto.SHA256},
		)
		if err != nil {
			return nil, err
		}

		decryptedBytes = append(decryptedBytes, decryptedBlockBytes...)
	}

	return decryptedBytes, nil
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
