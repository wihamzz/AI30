package main 

import (
	"fmt"
)

var n = 0

func f() {
    n++
}

func main() {
    for i := 0; i < 10000; i++ {
        go f()
    }

    fmt.Println("Appuyez sur entrée")
    //fmt.Scanln()
    fmt.Println("n:", n)
}

/* 

exercice 1 :
toutes les goroutines se lancent en même temps, et des fois certaines d'entres elles font le même n++, ce qui fait que on perd quelques valeurs au passage (en particulier si le pc est rapide)
si on enleve le scanln, le main peut se terminer avant auquel cas on aura aucun résultat 

*/ 