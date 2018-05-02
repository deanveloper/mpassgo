package mpassgo

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"errors"

	"golang.org/x/crypto/scrypt"
	"unicode/utf8"
)

var ErrInvalidTemplate = errors.New("invalid template provided")
var saltPrefix = []byte("com.lyndir.masterpassword")

// A convenience function that takes all needed inputs and generates the password from there.
func GetPassword(name, masterPass, website []byte, counter int, set TemplateSet) ([]byte, error) {
	key, err := GetKey(name, masterPass)
	if err != nil {
		return nil, err
	}
	defer zeroSlice(key)

	seed := GetSeed(key, website, counter)
	defer zeroSlice(seed)

	template := GetTemplate(seed, set)
	password := SeedToPassword(seed, template)

	return password, nil
}

// Gets the key for a given password.
// A side effect of using this function is that the password
// slice gets zero'd. This is done to discourage API users from
// using the master password anywhere else. Both the name and
// password should be UTF8 encoded.
func GetKey(name, password []byte) ([]byte, error) {

	saltLen := len(saltPrefix) + 4 + len(name)
	salt := bytes.NewBuffer(make([]byte, saltLen))
	salt.Write(saltPrefix)
	salt.Write(convertNum(len(name)))
	salt.Write(name)

	return scrypt.Key(password, salt.Bytes(), 32768, 8, 2, 64)
}

// Gets the seed provided a key and a counter. The website
// should be UTF8 encoded.
func GetSeed(key, website []byte, counter int) []byte {
	seededLen := len(saltPrefix) + 4 + len(website) + 4
	seeded := bytes.NewBuffer(make([]byte, 0, seededLen))

	seeded.Write(saltPrefix)
	seeded.Write(convertNum(len(website)))
	seeded.Write(website)
	seeded.Write(convertNum(counter))

	hasher := hmac.New(sha256.New, key)
	return hasher.Sum(seeded.Bytes())
}

// Figures out what the password template is going to be.
func GetTemplate(seed []byte, templates TemplateSet) []byte {
	return templates[seed[0]%byte(len(templates))]
}

// Encodes the seed into the template, generating the password.
// Panics if an invalid template string is provided.
func SeedToPassword(seed []byte, template []byte) []byte {
	pass := make([]byte, len(template))
	for i, e := range template {
		r, _ := utf8.DecodeRune([]byte{e})
		runes := runeMap[r]
		if runes == "" {
			panic(ErrInvalidTemplate)
		}

		pass[i] = runes[seed[i+1]%byte(len(runes))]
	}

	return pass
}

func zeroSlice(slice []byte) {
	for i := range slice {
		slice[i] = 0
	}
}

func convertNum(i int) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(i))
	return b
}
