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
var seedPrefix = []byte("com.lyndir.masterpassword")

// A convenience function that takes all needed inputs and generates the password from there.
func GetPassword(name, website, masterPass []byte, counter int, set TemplateSet) ([]byte, error) {
	mKey, err := GetMasterKey(name, masterPass)
	if err != nil {
		return nil, err
	}
	defer zeroSlice(mKey)

	sKey := GetSiteKey(mKey, website, counter)
	defer zeroSlice(sKey)

	template := GetTemplate(sKey, set)
	password := SiteKeyToPassword(sKey, template)

	return password, nil
}

// Gets the key for a given password.
// A side effect of using this function is that the password
// slice gets zero'd. This is done to discourage API users from
// using the master password anywhere else. Both the name and
// password should be UTF8 encoded.
func GetMasterKey(name, password []byte) ([]byte, error) {

	seedLen := len(seedPrefix) + 4 + len(name)
	seed := bytes.NewBuffer(make([]byte, seedLen))
	seed.Write(seedPrefix)
	seed.Write(convertNum(len(name)))
	seed.Write(name)

	return scrypt.Key(password, seed.Bytes(), 32768, 8, 2, 64)
}

// Gets the seed provided a key and a counter. The website
// should be UTF8 encoded.
func GetSiteKey(mKey, website []byte, counter int) []byte {
	seedLen := len(seedPrefix) + 4 + len(website) + 4
	seed := bytes.NewBuffer(make([]byte, 0, seedLen))

	seed.Write(seedPrefix)
	seed.Write(convertNum(len(website)))
	seed.Write(website)
	seed.Write(convertNum(counter))

	hasher := hmac.New(sha256.New, mKey)
	hasher.Write(seed.Bytes())
	return hasher.Sum(nil)
}

// Figures out what the password template is going to be.
func GetTemplate(sKey []byte, templates TemplateSet) []byte {
	return templates[sKey[0]%byte(len(templates))]
}

// Encodes the seed into the template, generating the password.
// Panics if an invalid template string is provided.
func SiteKeyToPassword(sKey []byte, template []byte) []byte {
	pass := make([]byte, len(template))
	for i, e := range template {
		r, _ := utf8.DecodeRune([]byte{e})
		runes := runeMap[r]
		if runes == "" {
			panic(ErrInvalidTemplate)
		}

		pass[i] = runes[sKey[i+1]%byte(len(runes))]
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
