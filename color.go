package gfxos

func color888(r int, g int, b int) int {
	new_r := mapRange(r, 0, 255, 0, 31)
	new_g := mapRange(g, 0, 255, 0, 63)
	new_b := mapRange(b, 0, 255, 0, 31)

	new_g = new_g << 5
	new_r = new_r << 11

	color := new_r | new_g | new_b
	return color
}
