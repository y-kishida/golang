package main

import (
	"fmt"
	"time"
)

func receiver(name string, ch chan int)  {
	for  {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main()  {
	ch1 := make(chan int, 2)

	go receiver("1.goroutine", ch1)
	go receiver("2.goroutine", ch1)
	go receiver("3.goroutine", ch1)

	i := 0
	for i < 100 {
		ch1 <-i
		i++
	}
	close(ch1)
	time.Sleep(3 * time.Second)
}
