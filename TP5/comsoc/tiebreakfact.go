package comsoc

import (
	"fmt"
)

//on donne en parametre un vecteur (orderedAlts) qui spécifie l'ordre qu'il faudra utiliser pour selectionner les alternatives en cas d'égalité (un ordre que spécifiera l'appelant)
//cette fonction va alors renvoyer la fonction tiebreak correspondant spécifiquement au vecteur de tiebreak passé en paramètres à la factory
//l'avantage ici est qu'on peut personnaliser la fonction tiebreak retournée en fonction du orderedalts passé en param de la factory
//sans factory, on devrait tout le temps repasser les ordered alts à la fonction de tiebreak
//la fonction retournée par factory utilisera toujours les orderedalts qu'on a passé à la factory au début
//elle se "souvient" des orderedAlts grace à la closure (fonction retournée)
func TieBreakFactory(orderedAlts []Alternative) (func ([]Alternative) (Alternative, error)) {
	return func (bestAlts []Alternative) (alt Alternative, err error) { //une fonction de tiebreak possible
		if (len(orderedAlts) == 0) {
			alt = bestAlts[0]
			return 
		}

		if len(bestAlts) == 0 {
			alt = -1
			err = fmt.Errorf("alternatives slice is empty")
			return 
		}

		for _, v := range orderedAlts { //on parcourt les alternatives classées dans l'ordre croissant, et on teste si une de ces dernières est dans les bestAlts
			for _, w := range bestAlts {
				if v == w {
					alt = v
					return 
				}
			}
		} 
		alt = -1
		err = fmt.Errorf("orderedAlts and bestAlts have no elements in common") 
		return 
	}
}
