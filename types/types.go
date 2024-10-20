package types

type OsProvider interface {
	GetOs() string
	GetHomeDir() (string, error)
}
