package application

import (
	"context"

	me "github.com/octoposprime/op-be-dlr/internal/domain/model/entity"
)

// GetDlrsByFilter returns the dlrs that match the given filter.
func (a QueryAdapter) GetDlrsByFilter(ctx context.Context, dlrFilter me.DlrFilter) (me.Dlrs, error) {
	return a.Service.GetDlrsByFilter(ctx, dlrFilter)
}
