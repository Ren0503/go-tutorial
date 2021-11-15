package main

import "fmt"

func main() {
	// decimal
	fmt.Println(42)

	// decimal & binary
	fmt.Printf("%d - %b \n", 42, 42)

	// decimal, binary & hexadecimal
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)

	// Loop
	for i := 0; i < 100; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
