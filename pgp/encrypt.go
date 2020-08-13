package pgp

import (
	"bytes"
	_ "crypto/sha256"
	"fmt"

	"golang.org/x/crypto/openpgp"
	_ "golang.org/x/crypto/ripemd160"
)

func Encrypt(entity *openpgp.Entity, message []byte) ([]byte, error) {
	// Create buffer to write output to
	buf := new(bytes.Buffer)

	// Create encoder -- this is part of https://github.com/jchavannes/go-pgp/
	// encoderWriter, err := armor.Encode(buf, "Message", make(map[string]string))
	// if err != nil {
	// 	return []byte{}, fmt.Errorf("Error creating OpenPGP armor: %v", err)
	// }

	// Create encryptor with buffer
	encryptorWriter, err := openpgp.Encrypt(buf, []*openpgp.Entity{entity}, nil, nil, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("Error creating entity for encryption: %v", err)
	}

	// Create compressor with encryptor -- this is part of https://github.com/jchavannes/go-pgp/
	// compressorWriter, err := gzip.NewWriterLevel(encryptorWriter, gzip.BestCompression)
	// if err != nil {
	// 	return []byte{}, fmt.Errorf("Invalid compression level: %v", err)
	// }

	// // Write message to compressor -- this is part of https://github.com/jchavannes/go-pgp/
	// messageReader := bytes.NewReader(message)
	// _, err = io.Copy(compressorWriter, messageReader)
	// if err != nil {
	// 	return []byte{}, fmt.Errorf("Error writing data to compressor: %v", err)
	// }

	// compressorWriter.Close()
	encryptorWriter.Close()
	// encoderWriter.Close()

	// Return buffer output - an encoded, encrypted, and compressed message
	return buf.Bytes(), nil
}
