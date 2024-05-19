package cinmemory

const (
	Non Strategy = iota
)

// Strategy it's the enum of the caching strategy.
//   - const Non = 0
type Strategy uint8

// IsValid validates the value of the strategy.
func (s Strategy) IsValid() bool {
	return s == Non
}

// String converts the current value of the type to a string.
func (s Strategy) String() string {
	if s == Non {
		return "Non"
	}

	return ""
}
