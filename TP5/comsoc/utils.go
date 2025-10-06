package comsoc

//une fonction de vote par fichier

import (
	"slices"
	"fmt"
)

type Alternative int           //les personnes pour lesquelles on peut voter
type Profile [][]Alternative   //profils de préférences, les préférences sont un ordre (total) et sont aussi transitives (0 meilleure alternative / n-1 pire)
type Count map[Alternative]int //décompte qui associe à chaque alternative un entier (un score) (renvoie le nombre de points qu'une alternative a obtenu)
// il peut y avoir des égalités

// renvoie l'indice ou se trouve alt dans prefs
func rank(alt Alternative, prefs []Alternative) int {
	for i, v := range prefs {
		if v == alt {
			return i
		}
	}
	return -1
}

// renvoie vrai ssi alt1 est préférée à alt2
func isPref(alt1, alt2 Alternative, prefs []Alternative) bool {
	return (rank(alt1, prefs) < rank(alt2, prefs)) && (rank(alt1, prefs) != -1) && (rank(alt1, prefs) != -1)
}

// renvoie les meilleures alternatives pour un décompte donné
func MaxCount(count Count) (bestAlts []Alternative) {
	var max int
	for _, f := range count {
		if max < f {
			max = f
		}
	}

	for k, f := range count {
		if f == max {
			bestAlts = append(bestAlts, k)
		}
	}

	return
}

func exists(al Alternative, sl []Alternative) bool {
	return slices.Contains(sl, al)
}

// vérifie les préférences d'un agent, par ex. qu'ils sont tous complets et que chaque alternative n'apparaît qu'une seule fois
func checkPrefs(prefs []Alternative, alts []Alternative) error { //préférence complete : toutes les alternatives sont dans le profil d'un agent
	if len(prefs) != len(alts) {
		return fmt.Errorf("len prefs != len alts")
	}

	m := make(Count)
	for _, v := range prefs {
		m[v]++
	}
	
	for _, f := range m {
		if f > 1 {
			return fmt.Errorf("alternatives appearing at least 2 times")
		}
	}

	for _, v := range alts {
		if !exists(v, prefs){
			return fmt.Errorf("prefs is not complete")
		}
	}
	
	return nil
}

// vérifie le profil donné, par ex. qu'ils sont tous complets et que chaque alternative de alts apparaît exactement une fois par préférences
func checkProfileAlternative(prefs Profile, alts []Alternative) error {
	for _, f := range prefs {
		err := checkPrefs(f, alts)
		if err != nil {
			return err
		}
	}
	return nil
}

func getAlternatives(p Profile) (alts []Alternative) {
	m := make(Count)
	for _, sub_p := range p {
		for _, alt := range sub_p {
			m[alt]++
		}
	}

	for k := range m {
		alts = append(alts, k)
	}

	return 
}


// le swf part d'un profil et renvoie un décompte (renvoie un ordre sur l'ensemble des alternatives pour un profil donné)
// on part du profil qui est l'ensemble des préférences à une préférence collective

// la scf 

//approval threshold, le seuil pour l'agent 2 : [2]

//tie breaks

// fonction tiebreak : prend l'alternative préférée parmis un ensemble d'alternatives
// fonction factory tie break : on peut créer un vecteur qui donne l'ordre strict des alternatives en cas d'égalité et on prend celui qui a l'indice le plus petit
// ce factory va créer des fonctions qui étant donné un vecteur qui est le vecteur de tiebreak en cas d'égalité renvoie la fonction correspondante

//swf/scf factory
