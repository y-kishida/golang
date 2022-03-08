package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 1
	fl64 := float64(i)
	fmt.Println(fl64)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)

	var s string = "100"

	str, _ := strconv.Atoi(s)
	fmt.Println(str)
	fmt.Printf("str = %T\n", str)

	var i2 int = 200
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("%T\n", s2)

	var h string = "Hello"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
}
