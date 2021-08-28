package main

import "fmt"

func main() {
	x := 30

	/* switch x {
	case 30:
		fmt.Println("x = 30")
	case 31:
		fmt.Println("x = 31")
	} */
	switch {
	case x > 30:
		fmt.Println("x는 30보다 크다.")
	case x == 30:
		fmt.Println("x는 30이다.")
	case x < 30:
		fmt.Println("x는 30보다 작다.")
	}
}
