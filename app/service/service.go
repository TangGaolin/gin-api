package service

import "context"

type Service struct {

}

func New()(s *Service) {
	s = &Service{}
	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {

	return nil
}