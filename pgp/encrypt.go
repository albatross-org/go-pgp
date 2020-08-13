package pgp

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/openpgp"

	// Support sha256
	_ "crypto/sha256"

	// Support ripemd160
	_ "golang.org/x/crypto/ripemd160"
)

func Encrypt(entity *openpgp.Entity, message []byte) ([]byte, error) {
	buf := &bytes.Buffer{}

	encryptorWriter, err := openpgp.Encrypt(buf, []*openpgp.Entity{entity}, nil, nil, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("Error creating entity for encryption: %v", err)
	}

	_, err = encryptorWriter.Write(message)
	if err != nil {
		return []byte{}, fmt.Errorf("Error writing message to buffer: %s", err)
	}

	encryptorWriter.Close()
	return buf.Bytes(), nil
}
