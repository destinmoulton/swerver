package pw

import (
	"github.com/destinmoulton/swerver/app/lib/crypto"
)

// EncryptPassword encrypts a password with a 32byte key
func EncryptPassword(key, pwd string) string {
	enc, err := crypto.Encrypt([]byte(key), pwd)
	if err != nil {
		panic(err)
	}
	return enc
}

func DecryptPassword(key, pwd string) string {
	dec, err := crypto.Decrypt([]byte(key), pwd)
	if err != nil {
		panic(err)
	}
	return dec
}
