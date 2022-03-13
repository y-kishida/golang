package main

import "fmt"

func main()  {
	a := 0
	if a == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("no")
	}
	if b := 100; b == 100 {
		fmt.Println("100")
	}

	x := 0
	if x := 2; true {
		fmt.Println(x)
	}
	fmt.Println(x)
}
