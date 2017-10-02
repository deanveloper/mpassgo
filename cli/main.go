package main

import (
    "github.com/deanveloper/mpassgo"
    "github.com/awnumar/memguard"
    "fmt"
)

func main() {
    name, _ := memguard.NewFromBytes([]byte("dean bassett"), false)
    site, _ := memguard.NewFromBytes([]byte("mtu.edu"), false)
    masterPass, _ := memguard.NewFromBytes([]byte("pass"), false)
    pass := mpassgo.GetPassword(name, site, masterPass, 1, mpassgo.Long)

    fmt.Println(string(pass.Buffer()))
}