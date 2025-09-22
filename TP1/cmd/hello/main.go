package main

import "fmt"

func pair() {
	for i := 0 ; i <= 1000 ; i += 2 {
		fmt.Println(i)
	}
}

func main() {
	pair()
}
