package boards

func mapToState(coord string) (int, int) {
	if len(coord) > 2 {
		return int(coord[0] - 65), 9
	}
	x := int(coord[0] - 65)
	y := int(coord[1] - 49)
	return x, y
}
