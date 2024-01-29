package main

import (
	   "fmt"
	   "regexp"
)

func main () {

	text := "Hello my email address is bogdan@yahoo.com and Phone is 077-033-0445"

	phoneRegex := regexp.MustCompile(`\d[3]-\d[3]-\d[4]`)
	phone := phoneRegex.FindString(text)

	fmt.Printf("Phone is: %s", phone)
}