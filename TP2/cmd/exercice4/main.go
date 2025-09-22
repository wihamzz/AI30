package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	//"runtime" //connaitre nombre coeurs cpu
	"sync"
	"time"
)

var nbworkers = 5 //si on a moins de coeurs que de workers, le gain de temps va s'effacer car les goroutines vont devoir se partager des coeurs
var wg sync.WaitGroup // ne pas le déclarer en global en temps normal, la c'est juste pour faire mes tests

func Fill(tab []int, v int) {
	for i := range tab {
		tab[i] = v
	}
}

// serait mieux avec une borne start et end pour chaque tableau plutot que total et start
func FillConc(tab []int, v int) {
	size_sub_tabs := len(tab) / nbworkers
	var total int
	for range nbworkers {
		start := total //copie locale pour l'itération --> chaque goroutine aura son propre start
		wg.Add(1)
		go func() { //lambda fonction permet notamment d'envoyer plusieurs instructions dans une goroutine, on peut lire les variables extérieures (closure)
			defer wg.Done()
			for i := start; i < start+size_sub_tabs; i++ {
				tab[i] = v
			}
		}() //permet de donner tout de suite un argument
		total = total + size_sub_tabs // !!! Toute les goroutines lisent total qui est une variable de la fonction externe, mais elle change au fil des itérations
		// ce qui fait que lorsqu'une goroutine démarre, total n'a plus forcément la valeur 0, (alors que dans l'ordre d'exec de la fonction on pourrait croire que c'est bon)
		// elles sautent en gros à la valeur de total déjà incrémentée
	}
	wg.Wait()
	for i := total; i < len(tab); i++ {
		tab[i] = v
	}
}

func ForEach(tab []int, f func(int) int) {
	for i, v := range tab {
		tab[i] = f(v)
	}
}

// attention, dans un for i, v := range tab[10:100], le i commence a 0 et pas à 10 !! il faut donc faire le décalage manuellement (voir les deux boucles for)
func ForEachConc(tab []int, f func(int) int) {
	slice_size := len(tab) / nbworkers
	for i := range nbworkers {
		wg.Add(1)
		go func() { // bloc ligne 55, 56, 57 remplacable par wg.Go(func(){...})
			defer wg.Done()
			for j, v := range tab[slice_size*i : slice_size*(i+1)] { //dans [start:end], l'indice end n'est pas compris
				tab[j+slice_size*i] = f(v)
			}
		}() //potentiellement passer i en paramètre de la lambda fonction, sinon i risque de changer avant qu'une goroutine aie eu le temps d'y toucher
	}
	wg.Wait()
	for i, v := range tab[len(tab)-len(tab)%nbworkers:] { //jusqu'à la fin du tableau
		tab[i+len(tab)-len(tab)%nbworkers] = f(v)
	}
}

func Copy(src []int, dest []int) {
	if len(src) != len(dest) {
		log.Fatal("Copy : src and dest len are different")
		return
	}
	for i := range src {
		dest[i] = src[i]
	}
}

func CopyConc(src []int, dest []int) {
	if len(src) != len(dest) {
		log.Fatal("CopyConc : src and dest len are different")
		return
	}
	slice_size := len(src) / nbworkers
	for i := range nbworkers {
		wg.Go(func() {
			for j := range src[slice_size*i : slice_size*(i+1)] {
				dest[j+slice_size*i] = src[j+slice_size*i]
			}
		})
	}
	wg.Wait()
	for i := range src[len(src)-(len(src)%nbworkers):] {
		dest[i+len(src)-(len(src)%nbworkers)] = src[i+len(src)-(len(src)%nbworkers)]
	}
}

func Equal(tab1 []int, tab2 []int) bool {
	//paralleliser equal sans channel est compliqué car il faut attendre que toutes les goroutines aient fini, (pas trop de gain de perf possible)
	//alors qu'en séquentiel, dès qu'on trouve quelque chose de différent on s'arrete
	if len(tab1) != len(tab2) {
		return false
	}

	for i := range tab1 {
		if tab1[i] != tab2[i] {
			return false
		}
	}
	return true
}

