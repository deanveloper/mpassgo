package mpassgo

import (
    "encoding/binary"
    "crypto/hmac"
    "crypto/sha256"
    "golang.org/x/crypto/scrypt"
    "bytes"
)

var saltPrefix = []byte("com.lyndir.masterpassword")

func GetPassword(name, site, pass []byte, counter int, templateSet TemplateSet) []byte {

    // first, make the buffer for the key's salt
    saltLength := len(saltPrefix) + 4 + len(name)
    salt := bytes.NewBuffer(make([]byte, saltLength))

    // copy stuff into salt buffer
    salt.Write(saltPrefix)
    salt.Write(convertNum(len(name)))
    salt.Write(name)

    ZeroSlice(name)

    // make key based on salt
    key, err := scrypt.Key(pass, salt.Bytes(), 32768, 8, 2, 64)
    checkErr(err)

    ZeroSlice(pass)

    // make buffer for what we're seeding
    seedLength := len(saltPrefix) + 4 + len(site) + 4
    seedBuf := bytes.NewBuffer(make([]byte, seedLength))

    seedBuf.Write(saltPrefix)
    seedBuf.Write(convertNum(len(site)))
    seedBuf.Write(site)
    seedBuf.Write(convertNum(counter))

    ZeroSlice(site)

    seed := hmac.New(sha256.New, key)
    seed.Write(seedBuf.Bytes())

    template := templateSet[int(seed.Sum(nil)[0]) % len(templateSet)]

    password := bytes.NewBuffer(make([]byte, len(template)))
    checkErr(err)
    for i, r := range template {
        passChars := charTemplates[r]
        passChar := passChars[int(seed.Sum(nil)[i + 1]) % len(passChars)]
        password.WriteByte(passChar)
    }

    return password.Bytes()
}

func ZeroSlice(s []byte) {
    for i := range s {
        s[i] = 0
    }
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func convertNum(i int) []byte {
    b := make([]byte, 4)
    binary.BigEndian.PutUint32(b, uint32(i))
    return b
}