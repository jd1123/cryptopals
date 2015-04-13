package padding

import (
	"log"
	"testing"
)

func TestPKCS7(t *testing.T) {
	b := []byte("This is not a the correct blocksize")
	padded := PKCS7(b, 16)
	if padded[len(padded)-1] != byte(13) {
		t.Errorf("Padding did not work correctly")
	}
	b = []byte("YELLOW SUBMARINE")
	if len(b) != len(PKCS7(b, 16)) {
		t.Errorf("Full blocks should not be padded")
	}
}

func TestValidatePKCS7(t *testing.T) {
	b := []byte("This is not a the correct blocksize")
	padded := PKCS7(b, 16)
	res, err := ValidatePKCS7(padded, 16)
	if err != nil {
		t.Errorf("Padding should validate for known good result. It does not.")
	}

	if len(res) > len(padded) {
		t.Errorf("The validated result should not be a greater length than the padded message It is")
	}

	if len(res) != len(b) {
		t.Errorf("The validated result should equal the original message. It does not.")
		log.Println(":\nb:", b, "\nres:", res, "\n")
	}
}
