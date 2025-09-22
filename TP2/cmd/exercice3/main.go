package main 

import (
	"fmt"
	"time"
)

func bonneAnneeSleep() {
	for i := 5 ; i > 0 ; i -- {
		time.Sleep(500*time.Millisecond)
		fmt.Println(i)
	}
	fmt.Println("BONNE ANNEE")
}

func bonneAnneeAfter() {  // un peu bizarre
	for i := 5 ; i > 0 ; i -- {
		fmt.Println(i)
		<- time.After(500*time.Millisecond)
	}
	fmt.Println("BONNE ANNEE")
}

func bonneAnneeTick() {
	c := time.Tick(500*time.Millisecond)
	cmp := 5
	for range c {
		if cmp == 0 {
			fmt.Println("BONNE ANNEE")
			break
		}
		fmt.Println(cmp)
		cmp--
	}
}

func main() {
	bonneAnneeAfter()
}