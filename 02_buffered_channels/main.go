package main

import "fmt"

func main(){
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	fmt.Println(<-c)
	fmt.Println(<-c)
}
