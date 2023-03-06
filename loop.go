package main

import "fmt"

func main() {
	for i := 1; i <= 6; i++ {
		for j := 0; j < i; j++ {
			if j == i-1 || j == 0 || i == 6 {
				fmt.Print("*")
			} else if i > 3 && j == 1 {
				fmt.Print(".")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
