package comsoc

//une fonction de vote par fichier

import (
	"fmt"
)

type Alternative int           //les personnes pour lesquelles ont peut voter
type Profile [][]Alternative   //profils de préférences, les préférences sont un ordre (totale) et sont aussi transitives (0 meilleur alternative / n-1 pire)
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
func maxCount(count Count) (bestAlts []Alternative) {
	var max int
	for _, field := range count {
		if max < field {
			max = field
		}
	}

	for key, field := range count {
		if field == max {
			bestAlts = append(bestAlts, key)
		}
	}

	return
}

// vérifie les préférences d'un agent, par ex. qu'ils sont tous complets et que chaque alternative n'apparaît qu'une seule fois
func checkPrefs(prefs []Alternative, alts []Alternative) error {

}

// // vérifie le profil donné, par ex. qu'ils sont tous complets et que chaque alternative de alts apparaît exactement une fois par préférences
// func checkProfileAlternative(prefs Profile, alts []Alternative) error


// le swf part d'un profil et renvoie un décompte (renvoie un ordre sur l'ensemble des alternatives pour un profil donné)
// on part du profil qui est l'ensemble des préférences à une préférence collective

// la scf 

//approval threshold, le seuil pour l'agent 2 : [2]

//tie breaks

// fonction tiebreak : prend l'alternative préférée parmis un ensemble d'alternatives
// fonction factory tie break : on peut créer un vecteur qui donne l'ordre strict des alternatives en cas d'égalité et on prend celui qui a l'indice le plus petit
// ce factory va créer des fonctions qui étant donné un vecteur qui est le vecteur de tiebreak en cas d'égalité renvoie la fonction correspondante

//swf/scf factory

func main() {

	// Définition des alternatives (candidats/options)
	alt1 := Alternative(1) // Alice
	alt2 := Alternative(2) // Bob
	alt3 := Alternative(3) // Charlie
	alt4 := Alternative(4) // Diana

	// Liste de toutes les alternatives
	alternatives := []Alternative{alt1, alt2, alt3, alt4}

	// Préférences individuelles d'agents
	// Agent 1: Alice > Bob > Charlie > Diana
	prefs1 := []Alternative{alt1, alt2, alt3, alt4}

	// Agent 2: Bob > Charlie > Alice > Diana
	prefs2 := []Alternative{alt2, alt3, alt1, alt4}

	// Agent 3: Charlie > Alice > Diana > Bob
	prefs3 := []Alternative{alt3, alt1, alt4, alt2}

	// Agent 4: Diana > Bob > Alice > Charlie
	prefs4 := []Alternative{alt4, alt2, alt1, alt3}

	// Agent 5: Alice > Charlie > Bob > Diana
	prefs5 := []Alternative{alt1, alt3, alt2, alt4}

	// Profil complet (toutes les préférences)
	profile := Profile{prefs1, prefs2, prefs3, prefs4, prefs5}

	// Exemples de décomptes pour différents systèmes de vote
	// Décompte 1: vote de Borda
	count1 := Count{
		alt1: 12, // Alice
		alt2: 8,  // Bob
		alt3: 10, // Charlie
		alt4: 5,  // Diana
	}

	// Décompte 2: vote majoritaire
	count2 := Count{
		alt1: 2, // Alice
		alt2: 1, // Bob
		alt3: 1, // Charlie
		alt4: 1, // Diana
	}

	// Décompte 3: égalité
	count3 := Count{
		alt1: 3, // Alice
		alt2: 3, // Bob
		alt3: 2, // Charlie
		alt4: 2, // Diana
	}

	// Affichage des données de test
	fmt.Println("=== DONNÉES DE TEST ===")
	fmt.Println()

	fmt.Println("Alternatives:", alternatives)
	fmt.Println()

	fmt.Println("Préférences des agents:")
	for i, prefs := range profile {
		fmt.Printf("Agent %d: %v\n", i+1, prefs)
	}
	fmt.Println()

	fmt.Println("Exemple de décompte 1 (Borda):", count1)
	fmt.Println("Exemple de décompte 2 (Majoritaire):", count2)
	fmt.Println("Exemple de décompte 3 (Égalité):", count3)
	fmt.Println()

	// Tests avec la fonction rank
	fmt.Println("=== TESTS FONCTION RANK ===")
	fmt.Printf("Rang de Alice dans prefs1: %d\n", rank(alt1, prefs1))
	fmt.Printf("Rang de Bob dans prefs1: %d\n", rank(alt2, prefs1))
	fmt.Printf("Rang de Charlie dans prefs2: %d\n", rank(alt3, prefs2))
	fmt.Printf("Rang de Diana dans prefs3: %d\n", rank(alt4, prefs3))
	fmt.Println()

	// Données de test supplémentaires pour cas particuliers
	fmt.Println("=== CAS PARTICULIERS ===")

	// Préférences partielles (pour tester les erreurs)
	prefsIncomplete := []Alternative{alt1, alt3} // Manque alt2 et alt4
	fmt.Println("Préférences incomplètes:", prefsIncomplete)

	// Préférences avec doublons (pour tester les erreurs)
	prefsDuplicates := []Alternative{alt1, alt2, alt1, alt4} // alt1 répété
	fmt.Println("Préférences avec doublons:", prefsDuplicates)

	// Profil vide
	emptyProfile := Profile{}
	fmt.Println("Profil vide:", emptyProfile)

	// Test avec alternative inexistante
	altInexistante := Alternative(99)
	fmt.Printf("Rang d'une alternative inexistante: %d\n", rank(altInexistante, prefs1))

}
