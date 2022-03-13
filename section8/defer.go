package main

import (
	"fmt"
	"os"
)

func TestDefer()  {
	defer fmt.Println("END")
	fmt.Println("START")
}

func main()  {
	TestDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello"))
}

