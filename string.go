package main

import "fmt"

func main()  {
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n",s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n",si)

	fmt.Println(`test
		test
test`)

	fmt.Println("\"")
	fmt.Println(`""`)
	fmt.Println(string(s[0]))
}
