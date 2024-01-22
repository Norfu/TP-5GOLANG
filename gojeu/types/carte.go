package types

type Carte struct {
	Longueur int
	Largeur  int
	Cases    [][]Case
}

func NewCarte(Longueur int, Largeur int) Carte {

	carte := Carte{
		Longueur: Longueur,
		Largeur:  Largeur,
		Cases:    make([][]Case, Longueur),
	}

	for i := range carte.Cases {
		carte.Cases[i] = make([]Case, Largeur)

		for j := range carte.Cases[i] {
			carte.Cases[i][j] = NewCase()
		}
	}

	return carte

}
