package main

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestExampleOne(t *testing.T) {
	input := "bbabb"

	g := Grammar{
		Productions: []Production{
			Production{
				Left: "S",
				Right: [][]string{
					[]string{"A", "B"},
				},
			},
			Production{
				Left: "A",
				Right: [][]string{
					[]string{"B", "B"},
					[]string{"a"},
				},
			},
			Production{
				Left: "B",
				Right: [][]string{
					[]string{"A", "B"},
					[]string{"b"},
				},
			},
		},
	}

	val, err := CYK(input, g)
	assert.Equal(t, err, nil)
	assert.Equal(t, val, false)
}

func TestExampleTwo(t *testing.T) {
	input := "baaba"

	g := Grammar{
		Productions: []Production{
			Production{
				Left: "S",
				Right: [][]string{
					[]string{"A", "B"},
					[]string{"B", "C"},
				},
			},
			Production{
				Left: "A",
				Right: [][]string{
					[]string{"B", "A"},
					[]string{"a"},
				},
			},
			Production{
				Left: "B",
				Right: [][]string{
					[]string{"C", "C"},
					[]string{"b"},
				},
			},
			Production{
				Left: "C",
				Right: [][]string{
					[]string{"A", "B"},
					[]string{"a"},
				},
			},
		},
	}

	val, err := CYK(input, g)
	assert.Equal(t, err, nil)
	assert.Equal(t, val, true)
}

func TestMatrixComputations(t *testing.T) {
	val := matrixmultiply([][][]string{
		[][]string{
			[]string{"A", "B"},
			[]string{"B", "C"},
		},
	})

	assert.Equal(t, val, [][]string{
		[]string{"A", "B"},
		[]string{"A", "C"},
		[]string{"B", "B"},
		[]string{"B", "C"},
	})

	val2 := matrixfromstring("abcd")

	assert.Equal(t, val2, [][]string{
		[]string{"a", "bcd"},
		[]string{"ab", "cd"},
		[]string{"abc", "d"},
	})

	val3 := matrixfromstring("this is a triumph")
	val4 := matrixfromstring("this is a triumph")
	val5 := matrixfromstring("this is a disaster")

	assert.Equal(t, val3, val4)
	assert.NotEqual(t, val3, val5)
}
