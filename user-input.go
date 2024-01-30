package main

import "fmt"

func main(){
	fmt.Println("Enter your name: ")
	var firstname string
	fmt.Scanln(& firstname)

	fmt.Println("Your full name is:"+firstname)

}