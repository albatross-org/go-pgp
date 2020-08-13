package pgp

import (
	"bytes"
	_ "crypto/sha256"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/openpgp"
	_ "golang.org/x/crypto/ripemd160"
)

func Decrypt(entity *openpgp.Entity, encrypted []byte) ([]byte, error) {
	// New code, load the encrypted buffer without decoding any armour.
	block := bytes.NewReader(encrypted)

	// Decrypt message
	entityList := openpgp.EntityList{entity}
	messageReader, err := openpgp.ReadMessage(block, entityList, nil, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("Error reading message: %v", err)
	}
	read, err := ioutil.ReadAll(messageReader.UnverifiedBody)
	if err != nil {
		return []byte{}, fmt.Errorf("Error reading unverified body: %v", err)
	}

	// Return output - an decrypted message.
	return read, nil
}
