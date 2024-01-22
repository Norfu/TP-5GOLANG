package enum

import "math/rand"

type Environnement string

const (
	Foret    Environnement = "FORET"
	Montagne Environnement = "MONTAGNE"
	Plaine   Environnement = "PLAINE"
	Ville    Environnement = "VILLE"
)

// func (s TypeCase) String() string {
// 	switch s {
// 	case Foret:
// 		return "foret"
// 	case Montagne:
// 		return "montagne"
// 	case Plaine:
// 		return "plaine"
// 	case Ville:
// 		return "ville"
// 	}
// 	return "unknown"
// }

//Fonction qui renvoie un environnement aléatoire
func (e Environnement) GenerateRandomEnvironnement() Environnement {
	random := rand.Intn(3)

	switch random {
	case 0:
		return Foret
	case 1:
		return Montagne
	case 2:
		return Plaine
	case 3:
		return Ville
	default:
		return "error"
	}
}
