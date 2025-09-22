package main 

import (
	"fmt"
)

var n = 0
var ch chan int

func chan_f() {
	n = <- ch
    n++
	ch <- n
}

func main() {
	ch = make(chan int, 1) 
	/*
	Unbuffered channels must have a consumer waiting, else the producer will block.
	Buffered size one means that there is no wait, as long as the consumer empties that buffer before the next producer writes to it.
	So a buffer of size one is useful when the consumer is generally as fast as the producer, meaning that there is enough consumers to match the producers output.
	If your producers are outperforming the consumers, but only for a short time (a burst), a larger buffer will be fine.
	However if the producers are consistently, or very often, outperforming the consumers, the buffer will fill up, and the producers will slow down to the same speed as the consumers (when the buffer is full the producers have to block and wait for a consumer to make some space in the buffer for the producer to write to).
	If this is happening, it's time for the developer to assign more resources to the consumer (eg. more goroutines).
	*/
	ch <- n
    for i := 0; i < 10000; i++ {
        go chan_f()
    }

    fmt.Println("Appuyez sur entrée")
    fmt.Scanln()
    fmt.Println("n:", n)
}
