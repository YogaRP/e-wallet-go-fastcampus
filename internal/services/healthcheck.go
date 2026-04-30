package services

import "ewallet-fastcampus/internal/interfaces"

type Healthcheck struct {
	HealthcheckRepository interfaces.IHalthcheckRepo
}

func (s *Healthcheck) HealthcheckServices() (string, error) {
	return "Service healthy", nil
}
