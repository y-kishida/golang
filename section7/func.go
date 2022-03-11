package main

import "fmt"

func Plus(x, y int) int  {
	return x + y
}

func Div(x, y int) (int, int)  {
	q := x / y
	r := x * y
	return q, r
}

func Double(price int) (result int)  {
	return price * 2
}

func main()  {
	i := Plus(1,2)
	fmt.Println(i)

	i2, _ := Div(3,4)
	fmt.Println(i2)

	i4 := Double(2000)
	fmt.Println(i4)
}
