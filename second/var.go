package main

import "fmt"

func main() {

	var name string = "Shekhar"
	var age = 30
	//Note that whatever declared must be used
	fmt.Printf("Hello, %s\n, %d\n", name, age)

	usn := "rj11"

	fmt.Printf("USN : %s\n", usn)
	fmt.Println("Hello ", name)
	//int
	//string
	//float
	//bool
	//complex
	var a complex128 = 1 + 2i
	fmt.Printf("Complex number: %v\n", a)
}
