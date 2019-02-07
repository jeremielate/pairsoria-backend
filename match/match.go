package match

import (
	"math"

	"github.com/cpmech/gosl/graph"
	uuid "github.com/satori/go.uuid"
)

type Genre uint

const (
	Male Genre = iota
	Female
	Unspecified
)

type Patient struct {
	id     uuid.UUID
	age    int
	genre  Genre
	points uint64
}

// calcule la proximité d'un patient avec un autre selon des critères définis.
// TODO: à ameliorer pour un meilleur matching
func patientProximity(a, b *Patient) float64 {
	if a.id == b.id {
		// TODO: don't match with yourself
		return math.MaxFloat64
	}
	ageDiff := math.Abs(float64(a.age) - float64(b.age))
	pointsDiff := math.Abs(float64(a.points) - float64(b.points))

	var difference float64
	if a.genre != b.genre {
		// valeur prise au pif
		difference += 20
	}

	return pointsDiff + ageDiff + difference
}

func matchPatients(patients []*Patient, compare func(a, b *Patient) float64) [][2]*Patient {
	var mnk graph.Munkres

	length := len(patients)
	mnk.Init(length, length)

	matrix := make([][]float64, 0, length)
	for _, patientX := range patients {
		row := make([]float64, length)
		for y, patientY := range patients {
			row[y] = compare(patientX, patientY)
		}
		matrix = append(matrix, row)
	}

	mnk.SetCostMatrix(matrix)
	mnk.Run()

	res := make([][2]*Patient, len(mnk.Links))
	for x, y := range mnk.Links {
		res[x] = [2]*Patient{patients[x], patients[y]}
	}

	return res
}

func MatchPatients(patients []*Patient) [][2]*Patient {
	return matchPatients(patients, patientProximity)
}
