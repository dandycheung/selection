package selection

// Random is an interface that represents the subset of
// random functions necessary for selections
type Random interface {
	// Intn returns, as an int, a non-negative pseudo-random number in [0,n)
	Intn(n int) int
}
