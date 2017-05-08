package pgp_test

import (
	"fmt"
	"testing"
	"github.com/jchavannes/go-pgp/pgp"
)

func TestSign(t *testing.T) {
	entity, err := pgp.GetEntity([]byte(TestPublicKey), []byte(TestPrivateKey))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Test message: %s\n", TestMessage)
	signature, err := pgp.Sign(entity, []byte(TestMessage))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Signature: %s\n", signature)

	pubKeyPacket, err := pgp.GetPublicKeyPacket([]byte(TestPublicKey))
	if err != nil {
		t.Error(err)
	}
	err = pgp.Verify(pubKeyPacket, []byte(TestMessage), signature)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("Signature verified.")
	}
}