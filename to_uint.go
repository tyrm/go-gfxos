package gfxos

func toUint(input int) int {
	if input < 0 {
		input = input + 256
	}

	return input
}