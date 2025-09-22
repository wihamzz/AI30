package main

import (
	"math/rand"
	"fmt"
	"sort"
)

func Fill(sl []int) {
	for i := range sl {
		sl[i] = rand.Int()
	}
}

func Moyenne(sl []int) float64 {
	var sum int
	for _, val := range sl {
		sum += val
	}
	fmt.Println(sum)
	return float64(sum/len(sl))
}

func ValeursCentrales(sl []int) []int { //valeur au milieu de la slice si impair, si pair les deux valeurs du milieu
	var res []int
	if (len(sl)%2 == 1){
		center := sl[(len(sl)/2)]
		res = append(res, center)
	} else {
		center1 := sl[len(sl)/2 -1]
		center2 := sl[(len(sl)/2)]
		res = append(res, center1, center2)
	}
	return res
}

func ValeursCentrales2(sl []int) []int { //il faut que le tableau soit trié (la V1 ne marche pas) valeur centrale d'un tableau impair : mediane, tableau pair 
	copysl := make([]int, len(sl))
	copy(copysl, sl) //copy fonction builtin pour copier une slice dans une autre (erreur dans la doc --> faire une PR ????)
	sort.Ints(copysl)
	center := (len(copysl)/2)
	if (len(copysl)%2 == 1){
		// return copysl[center : center +1] //potentielle fuite memoire, il vaut mieux créer un nouveau slice, car ici, si le tableau sous jacent fait 1to, on garde le tableau en mémoire
		return []int{copysl[center]}
	} else {
		// return copysl[center -1 : center +1]
		return []int{copysl[center] -1, copysl[center]}
	}
}

func Plus1(sl []int) {
	for i := range sl {
		sl[i] += 1
	}
}

func Compte(sl []int) {
	for _, v := range sl {
		if v > -10 || v < 10 {
			fmt.Print(v)
		}
	}
}

func main() {
	fmt.Println("hello world")
	sl := []int{1,2,3,4,5}
	Fill(sl)
	sl2 := []int{1,2,3,4,5}
	sl3 := []int{5,2,1,3}
	fmt.Println(ValeursCentrales2(sl2))
	fmt.Println(ValeursCentrales2(sl3))
	Plus1(sl3)
	fmt.Println(sl3)
	Compte(sl3)
}