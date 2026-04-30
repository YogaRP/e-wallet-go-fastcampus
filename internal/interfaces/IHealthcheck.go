package interfaces

type IHalthcheckServices interface {
	HealthcheckServices() (string, error)
}

type IHalthcheckRepo interface {
}
