package application

import (
	"context"

	me "github.com/octoposprime/op-be-dlr/internal/domain/model/entity"
)

// CreateDlr sends the given dlr to the application layer for creating a new dlr.
func (a CommandAdapter) CreateDlr(ctx context.Context, dlr me.Dlr) (me.Dlr, error) {
	return a.Service.CreateDlr(ctx, dlr)
}

// UpdateDlrBase sends the given base values of the dlr to the repository of the infrastructure layer for updating base values of dlr data.
func (a CommandAdapter) UpdateDlrBase(ctx context.Context, dlr me.Dlr) (me.Dlr, error) {
	return a.Service.UpdateDlrBase(ctx, dlr)
}

// UpdateDlrCore sends the given core values of the dlr to the repository of the infrastructure layer for updating core values of dlr data.
func (a CommandAdapter) UpdateDlrCore(ctx context.Context, dlr me.Dlr) (me.Dlr, error) {
	return a.Service.UpdateDlrCore(ctx, dlr)
}

// UpdateDlrStatus sends the given status value of the dlr to the repository of the infrastructure layer for updating status of dlr data.
func (a CommandAdapter) UpdateDlrStatus(ctx context.Context, dlr me.Dlr) (me.Dlr, error) {
	return a.Service.UpdateDlrStatus(ctx, dlr)
}

// DeleteDlr sends the given dlr to the application layer for deleting data.
func (a CommandAdapter) DeleteDlr(ctx context.Context, dlr me.Dlr) (me.Dlr, error) {
	return a.Service.DeleteDlr(ctx, dlr)
}
