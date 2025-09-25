package comsoc

import (

)

//chaque profil vote pour les candidats qu'il veut, il peut voter pour tout le monde ou encore pour personne. Le threshold permet de caractériser le nombre de candidats pour lequel le profile vote
func ApprovalSWF(p Profile, thresholds []int) (count Count, err error) {
	err = checkProfileAlternative(p, getAlternatives(p))
	if err != nil {
		return count, err
	}
	count = make(Count)
	for i, voter := range p { //je suis dans le tableau de préférences pour un profil par ex : {1,3,2}, si le threshold est a 2 pour ce profil, j'ajoute 1,3 à count
		th := thresholds[i]
		for _, field := range voter {
			if th == 0 {
				break
			}
			count[field]++
			th--
		}
	}
	return count, nil
}

func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error) {
	err = checkProfileAlternative(p, getAlternatives(p))
	if err != nil {
		return bestAlts, err
	}
	count := make(Count)
	for i, voter := range p { //je suis dans le tableau de préférences pour un profil par ex : {1,3,2}, si le threshold est a 2 pour ce profil, j'ajoute 1,3 à count
		th := thresholds[i]
		for _, field := range voter {
			if th == 0 {
				break
			}
			count[field]++
			th--
		}
	}
	return maxCount(count), nil

}
