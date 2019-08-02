package gfxos

func toUint(input int) int {
	if input < 0 {
		input = input + 256
	}

	return input
}

func toUint12(input int) int {
	if input < 0 {
		input = input + 4096
	}

	return input
}