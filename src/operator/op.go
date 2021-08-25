package main

import "fmt"

func main() {
	a := 4
	b := 2
	c := 21
	c = 21 % 10

	fmt.Printf("a & b = %v\n", a&b)
	fmt.Printf("a | b = %v\n", a|b)
	fmt.Printf("a ^ b = %v\n", a|b)
	fmt.Printf("^a = %v\n", ^a)
	fmt.Printf("^b = %v\n", ^b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", a>>1)
	fmt.Printf("%v\n", a<<2)

	d := 3 > 4

	fmt.Println(d)

	fmt.Println(true && false)
	fmt.Println(true || false)

	if 3 > 4 {
		fmt.Println("3 > 4")
	} else if 3 == 4 {
		fmt.Println("3 == 4")
	} else {
		fmt.Println("3 < 4")
	}

}
