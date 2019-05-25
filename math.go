package selection

// BytesWithinTolerance returns whether the abs of a-b is
// within some tolerance
func BytesWithinTolerance(a byte, b byte, tolerance byte) bool {
	if a > b {
		return (a - b) <= tolerance
	}

	return (b - a) <= tolerance
}
