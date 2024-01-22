package personnage

import "gojeu/types"

type IPersonnage interface {
	SetNom(nom string)
	SetPdv(pdv int)
	SetAttaque(attaque int)
	SetDefense(defense int)
	SetArme(arme types.Arme)
	SetInventaire(objet string)
	GetNom() string
	GetPdv() int
	GetAttaque() int
	GetDefense() int
	GetArme() []types.Arme
	GetInventaire() []string
	Parler(IPersonnage)
	Marcher(types.Case)
	Boire()
	Manger()
	Attaquer(IPersonnage)
}
