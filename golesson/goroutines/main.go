package main

import (
	"fmt"
	"time"
)

func main() {

	go sayHello()
	go sayWorld()

	time.Sleep(5 * time.Second)
}

func sayHello() {
	for {
		fmt.Println("Hello")
		time.Sleep(1 * time.Second)
	}
}

func sayWorld() {
	for {
		fmt.Println("World")
		time.Sleep(2 * time.Second)
	}
}
