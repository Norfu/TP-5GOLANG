package personnage

import (
	"gojeu/types"
)

type Gobelin struct {
	Personnage
}

// Constructeur
func NewGobelin(Nom string, Pdv int, Attaque int, Defense int) IPersonnage {

	massue := types.Arme{DegatMax: 12, DistanceMax: 4}

	return &Gobelin{
		Personnage: Personnage{
			Nom:        Nom,
			Pdv:        Pdv,
			Attaque:    Attaque,
			Defense:    Defense,
			Arme:       []types.Arme{massue},
			Inventaire: []string{},
		},
	}

}
