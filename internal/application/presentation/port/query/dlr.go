package application

import (
	"context"

	me "github.com/octoposprime/op-be-dlr/internal/domain/model/entity"
)

// QueryPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type DlrQueryPort interface {
	// GetDlrsByFilter returns the dlrs that match the given filter.
	GetDlrsByFilter(ctx context.Context, dlrFilter me.DlrFilter) (me.Dlrs, error)
}
