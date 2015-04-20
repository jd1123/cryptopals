package text

import "testing"

var secrets = []string{"yellow submarine", "longer than one block string so we can see an uneven string", ""}
var keys = []string{"valid key1234567", "invalidkey"}

func TestEncryptCBC(t *testing.T) {
	pt := NewPlaintext([]byte(secrets[0]))
	ct, _ := pt.EncryptCBC([]byte(keys[0]), nil)
	plain := NewCiphertext(ct, 16)
	if err != nil {
		t.Errorf("Error in CBC Encryption with known good values:", err)
	}
	if len(ct) != ((len(secrets[0]) % 16) + len(secrets[0]) + 16) {
		t.Errorf("Error: ciphertext not the proper length")
	}
}
