package personnage

import (
	"gojeu/types"
)

type Magicien struct {
	Personnage
}

// Constructeur
func NewMagicien(Nom string, Pdv int, Attaque int, Defense int) IPersonnage {

	batonMagique := types.Arme{DegatMax: 8, DistanceMax: 10}

	return &Magicien{
		Personnage: Personnage{
			Nom:        Nom,
			Pdv:        Pdv,
			Attaque:    Attaque,
			Defense:    Defense,
			Arme:       []types.Arme{batonMagique},
			Inventaire: []string{},
		},
	}

}
