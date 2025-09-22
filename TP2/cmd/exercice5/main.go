package main

import (
	"fmt"
	"sync"
)

type Agent interface {
	Start()
}

type PingAgent struct {
	ID string
	c chan string
}

type PongAgent struct {
	ID string
	c chan string
}

func (pingA *PingAgent) Start() {
	pingA.c <- "ping"
}

func (pongA *PongAgent) Start() { //il faut fermer le channel dans le main après utilisation, sinon cette boucle reste bloquée --> memory leak
	for i := range pongA.c { //se bloque (se met en écoute bloquante) si jamais il n'y a rien à lire dans le channel, et bloque par conséquent la goroutine
		fmt.Println(i)
		fmt.Println("pong")
	}
}

func Dispatch(ag []PingAgent, ch chan string) {
	var wg sync.WaitGroup
	var wgPingers sync.WaitGroup
	var ponga PongAgent
	ponga.c = ch
	wg.Go(func() { //important d'attendre la fin de cette goroutine en l'ajoutant en wait group elle aussi
		ponga.Start()
	})
	for _, v := range ag { //éviter d'utiliser directement cette variable, car elle change et une goroutine pourrait prendre la mauvaise (la passer en param de la lambda)
		wgPingers.Add(1)
		go func(agent PingAgent) {
			defer wgPingers.Done()
			agent.c = ch
			agent.Start()
		}(v)
	}
	wgPingers.Wait()
	close(ch)
	wg.Wait()
}


func main() {
	var pinga1 PingAgent
	var pinga2 PingAgent
	var pinga3 PingAgent
	var pinga4 PingAgent
	var agents []PingAgent
	agents = append(agents, pinga1, pinga2, pinga3, pinga4)
	ch := make(chan string)

	Dispatch(agents, ch)

}