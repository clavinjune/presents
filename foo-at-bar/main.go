package main

import "fmt"

func show() {
	fmt.Println("Foobar") // HL2
}

func main() { // HL1
	show() // HL1
} // HL1
