package main

import (
"fmt"
//	"math/rand"
//	"time"
)

func main() {
	var favoriteNum, favoriteNum1 = hello(1, 2)
	fmt.Printf("Welcome to IWannaGo, My favorite number: %d, %d\n", favoriteNum, favoriteNum1)
	loop()

	fmt.Println("counting")

	for i, _ := range make([]int, 10) {
		go fmt.Println(i)
	}

	fmt.Println("done")
}

func hello(a int, b int) (int, int) {
	return a / b, a % b
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Welcome to Go: %d\n", i)
	}

	for i := 0; i < 10; i++ {
		var favoriteNum, favoriteNum1 = hello(1, 2)
		fmt.Printf("Welcome to IWannaGo, My favorite number: %d, %d\n", favoriteNum, favoriteNum1)
	}
}

