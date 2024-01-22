package personnage

import (
	"gojeu/types"
)

type Orc struct {
	Personnage
}

// Constructeur
func NewOrc(Nom string, Pdv int, Attaque int, Defense int) IPersonnage {

	poing := types.Arme{DegatMax: 5, DistanceMax: 2}

	return &Orc{
		Personnage: Personnage{
			Nom:        Nom,
			Pdv:        Pdv,
			Attaque:    Attaque,
			Defense:    Defense,
			Arme:       []types.Arme{poing},
			Inventaire: []string{},
		},
	}
}