//meilleur si slices identiques
func EqualConc(tab1 []int, tab2 []int) { //on écrit dans le channel, output channel, utilisable uniquement pour envoyer des bool
	ch := make(chan bool, nbworkers+1) //utilisé localement donc pas besoin de le prendre en parametre
	if len(tab1) != len(tab2) {
		ch <- false
		return
	}
	slice_size := len(tab1) / nbworkers
	for i := range nbworkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range tab1[slice_size*i:slice_size*(i+1)] {
				if tab1[j+slice_size*i] != tab2[j+slice_size*i] {
					ch <- false
					break
				}
			}
		}() //potentiellement passer i en paramètre de la lambda fonction, sinon i risque de changer avant qu'une goroutine aie eu le temps d'y toucher
	}
	wg.Wait()
	ch <- true
	close(ch)
}

func FasterEqualConc(tab1 []int, tab2 []int, ch chan <- bool) {  //meilleur si valeurs différentes dans une des slices (surtout vers le début)
	if len(tab1) != len(tab2) {
		ch <- false
		return
	}
	done_ch := make(chan struct{}) //utilisé pour envoyer un signal (ne prend pas de place en mémoire)
	slice_size := len(tab1) / nbworkers
	for i := range nbworkers {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := range tab1[slice_size*i:slice_size*(i+1)] {
				select {
				case <-done_ch : //tant que le channel n'est pas fermé, ce cas est bloqué car le channel est vide
					return
				default:
					if tab1[j+slice_size*i] != tab2[j+slice_size*i] {
						ch <- false
						close(done_ch) //lorsqu'un channel est fermé, il est ouvert à la lecture (donc le cas ci dessus s'ouvre aussi)
						return
					}
				}
			}
		}(i)
	}
	wg.Wait()
	ch <- true
	close(ch)
}

// trouve la valeur et le premier élément de tab vérifiant f (-1, 0 si rien n'est trouvé)
func Find(tab []int, f func (int) bool) (index int, val int) {
	for index,val := range tab {
		if f(val) {
			return index, val
		}
	}
	return -1, 0
}

func FindConc(tab []int, f func (int) bool) (index int, val int) {
	ch := make(chan int, 1) //pas obligé de stocker 2 valeurs, on peut juste stocker l'indice et l'utiliser pour avoir la valeur (tab[i])
	done := make(chan struct{}, nbworkers) //signal channel
	slice_size := len(tab)/nbworkers
	tab_rest := len(tab)%nbworkers
	for i := range nbworkers {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for id, val := range tab[slice_size*i:slice_size*(i+1)] {
				select {
				case <- done : //si ce case est dispo, on passe par là et la goroutine s'arrête (donc toutes les goroutines seront notifiées)
					return
				default :
					if f(val) { //si on trouve la valeur qu'on cherche, on envoie l'indice au channel, on le ferme (en écriture donc) on envoie un signal done, et on quitte la goroutine
						ch <- id+slice_size*i //on a pas l'id reel car on travaille sur des portions du tableau principal
						// close(ch) // fermer le channel ici peut être risqué, ne fermer un channel que quand on est sur qu'il ne va plus rien avoir + qu'on a besoin de le fermer (pour faire un range dessus par ex) --> fermer un channel n'est pas obligatoire
						// for range nbworkers{ //marche mais pas besoin de faire ça
						// 	done <- struct{}{}
						// }
						close(done) //fermer le channel permet de faire en sorte (s'il est vide) de permettre la lecture de la zero value, et donc de permettre le passage dans le select
					}
				}
			}
		}(i)
	}
	wg.Wait() // possible de mettre le wait dans une goroutine et d'enlever le buffer de ch, mais pas très interessant ici

	if len(ch) >= 1 {
		index := <- ch
		return index, tab[index]
	}
	
	for id, val := range tab[len(tab)-tab_rest:] {
		if f(val) {
			return id+len(tab)-tab_rest, val
		}
	}

	return -1, 0
}

