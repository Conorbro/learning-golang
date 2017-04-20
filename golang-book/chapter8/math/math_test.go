package math

import "testing"

type testpairMax struct {
	values []float64
	max    float64
}

type testpairMin struct {
	values []float64
	min    float64
}

var minTests = []testpairMin{
	{[]float64{1, 2, 3, 4}, 1},
	{[]float64{5, 6, 7, 8}, 5},
}

var maxTests = []testpairMax{
	{[]float64{1, 2, 3, 4}, 4},
	{[]float64{5, 6, 7, 8}, 8},
}

func TestAverage(t *testing.T) {
	v := Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

func TestMax(t *testing.T) {
	for _, pair := range maxTests {
		res := Max(pair.values)
		if res != pair.max {
			t.Error(
				"For", pair.values,
				"Expected", pair.max,
				"got", res,
			)
		}

	}
}

func TestMin(t *testing.T) {
	for _, pair := range minTests {
		res := Min(pair.values)
		if res != pair.min {
			t.Error(
				"For", pair.values,
				"Expected", pair.min,
				"Got", res,
			)
		}
	}
}
