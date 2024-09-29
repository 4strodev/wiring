package wiring

type abstractionLifeCycle uint8

const (
	SINGLETON abstractionLifeCycle = iota
	TRANSIENT
)
