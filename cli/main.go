package main

import (
    "github.com/deanveloper/mpassgo"
    "fmt"
)

func main() {
    num := 5
    fmt.Println(string(mpassgo.GetPassword([]byte("dean bassett"), []byte("mtu.edu"), []byte("pass"), &num, mpassgo.LONG)))
}