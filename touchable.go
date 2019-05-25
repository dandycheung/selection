package selection

// Touchable is a selection that can be clicked on
type Touchable interface {
	Touch(x int, y int, rand Random) error
}
