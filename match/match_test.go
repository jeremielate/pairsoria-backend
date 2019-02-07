package match

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/uuid"
)

func TestMatch(t *testing.T) {
	list := []*Patient{
		&Patient{Id: uuid.New(), Genre: Male, Age: 18},
		&Patient{Id: uuid.New(), Genre: Female, Age: 27},
		&Patient{Id: uuid.New(), Genre: Female, Age: 34},
		&Patient{Id: uuid.New(), Genre: Male, Age: 79},
		&Patient{Id: uuid.New(), Genre: Unspecified, Age: 21},
		&Patient{Id: uuid.New(), Genre: Female, Age: 40},
		&Patient{Id: uuid.New(), Genre: Male, Age: 30},
		&Patient{Id: uuid.New(), Genre: Unspecified, Age: 56},
	}
	res := MatchPatients(list)

	for _, couple := range res {
		fmt.Println(couple.One, couple.Two)
	}
}

func TestMatchAge(t *testing.T) {
	list := []*Patient{
		&Patient{Id: uuid.New(), Genre: Unspecified, Age: 18},
		&Patient{Id: uuid.New(), Genre: Unspecified, Age: 41},
		&Patient{Id: uuid.New(), Genre: Unspecified, Age: 34},
		&Patient{Id: uuid.New(), Genre: Unspecified, Age: 35},
		&Patient{Id: uuid.New(), Genre: Unspecified, Age: 57},
		&Patient{Id: uuid.New(), Genre: Unspecified, Age: 40},
		&Patient{Id: uuid.New(), Genre: Unspecified, Age: 19},
		&Patient{Id: uuid.New(), Genre: Unspecified, Age: 56},
	}

	res := MatchPatients(list)

	for _, couple := range res {
		fmt.Println(couple.One, couple.Two)
		if math.Abs(float64(couple.One.Age-couple.Two.Age)) > 1 {
			t.Fail()
		}
	}
}
