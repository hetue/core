package kernel

type Step interface {
	Runnable() bool

	Run() error
}
