package application

import (
	"context"

	me "github.com/octoposprime/op-be-dlr/internal/domain/model/entity"
)

// CommandPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type DlrCommandPort interface {
	// CreateDlr sends the given dlr to the application layer for creating a new dlr.
	CreateDlr(ctx context.Context, dlr me.Dlr) (me.Dlr, error)

	// UpdateDlrBase sends the given base values of the dlr to the repository of the infrastructure layer for updating base values of dlr data.
	UpdateDlrBase(ctx context.Context, dlr me.Dlr) (me.Dlr, error)

	// UpdateDlrCore sends the given core values of the dlr to the repository of the infrastructure layer for updating core values of dlr data.
	UpdateDlrCore(ctx context.Context, dlr me.Dlr) (me.Dlr, error)

	// UpdateDlrStatus sends the given status value of the dlr to the repository of the infrastructure layer for updating status of dlr data.
	UpdateDlrStatus(ctx context.Context, dlr me.Dlr) (me.Dlr, error)

	// DeleteDlr sends the given dlr to the application layer for deleting data.
	DeleteDlr(ctx context.Context, dlr me.Dlr) (me.Dlr, error)
}
