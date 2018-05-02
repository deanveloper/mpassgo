package main

import (
	"github.com/deanveloper/mpassgo"
	"golang.org/x/crypto/ssh/terminal"
	"fmt"
	"bufio"
	"os"
)

func main() {

	fmt.Println()

	// get name
	var name []byte
	var err error
	nameStr, ok := os.LookupEnv("MPW_FULLNAME")
	if ok {
		name = []byte(nameStr)
	} else {
		fmt.Println("Full Name")
		fmt.Print("Input: ")
		name, err = readLine()
		if err != nil {
			fmt.Println("Error reading site: ", err)
			return
		}
	}
	defer zeroSlice(name)

	fmt.Println()

	// get master password
	fmt.Println("Password: ")
	fmt.Print("Input: ")
	masterPass, err := terminal.ReadPassword(0)
	defer zeroSlice(masterPass)

	// if terminal is not available (ie executing from IDE) then just let it echo
	if err != nil {
		fmt.Println("[error: cannot use native password reader. password will echo!]")
		fmt.Print("Input: ")
		masterPass, err = readLine()
		if err != nil {
			fmt.Println("Error reading master password: ", err)
			return
		}
	} else {
		fmt.Println()
	}

	fmt.Println()

	// get website
	fmt.Println("Website")
	fmt.Print("Input: ")
	site, err := readLine()
	if err != nil {
		fmt.Println("Error reading site: ", err)
		return
	}
	defer zeroSlice(site)

	// get counter
	fmt.Println("Counter [blank for 1]")
	fmt.Print("Input: ")
	var counter int
	_, err = fmt.Scanln(&counter)
	for err != nil {
		fmt.Println("error: Please input a number!")
		_, err = fmt.Scanln(&counter)
	}
	defer func() { counter = 0 }()

	fmt.Println()

	// get password type (note, since string must be used,
	fmt.Println("Password Type [blank for maximum]")
	fmt.Print("Input: ")
	var pwTypeStr string
	fmt.Scanln(&pwTypeStr)

	pwType, ok := mpassgo.GetPasswordType(pwTypeStr)
	for !ok {
		fmt.Println("Try again! Valid password types: ")

		for i, typ := range mpassgo.ValidPasswordTypes() {
			if i != 0 {
				fmt.Print(", ")
			}
			fmt.Print(typ)
		}

		fmt.Print("Input: ")
		fmt.Scanln(&pwTypeStr)
	}

	fmt.Println()

	pass, err := mpassgo.GetPassword(name, site, masterPass, counter, pwType)
	if err != nil {
		fmt.Println("Error generating password: ", err)
	}

	fmt.Println(string(pass))
}

func readLine() ([]byte, error) {
	read := bufio.NewReader(os.Stdin)
	return read.ReadSlice(10) // 10 is newline
}

func zeroSlice(slice []byte) {
	for i := range slice {
		slice[i] = 0
	}
}
