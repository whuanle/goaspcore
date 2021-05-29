package ioc

type ServiceLifetime int

const (
	Transient ServiceLifetime= iota
	Scoped
	Singleton
)
