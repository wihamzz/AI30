package comsoc

import (
	"gonum.org/v1/gonum/stat/combin"
)

//idée : 

/*
	Générer les 2-combinaisons parmis l'ensemble des alternatives
	Parcourir le profil
	Pour chaque sous profil, récupérer l'indice des deux éléments de chaque combinaison, et les comparer
	Dans une map contenant chaque alternative, ajouter + 1 à celui qui a le plus petit indice
	Comparer les scores via maxcount
	Si len(maxcount(..)) > 1 retourner un tableau vide, sinon retourner un tableau avec le gagnant en position 0 
*/

//un gagnant de condorcet est une alternative qui gagne contre toutes les autres en duel
func CondorcetWinner(p Profile) (bestAlts []Alternative, err error) {
	err = checkProfileAlternative(p, getAlternatives(p))
	if  err != nil {
		return bestAlts, err
	}

	temp_alts := getAlternatives(p)
	combinations := combin.Combinations(len(temp_alts), 2)
	alts := make([][]Alternative, 0)
	temp_sl := make([]Alternative, 0)

	for _,v := range combinations {
		for _,w := range v {
			temp_sl = append(temp_sl, Alternative(w) + 1)
		}
		alts = append(alts, temp_sl)
		temp_sl = nil
	}

	total_score := make(Count)
	intermediate_score := make(Count)
	index := make(Count)

	for _, oneversusone := range alts {
		for _, subprofile := range p {
			for j, v := range subprofile {
				switch {
				case v == oneversusone[0]:
					index[oneversusone[0]] = j
				case v == oneversusone[1]:
					index[oneversusone[1]] = j
				}
			}
			if index[oneversusone[0]] < index[oneversusone[1]] {
				intermediate_score[oneversusone[0]]++
			} else {
				intermediate_score[oneversusone[1]]++
			}
		}
		switch {
		case intermediate_score[oneversusone[0]] > intermediate_score[oneversusone[1]]:
			total_score[oneversusone[0]]++
		case intermediate_score[oneversusone[0]] < intermediate_score[oneversusone[1]]:
			total_score[oneversusone[1]]++
		case intermediate_score[oneversusone[0]] == intermediate_score[oneversusone[1]]:
			total_score[oneversusone[0]]++
			total_score[oneversusone[1]]++
		}
		intermediate_score[oneversusone[0]] = 0
		intermediate_score[oneversusone[1]] = 0
	}

	result := make([]Alternative, 0)
	for k, v := range total_score {
		if v == len(alts) - 1 { // est le gagnant de condorcet car gagne contre tous les autres
			result = append(result, k)
			return result, nil
		}
	}
	return result, nil
}