// crée un nouveau tableau à partir de tab tel que les éléments du nouveau tableau sont composés des éléments du précédent tableau auxquels on a appliqué la fonction f
func Map(tab []int, f func (int) int) []int {
	new_tab := make([]int, len(tab))
	for i, v := range tab {
		new_tab[i] = f(v)
	}
	return new_tab
}

func MapConc(tab []int, f func(int) int) []int {
	new_tab := make([]int, len(tab))
	slice_size := len(tab)/nbworkers
	rest_size := len(tab)%nbworkers
	for i := range nbworkers {
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			for index, val := range tab[slice_size*i:slice_size*(i+1)] {
				new_tab[index+slice_size*i] = f(val)
			}
		}(i)
	}

	wg.Wait()

	for index, val := range tab[len(tab)-rest_size:] {
		new_tab[index+len(tab)-rest_size] = f(val)
	}

	return new_tab
}

// reduit le tableau à un entier (agrégation de tous les éléments avec la fonction f)
func Reduce(tab []int, init int, f func (int, int) int) int {
	
}




func isNum(n int) bool {
	return n == 424242
}

func add1(n int) int {
	return n + 1
}

func main() {
	tab := make([]int, 2<<25)
	tab2 := make([]int, 2<<25)
	tab3 := make([]int, 2<<25)
	tab4 := make([]int, 2<<25)
	tab5 := make([]int, 2<<25)
	tab6 := make([]int, 2<<25)
	tab7 := make([]int, 2<<25)
	Fill(tab5, 1)
	tab5[42000000] = 424242

	d := time.Now()
	Fill(tab, math.MaxInt64)
	fmt.Println("fill : ", time.Since(d))
	d = time.Now()
	ForEach(tab, rand.Intn)
	fmt.Println("foreach : ", time.Since(d))
	d = time.Now()
	Copy(tab, tab2)
	fmt.Println("copy : ", time.Since(d))
	d = time.Now()
	Equal(tab, tab2)
	fmt.Println("equal : ", time.Since(d))
	d = time.Now()
	Find(tab5, isNum)
	fmt.Println("find : ", time.Since(d))
	d = time.Now()
	Map(tab6, add1)
	fmt.Println("map : ", time.Since(d))
	d = time.Now()
	FillConc(tab3, math.MaxInt64)
	fmt.Println("fillconc : ", time.Since(d))
	d = time.Now()
	ForEachConc(tab3, rand.Intn)
	fmt.Println("foreachconc : ", time.Since(d))
	d = time.Now()
	CopyConc(tab3, tab4)
	fmt.Println("copyconc : ", time.Since(d))
	d = time.Now()
	EqualConc(tab3, tab4)
	fmt.Println("equalconc : ", time.Since(d)) //plus rapide car les tableaux sont égaux ici
	d = time.Now()
	FindConc(tab5, isNum) //vitesse qui dépend du nombre de workers
	fmt.Println("findconc : ", time.Since(d))
	d = time.Now()
	MapConc(tab7, add1)
	fmt.Println("mapconc : ", time.Since(d))
	
	// tab5 := make([]int, 2<<25)
	// tab6 := make([]int, 2<<25)
	// Fill(tab5, 1)
	// Fill(tab6, 1)
	// // CopyConc(tab5, tab6)
	// // ch := make(chan bool) //un channel non bufferisé necessite qu'une goroutine lise directement dessus en parallèle
	// // la fonction equal conc necessite de stocker la valeur dans le channel pendant un temps car 
	// // tab6[25000000] = 25
	// d = time.Now()
	// ch2 := make(chan bool, nbworkers + 1)
	// ch3 := make(chan bool, nbworkers + 1)
	// fmt.Println(Equal(tab5, tab6))
	// fmt.Println("test equal : ", time.Since(d))
	// d = time.Now()
	// EqualConc(tab5, tab6, ch2)
	// fmt.Println(<- ch2)
	// fmt.Println("test equal conc : ", time.Since(d))
	// d = time.Now()
	// FasterEqualConc(tab5, tab6, ch3)
	// fmt.Println(<- ch3)
	// fmt.Println("test faster equal conc : ", time.Since(d))
}
