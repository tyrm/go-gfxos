package gfxos

func mapRange(value int, fromLow int, fromHigh int, toLow int, toHigh int) int {
	return (((value - fromLow) * (toHigh - toLow)) / (fromHigh - fromLow)) + toLow
}
