package personnage

import (
	"gojeu/types"
)

// Classe abstraite
type Personnage struct {
	Nom        string
	Pdv        int
	Attaque    int
	Defense    int
	Arme       []types.Arme
	Inventaire []string
}

func (p *Personnage) SetNom(nom string) {
	p.Nom = nom
}

func (p *Personnage) GetNom() string {
	return p.Nom
}

func (p *Personnage) SetPdv(pdv int) {
	p.Pdv = pdv
}

func (p *Personnage) GetPdv() int {
	return p.Pdv
}

func (p *Personnage) SetAttaque(attaque int) {
	p.Attaque = attaque
}

func (p *Personnage) GetAttaque() int {
	return p.Attaque
}
func (p *Personnage) SetDefense(defense int) {
	p.Defense = defense
}

func (p *Personnage) GetDefense() int {
	return p.Defense
}
func (p *Personnage) SetArme(arme types.Arme) {
	p.Arme = append(p.Arme, arme)
}

func (p *Personnage) GetArme() []types.Arme {
	return p.Arme
}
func (p *Personnage) SetInventaire(objet string) {
	p.Inventaire = append(p.Inventaire, objet)
}

func (p *Personnage) GetInventaire() []string {
	return p.Inventaire
}

func GetPersonnage(personnageType string) (IPersonnage, error) {
	switch personnageType {
	case "chevalier":
		return NewChevalier("Aragorn", 100, 15, 20), nil
	case "elf":
		return NewElf("Legolas", 70, 20, 10), nil
	case "gobelin":
		return NewGobelin("Smirgold", 60, 15, 10), nil
	case "magicien":
		return NewMagicien("Gandalf", 70, 10, 10), nil
	case "nain":
		return NewNain("Gimli", 100, 10, 30), nil
	case "orc":
		return NewOrc("José", 120, 20, 5), nil
	default:
		return nil, nil
	}
}

// func CreatePersonnage(NomTypePerso string) ActionPersonnage {
// 	var personnage ActionPersonnage
// 	switch NomTypePerso {
// 	case "Elf":
// 		personnage = NewElf("José", 5, 5, 5)

// 	}
// 	Arme := make([]types.Arme, 0)
// 	Inventaire := make([]string, 0)
// 	poing := types.Arme{DegatMax: 5, DistanceMax: 2}
// 	personnage := Personnage{Nom, Pdv, Attaque, Defense, append(Arme, poing), Inventaire}
// 	return personnage
// }

func (p *Personnage) Parler(interlocuteur IPersonnage) {
	println(p.Nom, " parle avec ", interlocuteur.GetNom())
}

func (p *Personnage) Marcher(caseJeu types.Case) {
	if caseJeu.Occupe {
		println("Impossible il y a deja quelqu'un")
	} else {
		println(p.Nom, "marche jusqu'a la case")
		caseJeu.Occupe = true
	}

}

func (p *Personnage) Boire() {
	print(p.Nom, " boit")
}

func (p *Personnage) Manger() {
	print(p.Nom, " mange et gagne 5 pdv")
	p.Pdv += 5
}

func (p *Personnage) Attaquer(victime IPersonnage) {

	armeActuelle := p.GetArme()[0]

	degat := p.GetAttaque() + armeActuelle.DegatMax

	switch {
	case victime.GetDefense() == 0:
		victime.SetPdv(victime.GetPdv() - degat)
	case victime.GetDefense() < degat:
		degat -= victime.GetDefense()
		victime.SetDefense(0)
		victime.SetPdv(victime.GetPdv() - degat)
	case victime.GetDefense() > degat:
		victime.SetDefense(victime.GetDefense() - degat)
	default:
		println("L'attaque a échouée")
	}

	println(p.GetNom(), " a fait ", degat, " a ", victime.GetNom())
	println("Armure de ", victime.GetNom(), ":", victime.GetDefense())
	println("Pdv de ", victime.GetNom(), ":", victime.GetPdv())

}
