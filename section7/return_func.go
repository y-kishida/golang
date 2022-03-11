package main

import "fmt"

func ReturnFunc() func()  {
	return func() {
		fmt.Println("function")
	}
}

func main()  {
	f := ReturnFunc()
	f()
}
