package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	var length *int
	var specialChars *bool

	length = flag.Int("l", 32, "length of password")
	specialChars = flag.Bool("s", false, "use special characters")
	flag.Parse()

	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	alphabetWithSpecialCharacters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+{}|:<>?[]123456789"

	var password string

	for i := 0; i < *length; i++ {
		if *specialChars {
			password += string(alphabetWithSpecialCharacters[rand.Intn(len(alphabetWithSpecialCharacters))])
		} else {
			password += string(alphabet[rand.Intn(len(alphabet))])
		}
	}

	fmt.Printf("Your password is: %s\n", password)

}
