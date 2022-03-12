package main

import "fmt"

func main()  {
	n := 5
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3,4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("anything")
	}

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3,4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("anything")
	}
}