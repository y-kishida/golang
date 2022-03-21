package main

import (
	"fmt"
	"time"
)

func receiver(c chan int)  {
	for  {
		i := <-c
		fmt.Println(i)
	}
}

func main()  {
	ch := make(chan int)
	ch2 := make(chan int)

	go receiver(ch)
	go receiver(ch2)

	i := 0
	for i < 100 {
		ch <- i
		ch2<- i
		time.Sleep(50 * time.Millisecond)
		i++
	}
	
}
