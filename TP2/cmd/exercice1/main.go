package main

import (
	"fmt"
)

func compte(n int) {
	for i := range n {
		fmt.Println(i)
	}
}

func compteMsg(n int, msg string) {
	for i := range n {
		fmt.Println(msg, " ", i)
	}
}

func compteMsgFromTo(start int, end int, msg string) {
	for i := start ; i < end ; i ++ {
		fmt.Println(msg, " ", i)
	}
}

func main() {
	// compte(10)
	// fmt.Println("------------")
	// go compte(10)
	// go compteMsg(10, "gougou")
	// go compteMsg(10, "gougougaga")
	for i := range 10 {
		go compteMsgFromTo(i*10, i*10 + 10, "test")
	}
	fmt.Scanln()
}