package logger

const (
	Non Level = iota
	Debug
	Info
	Warn
	Error
)

// Level it's a logger level enum.
//   - const Non   = 0
//   - const Debug = 1
//   - const Info  = 2
//   - const Warn  = 3
//   - const Error = 4
type Level uint8

// IsValid validates the value of the logger level.
func (l Level) IsValid() bool {
	return l == Debug || l == Info || l == Warn || l == Error
}

// String converts the current value of the level to a string.
func (l Level) String() string {
	switch l {
	case Non:
		return ""
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	}

	return ""
}
