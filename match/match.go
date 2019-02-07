package match

import (
	"encoding/json"
	"math"

	"github.com/cpmech/gosl/graph"
	"github.com/google/uuid"
)

type Genre uint

const (
	Male Genre = iota
	Female
	Unspecified
)

type Patient struct {
	Id     uuid.UUID
	Age    int    `json:"age"`
	Genre  Genre  `json:"genre"`
	Points uint64 `json:"points"`
}

type Couple struct {
	a, b *Patient
}

func (c Couple) MarshallJSON() ([]byte, error) {
	var couple [2]uuid.UUID
	couple[0] = c.a.Id
	couple[1] = c.a.Id
	return json.Marshal(couple)
}

// calcule la proximité d'un patient avec un autre selon des critères définis.
// TODO: à ameliorer pour un meilleur matching
func patientProximity(a, b *Patient) float64 {
	if a.Id == b.Id {
		// TODO: don't match with yourself
		return math.MaxFloat64
	}
	ageDiff := math.Abs(float64(a.Age) - float64(b.Age))
	pointsDiff := math.Abs(float64(a.Points) - float64(b.Points))

	var difference float64
	if a.Genre != b.Genre {
		// valeur prise au pif
		difference += 20
	}

	return pointsDiff + ageDiff + difference
}

func matchPatients(patients []*Patient, compare func(a, b *Patient) float64) []Couple {
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

	res := make([]Couple, len(mnk.Links))
	for x, y := range mnk.Links {
		res[x] = Couple{patients[x], patients[y]}
	}

	return res
}

func MatchPatients(patients []*Patient) []Couple {
	return matchPatients(patients, patientProximity)
}
