package main

import "fmt"

func main()  {
	sl := []string{"A", "B", "C"}

	for i := range sl {
		fmt.Println(i, sl[i])
	}

}
