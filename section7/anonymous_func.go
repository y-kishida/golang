package main

import "fmt"

func main()  {
	f := func(x, y int) int {
		return x + y
	}
	i := f(1, 2)
	fmt.Println(i)
}