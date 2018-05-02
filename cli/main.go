package main

import (
	"bufio"
	"fmt"
	"github.com/deanveloper/mpassgo"
	"golang.org/x/crypto/ssh/terminal"
	"math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

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
		scanner.Scan()
		name = scanner.Bytes()
	}
	defer zeroSlice(name)

	fmt.Println()

	// get master password
	fmt.Println("Password")
	fmt.Print("Input: ")
	masterPass, err := terminal.ReadPassword(0)
	defer zeroSlice(masterPass)

	// if terminal is not available (ie executing from IDE) then just let it echo
	if err != nil {
		fmt.Println("[error: cannot use native password reader. password will echo!]")
		fmt.Print("Input: ")
		scanner.Scan()
		masterPass = scanner.Bytes()
	} else {
		fmt.Println()
	}

	fmt.Println()

	// get website
	fmt.Println("Website")
	fmt.Print("Input: ")
	scanner.Scan()
	site := scanner.Bytes()
	defer zeroSlice(site)

	fmt.Println()

	// get counter
	fmt.Println("Counter [blank for 1]")
	fmt.Print("Input: ")
	scanner.Scan()
	counter := readDigits(scanner.Bytes())
	for counter <= 0 {
		fmt.Println("error: Please input a positive number!")
		fmt.Print("Input: ")
		scanner.Scan()
		counter = readDigits(scanner.Bytes())
	}
	defer func() { counter = 0 }()

	fmt.Println()

	// get password type
	// (note, since string must be used in GetPasswordType, we use strings here)
	var pwType mpassgo.PasswordType
	fmt.Println("Password Type [blank for maximum]")
	fmt.Print("Input: ")
	scanner.Scan()
	pwTypeStr := scanner.Text()

	if pwTypeStr == "" {
		pwType, ok = mpassgo.Maximum, true
	} else {
		pwType, ok = mpassgo.GetPasswordType(pwTypeStr)
	}

	for !ok {
		fmt.Println("Try again! Valid password types: ")

		for i, typ := range mpassgo.ValidPasswordTypes() {
			if i != 0 {
				fmt.Print(", ")
			}
			fmt.Print(typ)
		}

		fmt.Print("Input: ")
		scanner.Scan()
		pwTypeStr = scanner.Text()

		if pwTypeStr == "" {
			pwType, ok = mpassgo.Maximum, true
		} else {
			pwType, ok = mpassgo.GetPasswordType(pwTypeStr)
		}
	}

	fmt.Println()

	println(counter)
	pass, err := mpassgo.GetPassword(name, site, masterPass, counter, pwType)
	if err != nil {
		fmt.Println("Error generating password: ", err)
	}

	fmt.Println(string(pass))
}

// turns an array of ASCII digits into an int. 1 if empty, 0 if error.
// also handles if it ends in a newline
func readDigits(slice []byte) int {

	result := 0
	for i := range slice {
		byt := slice[len(slice)-i-1]
		if byt >= 48 && byt <= 57 {
			result += int(byt-48) * int(math.Pow10(i))
		} else {
			return 0
		}
	}

	return result
}

func zeroSlice(slice []byte) {
	for i := range slice {
		slice[i] = 0
	}
}
