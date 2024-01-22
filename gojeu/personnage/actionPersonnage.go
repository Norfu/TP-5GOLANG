package personnage

import (
	"gojeu/types"
)

type ActionPersonnage interface {
	Parler(Personnage)
	Marcher(types.Case)
	Boire()
	Manger()
	Attaquer(Personnage)
}
