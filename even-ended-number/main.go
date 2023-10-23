package main

import "fmt"

func main() {
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {
			n := a * b
			s := fmt.Sprint(n)
			if s[0] == s[len(s)-1] {
				fmt.Println("it is even ended number")

			}
		}
	}
}
