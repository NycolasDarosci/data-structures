package main

import "fmt"

func main() {
	number := 2
	for number <= 100 {
		if number%2 == 0 {
			fmt.Println(number)
		}
		number += 1
	}

}
