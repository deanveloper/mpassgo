package mpassgo

import (
    "encoding/binary"
    "github.com/awnumar/memguard"
    "crypto/hmac"
    "crypto/sha256"
    "golang.org/x/crypto/scrypt"
    "fmt"
)

var saltPrefix = []byte("com.lyndir.masterpassword")

func GetPassword(name, site, pass *memguard.LockedBuffer, counter int, templateSet TemplateSet) *memguard.LockedBuffer {

    // first, make the buffer for the key's salt
    saltLength := len(saltPrefix) + 4 + len(name.Buffer())
    salt, err := memguard.New(saltLength, false)
    checkErr(err)

    // copy stuff into salt buffer
    salt.Move(saltPrefix)
    salt.MoveAt(convertNum(len(name.Buffer())), len(saltPrefix))
    salt.MoveAt(name.Buffer(), len(saltPrefix) + 4)

    for i, e := range pass.Buffer() {
        fmt.Println(i, e)
    }

    // make key based on salt
    key, err := scrypt.Key(pass.Buffer(), salt.Buffer(), 32768, 8, 2, 64)
    checkErr(err)

    // make buffer for what we're seeding
    seedLength := len(saltPrefix) + 4 + len(site.Buffer()) + 4
    seedBuf, err := memguard.New(seedLength, false)
    checkErr(err)

    seedBuf.Move(saltPrefix)
    seedBuf.MoveAt(convertNum(len(site.Buffer())), len(saltPrefix))
    seedBuf.MoveAt(site.Buffer(), len(saltPrefix) + 4)
    seedBuf.MoveAt(convertNum(counter), len(saltPrefix) + 4 + len(site.Buffer()))

    seed := hmac.New(sha256.New, key)
    seed.Write(seedBuf.Buffer())

    template := templateSet[int(seed.Sum(nil)[0]) % len(templateSet)]

    password, err := memguard.New(len(template), false)
    checkErr(err)
    for i, r := range template {
        passChars := charTemplates[r]
        passChar := passChars[int(seed.Sum(nil)[i + 1]) % len(passChars)]
        password.CopyAt([]byte{passChar}, i)
    }

    return password
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