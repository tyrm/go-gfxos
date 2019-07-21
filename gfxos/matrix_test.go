package gfxos

import (
	"fmt"
	"testing"
)


type testColor888 struct {
	red int
	green int
	blue int
	value int
}

var testColor888s = []testColor888{
	{ 255,0,0,0xf800},
	{ 0,255,0,0x7e0},
	{ 0,0,255,0x1f},
	{ 255,255,0,0xffe0},
	{ 0,255,255,0x7ff},
	{ 255,0,255,0xf81f},
	{ 255,255,255,0xffff},
}

func TestColr888(t *testing.T) {
	for _, testVal := range testColor888s {

		v := color888(testVal.red, testVal.green, testVal.blue)
		if v != testVal.value {
			t.Error(
				"For [",
				fmt.Sprintf("%x", testVal.red),
				fmt.Sprintf("%x", testVal.green),
				fmt.Sprintf("%x", testVal.blue),
				"] expected",
				fmt.Sprintf("%x", testVal.value),
				"got",
				fmt.Sprintf("%x", v),
			)
		}
	}
}

type testmap struct {
	fromLow int
	fromHigh int
	toLow int
	toHigh int
	value int
	result int
}

var testMapRanges = []testmap{
	{ 0,255,0,31,0,0},
	{ 0,255,0,31,63,7},
	{ 0,255,0,31,127,15},
	{ 0,255,0,31,171,20},
	{ 0,255,0,31,246,29},
	{ 0,255,0,31,247,30},
	{ 0,255,0,31,254,30},
	{ 0,255,0,31,255,31},
	{ 0,255,0,127,0,0},
	{ 0,255,0,127,63,31},
	{ 0,255,0,127,127,63},
	{ 0,255,0,127,171,85},
	{ 0,255,0,127,246,122},
	{ 0,255,0,127,247,123},
	{ 0,255,0,127,254,126},
	{ 0,255,0,127,255,127},
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