package math

import "testing"

type testpair struct {
	values []float64
	value  float64
}

var testsMin = []testpair{
	{[]float64{}, 0},
	{[]float64{1}, 1},
	{[]float64{1, 2}, 1},
	{[]float64{5, 1, 4, 2, 3}, 1},
	{[]float64{1, -1, 0}, 1},
}

func TestMin(t *testing.T) {
	for _, pair := range testsMin {
		v := Min(pair.values)
		if v != pair.value {
			t.Error(
				"For", pair.values,
				"expected", pair.value,
				"got", v,
			)
		}
	}
}

var testsMax = []testpair{
	{[]float64{}, 0},
	{[]float64{1}, 1},
	{[]float64{1, 2}, 2},
	{[]float64{1, 5, 2, 4, 3}, 5},
	{[]float64{-1, 1, 0}, 1},
}

func TestMax(t *testing.T) {
	for _, pair := range testsMax {
		v := Max(pair.values)
		if v != pair.value {
			t.Error(
				"For", pair.values,
				"expected", pair.value,
				"got", v,
			)
		}
	}
}
