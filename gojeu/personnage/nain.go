package personnage

import (
	"gojeu/types"
)

type Nain struct {
	Personnage
}

// Constructeur
func NewNain(Nom string, Pdv int, Attaque int, Defense int) IPersonnage {

	hache := types.Arme{DegatMax: 17, DistanceMax: 3}

	return &Nain{
		Personnage: Personnage{
			Nom:        Nom,
			Pdv:        Pdv,
			Attaque:    Attaque,
			Defense:    Defense,
			Arme:       []types.Arme{hache},
			Inventaire: []string{},
		},
	}

}
