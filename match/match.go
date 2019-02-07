package match

import (
	"encoding/json"
	"math"
	"strings"

	"github.com/cpmech/gosl/graph"
	"github.com/google/uuid"
)

type Genre uint

const (
	Male Genre = iota
	Female
	Unspecified
)

func (g Genre) MarshalJSON() ([]byte, error) {
	var s string
	switch g {
	case Female:
		s = "female"
	case Male:
		s = "male"
	default:
		s = "unspecified"
	}
	return json.Marshal(s)
}

func (g *Genre) UnmarshalJSON(b []byte) error {
	var s string

	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	case "female":
		*g = Female
	case "male":
		*g = Male
	default:
		*g = Unspecified
	}

	return nil
}

type Patient struct {
	Id     uuid.UUID `json:"id"`
	Age    int       `json:"age"`
	Genre  Genre     `json:"genre"`
	Points uint64    `json:"points"`
}

type Couple struct {
	One *Patient `json:"one"`
	Two *Patient `json:"two"`
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
