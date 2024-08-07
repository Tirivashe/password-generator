package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	var length *int
	var specialChars *bool
	var numbers *bool

	length = flag.Int("l", 32, "length of password")
	specialChars = flag.Bool("s", false, "use special characters")
	numbers = flag.Bool("n", false, "use numbers")

	flag.Parse()

	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if *specialChars {
		alphabet += "!@#$%^&*()_+"
	}

	if *numbers {
		alphabet += "0123456789"
	}

	var password string

	for i := 0; i < *length; i++ {
		password += string(alphabet[rand.Intn(len(alphabet))])
	}

	fmt.Printf("Your password is: %s\n", password)

}
