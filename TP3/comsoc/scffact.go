package comsoc

import (

)

func SCFFactory(scf func (p Profile) ([]Alternative, error), tb func ([]Alternative) (Alternative, error)) (func(Profile) (Alternative, error)) {
	return func(p Profile) (Alternative, error) {
		alts, err := scf(p)
		if err != nil {
			return -1, err
		}
		var alt Alternative
		if len(alts) > 1 {
			alt, err = tb(alts) 
			if err != nil {
				return -1, err
			}
		} else {
			alt = alts[0]
		}
		return alt, nil
	}
}