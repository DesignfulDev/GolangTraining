package main

import "fmt"

func main() {
	fmt.Printf("%d \t %b \t %#x \t %q \n", 42, 42, 42, 42)
	fmt.Printf("%d \t %b \t %#X \t %q \n", 42, 42, 42, 42)
	fmt.Println("########################################")

	for i := 1000000; i < 1000100; i++ {
		fmt.Printf("%d \t %b \t %U \t %#x \t %#o \t %q \n", i, i, i, i, i, i)
	}

	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %#x \n", i, i, i)
	}

	fmt.Println("Trying to configure SSH keys")
	fmt.Println("Trying on more time to set SSH keys to work on Git and GitHub")
}
