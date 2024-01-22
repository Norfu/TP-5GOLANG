package personnage

import (
	"gojeu/types"
)

type Chevalier struct {
	Personnage
}

// Constructeur
func NewChevalier(Nom string, Pdv int, Attaque int, Defense int) IPersonnage {

	excalibur := types.Arme{DegatMax: 15, DistanceMax: 4}

	return &Chevalier{
		Personnage: Personnage{
			Nom:        Nom,
			Pdv:        Pdv,
			Attaque:    Attaque,
			Defense:    Defense,
			Arme:       []types.Arme{excalibur},
			Inventaire: []string{},
		},
	}

}
