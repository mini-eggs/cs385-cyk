package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func readfile(filename string) (g Grammar, err error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &g)
	if err != nil {
		return
	}

	return
}

func TestExampleOne(t *testing.T) {
	// should fail
	input := "bbabb"
	grammar, err := readfile("example_input_one.json")
	assert.Equal(t, err, nil)
	val, err := CYK(input, grammar)
	assert.Equal(t, err, nil)
	assert.Equal(t, val, false)
}

func TestExampleTwo(t *testing.T) {
	// should pass
	input := "baaba"
	grammar, err := readfile("example_input_two.json")
	assert.Equal(t, err, nil)
	val, err := CYK(input, grammar)
	assert.Equal(t, err, nil)
	assert.Equal(t, val, true)
}

func TestExampleThree(t *testing.T) {
	// should pass
	input := "aabd"
	grammar, err := readfile("example_input_three.json")
	assert.Equal(t, err, nil)
	val, err := CYK(input, grammar)
	assert.Equal(t, err, nil)
	assert.Equal(t, val, true)
	// should fail
	input = "adbd"
	grammar, err = readfile("example_input_three.json")
	assert.Equal(t, err, nil)
	val, err = CYK(input, grammar)
	assert.Equal(t, err, nil)
	assert.Equal(t, val, false)
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
