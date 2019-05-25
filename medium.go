package selection

// Medium is a medium that can be clicked on
type Medium interface {
	Touch(x int, y int) error
}
