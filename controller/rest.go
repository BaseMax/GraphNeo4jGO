package controller

type Rest interface {
	Start() error
	Stop() error
}
