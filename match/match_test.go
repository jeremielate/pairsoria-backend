package match

import (
	"fmt"
	"testing"
)

func TestMatch(t *testing.T) {
	list := []*Mate{
		&Mate{4},
		&Mate{9},
		&Mate{34},
		&Mate{124},
		&Mate{13},
		&Mate{7},
		&Mate{213},
		&Mate{56},
	}
	res := Match(list)

	for _, couple := range res {
		fmt.Println(couple[0].value, couple[1].value)
	}
}
