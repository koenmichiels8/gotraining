package main

import "fmt"

func main() {
	for index := 1000; index < 1100; index++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", index, index, index, index)
	}

}
