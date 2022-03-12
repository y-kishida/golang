package main

import "fmt"

func anything(a interface{})  {
	fmt.Println(a)
}

func main()  {
	anything("aaa")
	anything(1)

	var x interface{} = 3
	i, isInt := x.(int)
	fmt.Println(i, isInt)

	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)
}
