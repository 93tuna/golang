package main

import "fmt"

func main() {
	/* for i := 2; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("%d단\n", i)
		for j := 1; j < 10; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	} */

	/* for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	} */

	/* for i := 0; i < 5; i++ {
		for j := 5; j > i; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	} */

	/* for i := 0; i < 5; i++ {
		for j := 0; j <= i && j <= 4-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	} */

	/* for i := 3; i >= 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 7-2*i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	} */

	for i := 0; i < 4; i++ {
		for j := 0; j < 3-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
