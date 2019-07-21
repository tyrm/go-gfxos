package gfxos

import (
	"fmt"
	"testing"
)

type testColor888 struct {
	red   int
	green int
	blue  int
	value int
}

var testColor888s = []testColor888{
	{255, 0, 0, 0xf800},
	{0, 255, 0, 0x7e0},
	{0, 0, 255, 0x1f},
	{255, 255, 0, 0xffe0},
	{0, 255, 255, 0x7ff},
	{255, 0, 255, 0xf81f},
	{255, 255, 255, 0xffff},
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
