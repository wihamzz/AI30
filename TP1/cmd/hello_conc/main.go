package main

import "fmt"

func pair() {
	for i := 0 ; i <= 1000 ; i += 2 {
		go fmt.Println(i)
	}
	fmt.Scanln()
}

func main() {
	pair()
}
