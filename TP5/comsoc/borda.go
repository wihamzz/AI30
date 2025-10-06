package comsoc

// import "fmt"

// "fmt"

func BordaSWF(p Profile) (count Count, err error) {
	err = checkProfileAlternative(p, getAlternatives(p))
	if  err != nil {
		return count, err
	}
	count = make(Count)

	for _, voter := range p{
		for j, val := range voter {
			count[val] += len(voter) - j - 1
		}
	}

	return count, nil
}

func BordaSCF(p Profile) (bestAlts []Alternative, err error) {
	err = checkProfileAlternative(p, getAlternatives(p))
	if  err != nil {
		return bestAlts, err
	}

	count := make(Count)
	for _, voter := range p{
		for j, val := range voter {
			count[val] += len(voter) - j - 1
		}
	}
	return MaxCount(count), nil
}
