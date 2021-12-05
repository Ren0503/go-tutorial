package main

import "fmt"

var name1 = "Todd"

func main() {
	// Solution 1
	fmt.Println("Hello ", name1)

	// Solution 2
	var name2 = "Todd"
	fmt.Println("Hello ", name2)

	// Solution 3
	name3 := "Todd"
	fmt.Println("Hello ", name3)

	// Solution 4
	name4 := `Todd`
	fmt.Println("Hello ", name4)
}
