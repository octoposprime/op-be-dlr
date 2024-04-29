package domain

import (
	me "github.com/octoposprime/op-be-dlr/internal/domain/model/entity"
)

// This is the service layer of the domain layer.
type Service struct {
}

// NewService creates a new *Service.
func NewService() *Service {
	return &Service{}
}

// ValidateDlr validates the dlr.
func (s *Service) ValidateDlr(dlr *me.Dlr) error {
	if err := dlr.Validate(); err != nil {
		return err
	}
	return nil
}
