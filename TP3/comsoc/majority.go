package comsoc

//Renvoie un décompte à partir d'un profil (Candidat 1 : 3 voix, candidat 2 : 2 voix, ...)
func MajoritySWF(p Profile) (count Count, err error) {
	err = checkProfileAlternative(p, getAlternatives(p))
	if  err != nil {
		return count, err
	}
	count = make(Count) //obligatoire de make la map avant de faire des choses dessus
	for _, voter := range p {
		count[voter[0]] ++
	}
	return count, nil
}

//Renvoie uniquement la ou les alternative(s) préférée(s)
func MajoritySCF(p Profile) (bestAlts []Alternative, err error) {
	err = checkProfileAlternative(p, getAlternatives(p))
	if  err != nil {
		return bestAlts, err
	}
	count := make(Count) 
	for _, voter := range p {
		count[voter[0]] ++
	}
	return maxCount(count), nil
}
