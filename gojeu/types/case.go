package types

import (
	"gojeu/enum"
)

type Case struct {
	TypeCase enum.Environnement
	Occupe   bool
}

func NewCase() Case {

	var envi enum.Environnement
	Case := Case{TypeCase: envi.GenerateRandomEnvironnement(), Occupe: false}

	return Case
}
