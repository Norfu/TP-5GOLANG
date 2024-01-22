package jeu

import (
	"gojeu/personnage"
	"gojeu/types"
)

type Jeu struct {
	Carte       types.Carte
	Personnages []personnage.IPersonnage
}

func (j *Jeu) Init() {
	j.Carte = types.NewCarte(1000, 1000)

	aragorn, _ := personnage.GetPersonnage("chevalier")
	gandalf, _ := personnage.GetPersonnage("magicien")

	j.Personnages = append(j.Personnages, aragorn, gandalf)

	j.Carte.Cases[0][1].Occupe = true
	j.Carte.Cases[1][0].Occupe = true

}

func (j *Jeu) Play() {
	j.Init()
	j.Personnages[0].Attaquer(j.Personnages[1])
}
