package match

import (
	"math"

	"github.com/cpmech/gosl/graph"
)

type Mate struct {
	value float64
}

func (mate *Mate) Preference(other *Mate) float64 {
	if mate == other {
		return math.MaxFloat64
	}
	return math.Abs(mate.value - other.value)
}

func Match(mates []*Mate) [][2]*Mate {
	var mnk graph.Munkres
	length := len(mates)
	mnk.Init(length, length)

	matrix := make([][]float64, 0, length)
	for _, mateX := range mates {
		row := make([]float64, length)
		for y, mateY := range mates {
			row[y] = mateX.Preference(mateY)
		}
		matrix = append(matrix, row)
	}

	mnk.SetCostMatrix(matrix)
	mnk.Run()

	res := make([][2]*Mate, len(mnk.Links))
	for x, y := range mnk.Links {
		res[x] = [2]*Mate{mates[x], mates[y]}
	}

	return res
}
