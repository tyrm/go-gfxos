package gfxos

import "testing"

type testToUint struct {
	input int
	value int
}

var testToUints = []testToUint{
	{-128, 128},
	{-1, 255},
	{0, 0},
	{127, 127},
}

func TestToUint(t *testing.T) {
	for _, testVal := range testToUints {

		v := toUint(testVal.input)
		if v != testVal.value {
			t.Error(
				"For", testVal.input,
				"expected",
				testVal.value,
				"got", v,
			)
		}
	}
}
