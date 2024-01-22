package personnage

import (
	"gojeu/types"
)

type Elf struct {
	Personnage
}

// Constructeur
func NewElf(Nom string, Pdv int, Attaque int, Defense int) IPersonnage {

	// templatePersonnage := New(Nom, Pdv, Attaque, Defense)
	arc := types.Arme{DegatMax: 10, DistanceMax: 18}
	// templatePersonnage.Arme = append(templatePersonnage.Arme, arc)
	// elf := Elf{Personnage: templatePersonnage}

	return &Elf{
		Personnage: Personnage{
			Nom:        Nom,
			Pdv:        Pdv,
			Attaque:    Attaque,
			Defense:    Defense,
			Arme:       []types.Arme{arc},
			Inventaire: []string{},
		},
	}

}
