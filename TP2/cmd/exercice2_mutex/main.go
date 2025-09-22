package main 

import (
	"fmt"
	"sync"
)

var n = 0
var mu sync.Mutex

func mutex_f() {
	mu.Lock()
    n++
	mu.Unlock()
}

func main() {
    for i := 0; i < 10000; i++ {
        go mutex_f()
    }

    fmt.Println("Appuyez sur entrée")
    fmt.Scanln()
    fmt.Println("n:", n)
}