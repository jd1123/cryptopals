package padding

import "testing"

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
