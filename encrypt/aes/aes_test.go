package aes

import (
	"testing"
)

const (
	key = "RknorRerlsraiOdff"
)

func TestEncrypt(t *testing.T) {
	t.Log(New(key).Encrypt("123456"))
}

func TestDecrypt(t *testing.T) {
	t.Log(New(key).Decrypt("GO-ri84zevE-z1biJwfQPw=="))
}

func BenchmarkEncryptAndDecrypt(b *testing.B) {
	b.ResetTimer()
	aes := New(key)
	for i := 0; i < b.N; i++ {
		encryptString, _ := aes.Encrypt("123456")
		aes.Decrypt(encryptString)
	}
}
