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
	val := matrixmerge([][]string{
		[]string{"E"},
		[]string{"A", "B"},
		[]string{"B", "C"},
	})

	assert.Equal(t, val, [][]string{
		[]string{"E", "A", "B"},
		[]string{"E", "A", "C"},
		[]string{"E", "B", "B"},
		[]string{"E", "B", "C"},
	})

	val6 := matrixmerge([][]string{
		[]string{"A", "B", "C"},
		[]string{"M", "N"},
		[]string{"X", "Y"},
	})

	assert.Equal(t, val6, [][]string{
		[]string{"A", "M", "X"},
		[]string{"A", "M", "Y"},
		[]string{"A", "N", "X"},
		[]string{"A", "N", "Y"},
		[]string{"B", "M", "X"},
		[]string{"B", "M", "Y"},
		[]string{"B", "N", "X"},
		[]string{"B", "N", "Y"},
		[]string{"C", "M", "X"},
		[]string{"C", "M", "Y"},
		[]string{"C", "N", "X"},
		[]string{"C", "N", "Y"},
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
