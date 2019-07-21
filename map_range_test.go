package gfxos

import "testing"

type testmap struct {
	fromLow  int
	fromHigh int
	toLow    int
	toHigh   int
	value    int
	result   int
}

var testMapRanges = []testmap{
	{0, 255, 0, 31, 0, 0},
	{0, 255, 0, 31, 63, 7},
	{0, 255, 0, 31, 127, 15},
	{0, 255, 0, 31, 171, 20},
	{0, 255, 0, 31, 246, 29},
	{0, 255, 0, 31, 247, 30},
	{0, 255, 0, 31, 254, 30},
	{0, 255, 0, 31, 255, 31},
	{0, 255, 0, 127, 0, 0},
	{0, 255, 0, 127, 63, 31},
	{0, 255, 0, 127, 127, 63},
	{0, 255, 0, 127, 171, 85},
	{0, 255, 0, 127, 246, 122},
	{0, 255, 0, 127, 247, 123},
	{0, 255, 0, 127, 254, 126},
	{0, 255, 0, 127, 255, 127},
}

func TestMapRange(t *testing.T) {
	for _, testVal := range testMapRanges {

		v := mapRange(testVal.value, testVal.fromLow, testVal.fromHigh, testVal.toLow, testVal.toHigh)
		if v != testVal.result {
			t.Error(
				"For", testVal.value,
				"expected", testVal.result,
				"got", v,
			)
		}
	}
}
