package base

type Library interface {
	Start() error

	AsyncStart() error

	Restart() error

	Stop() error

	Status() (string, error)

	Version() string
}
