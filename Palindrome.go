package main

import (
	"fmt"
)

func main() {
	n := 1210
	revnum := 0
	temp := n
	for n != 0 {
		rem := n % 10
		revnum = revnum*10 + rem
		n = n / 10
	}
	if revnum == temp {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}

}
