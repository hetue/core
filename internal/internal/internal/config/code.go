package config

type Code struct {
	Ok     uint8
	Failed uint8
}

func NewCode() *Code {
	return &Code{
		Ok:     0,
		Failed: 1,
	}
}
