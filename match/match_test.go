package match

import (
	"fmt"
	"math"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestMatch(t *testing.T) {
	list := []*Patient{
		&Patient{id: uuid.NewV4(), genre: Male, age: 18},
		&Patient{id: uuid.NewV4(), genre: Female, age: 27},
		&Patient{id: uuid.NewV4(), genre: Female, age: 34},
		&Patient{id: uuid.NewV4(), genre: Male, age: 79},
		&Patient{id: uuid.NewV4(), genre: Unspecified, age: 21},
		&Patient{id: uuid.NewV4(), genre: Female, age: 40},
		&Patient{id: uuid.NewV4(), genre: Male, age: 30},
		&Patient{id: uuid.NewV4(), genre: Unspecified, age: 56},
	}
	res := MatchPatients(list)

	for _, couple := range res {
		fmt.Println(couple[0], couple[1])
	}
}

func TestMatchAge(t *testing.T) {
	list := []*Patient{
		&Patient{id: uuid.NewV4(), genre: Unspecified, age: 18},
		&Patient{id: uuid.NewV4(), genre: Unspecified, age: 41},
		&Patient{id: uuid.NewV4(), genre: Unspecified, age: 34},
		&Patient{id: uuid.NewV4(), genre: Unspecified, age: 35},
		&Patient{id: uuid.NewV4(), genre: Unspecified, age: 57},
		&Patient{id: uuid.NewV4(), genre: Unspecified, age: 40},
		&Patient{id: uuid.NewV4(), genre: Unspecified, age: 19},
		&Patient{id: uuid.NewV4(), genre: Unspecified, age: 56},
	}

	res := MatchPatients(list)

	for _, couple := range res {
		fmt.Println(couple[0], couple[1])
		if math.Abs(float64(couple[0].age-couple[1].age)) > 1 {
			t.Fail()
		}
	}
}
