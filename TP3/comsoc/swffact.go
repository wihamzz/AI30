package comsoc

import (
	"sort"
)

//ici, on va pouvoir donner en param n'importe quelle méthode swf et n'importe quelle tiebreak, et la factory va retourner une fonction SWF
//on peut donc personnaliser la fonction qu'on va avoir en retour, en fonction d'un type de méthode SWF et d'un tie break
//on va faire une méthode swf généraliste (qui peut utiliser différentes méthodes) avec un tie break qui empechera les égalités
//pour cette fonction, il faut par contre recréer un tableau avec l'ordre final et total sans égalité
func SWFFactory(swf func (p Profile) (Count, error), tb func ([]Alternative) (Alternative, error)) (func(Profile) ([]Alternative, error)) {
	return func(p Profile) (alts []Alternative, err error) { //on retourne un tableau d'alternatives ici car c'est la fonction de bien etre social (on retourne tous les candidats et leurs scores)
		count, err := swf(p)
		if err != nil {
			return
		}
		var bestAlt Alternative
		tempAlts := maxCount(count)
		if len(tempAlts) > 1 {
			bestAlt, err = tb(tempAlts)
			if err != nil {
				return
			}
		} else {
			bestAlt = tempAlts[0]
		}

		count_keys := sortCountKeysByValue(count)
		
		alts = make([]Alternative, 0)
		alts[0] = bestAlt
		for _,v := range count_keys {
			if v != bestAlt {
				alts = append(alts, v)
			}
		}
		return
	}
}

func sortCountKeysByValue(c Count) []Alternative{ // il n'y a pas d'ordre dans les maps, on ne peut donc pas les trier (elles changent d'ordre en fonction des itérations)
	keys := make([]Alternative, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { //less est une fonction de comparaison si j'ai bien compris
		return c[keys[i]] > c[keys[j]] //on trie les clés en fonction de leur valeur 
	})
	return keys
}
