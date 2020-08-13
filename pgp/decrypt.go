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
	// Decode message -- this is part of https://github.com/jchavannes/go-pgp/
	// block, err := armor.Decode(bytes.NewReader(encrypted))
	// if err != nil {
	// 	return []byte{}, fmt.Errorf("Error decoding: %v", err)
	// }
	// if block.Type != "Message" {
	// 	return []byte{}, errors.New("Invalid message type")
	// }

	// New code, load the encrypted buffer without decoding any armour.
	block := bytes.NewBuffer(encrypted)

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

	// Uncompress message -- this is part of https://github.com/jchavannes/go-pgp/
	// reader := bytes.NewReader(read)
	// uncompressed, err := gzip.NewReader(reader)
	// if err != nil {
	// 	return []byte{}, fmt.Errorf("Error initializing gzip reader: %v", err)
	// }
	// defer uncompressed.Close()

	// out, err := ioutil.ReadAll(uncompressed)
	// if err != nil {
	// 	return []byte{}, err
	// }

	// Return output - an decrypted message.
	return read, nil
}
