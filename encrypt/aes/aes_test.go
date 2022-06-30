package aes

import (
	"testing"
)

const (
	key = "qU6y_AS7&gWtH0I2Z*unOoBxGTmjrYRX"
)

func TestEncrypt(t *testing.T) {
	t.Log(New(key).Encrypt("wb_mads"))
}

func TestDecrypt(t *testing.T) {
	t.Log(New(key).Decrypt("Ue3xxa1K_ppFfMWyJe3ebQ=="))
}

func BenchmarkEncryptAndDecrypt(b *testing.B) {
	b.ResetTimer()
	aes := New(key)
	for i := 0; i < b.N; i++ {
		encryptString, _ := aes.Encrypt("123456")
		aes.Decrypt(encryptString)
	}
}
