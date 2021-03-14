package database

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/CSUOS/rabums/ent"
	"github.com/CSUOS/rabums/pkg/config"
)

//GetClient return database client
func GetClient() *ent.Client {
	return client
}

//EncryptPassword get hashed password
func EncryptPassword(pw string) string {
	return hash(pw + config.Database.Salt)
}

//IsEqualPassword Check if given passwords are equal
func IsEqualPassword(encrypted, decrypted string) bool {
	return encrypted == EncryptPassword(decrypted)
}

func hash(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	b := hash.Sum(nil)
	return hex.EncodeToString(b)
}
