package presentation

import (
	"context"

	dto "github.com/octoposprime/op-be-dlr/pkg/presentation/dto"
	pb_dlr "github.com/octoposprime/op-be-shared/pkg/proto/pb/dlr"
)

// GetDlrsByFilter returns the dlrs that match the given filter.
func (a *Grpc) GetDlrsByFilter(ctx context.Context, filter *pb_dlr.DlrFilter) (*pb_dlr.Dlrs, error) {
	dlrs, err := a.queryHandler.GetDlrsByFilter(ctx, *dto.NewDlrFilter(filter).ToEntity())
	return dto.NewDlrFromEntities(dlrs).ToPbs(), err
}

// CreateDlr sends the given dlr to the application layer for creating new dlr.
func (a *Grpc) CreateDlr(ctx context.Context, dlr *pb_dlr.Dlr) (*pb_dlr.Dlr, error) {
	data, err := a.commandHandler.CreateDlr(ctx, *dto.NewDlr(dlr).ToEntity())
	return dto.NewDlrFromEntity(data).ToPb(), err
}

// UpdateDlrBase sends the given dlr to the application layer for updating dlr's base values.
func (a *Grpc) UpdateDlrBase(ctx context.Context, dlr *pb_dlr.Dlr) (*pb_dlr.Dlr, error) {
	data, err := a.commandHandler.UpdateDlrBase(ctx, *dto.NewDlr(dlr).ToEntity())
	return dto.NewDlrFromEntity(data).ToPb(), err
}

// UpdateDlrCore sends the given dlr to the application layer for updating dlr's core values.
func (a *Grpc) UpdateDlrCore(ctx context.Context, dlr *pb_dlr.Dlr) (*pb_dlr.Dlr, error) {
	data, err := a.commandHandler.UpdateDlrCore(ctx, *dto.NewDlr(dlr).ToEntity())
	return dto.NewDlrFromEntity(data).ToPb(), err
}

// UpdateDlrStatus sends the given dlr to the application layer for updating dlr status.
func (a *Grpc) UpdateDlrStatus(ctx context.Context, dlr *pb_dlr.Dlr) (*pb_dlr.Dlr, error) {
	data, err := a.commandHandler.UpdateDlrStatus(ctx, *dto.NewDlr(dlr).ToEntity())
	return dto.NewDlrFromEntity(data).ToPb(), err
}

// DeleteDlr sends the given dlr to the application layer for deleting data.
func (a *Grpc) DeleteDlr(ctx context.Context, dlr *pb_dlr.Dlr) (*pb_dlr.Dlr, error) {
	data, err := a.commandHandler.DeleteDlr(ctx, *dto.NewDlr(dlr).ToEntity())
	return dto.NewDlrFromEntity(data).ToPb(), err
}
