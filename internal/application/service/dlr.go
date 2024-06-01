package application

import (
	"context"

	"github.com/google/uuid"
	me "github.com/octoposprime/op-be-dlr/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-dlr/internal/domain/model/object"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// GetDlrsByFilter returns the dlrs that match the given filter.
func (a *Service) GetDlrsByFilter(ctx context.Context, dlrFilter me.DlrFilter) (me.Dlrs, error) {
	return a.DbPort.GetDlrsByFilter(ctx, dlrFilter)
}

// CreateDlr sends the given dlr to the repository of the infrastructure layer for creating a new dlr.
func (a *Service) CreateDlr(ctx context.Context, dlr me.Dlr) (me.Dlr, error) {
	dlr.Id = uuid.UUID{}
	if err := a.ValidateDlr(&dlr); err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "CreateUser", userId, err.Error()))
		return me.Dlr{}, err
	}
	if dlr.DlrStatus == mo.DlrStatusNONE {
		dlr.DlrStatus = mo.DlrStatusACTIVE
	}
	return a.DbPort.SaveDlr(ctx, dlr)
}

// UpdateDlrBase sends the given base values of the dlr to the repository of the infrastructure layer for updating base values of dlr data.
func (a *Service) UpdateDlrBase(ctx context.Context, dlr me.Dlr) (me.Dlr, error) {
	if user.Id.String() == "" || user.Id == (uuid.UUID{}) {
		err := mo.ErrorUserIdIsEmpty
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return me.User{}, err
	}
	var userFilter me.UserFilter
	userFilter.Id = user.Id
	users, err := a.GetUsersByFilter(ctx, userFilter)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return me.User{}, err
	}
	if users.TotalRows > 0 {
		dbUser := users.Users[0]
		dbUser.Tags = user.Tags
		dbUser.DlrType = user.DlrType
		if err := a.ValidateUser(&dbUser); err != nil {
			userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
			go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
			return me.User{}, err
		}
		return a.DbPort.SaveUser(ctx, dbUser)
	} else {
		return user, mo.ErrorUserNotFound
	}
}

// UpdateDlrCore sends the given core values of the dlr to the repository of the infrastructure layer for updating core values of dlr data.
func (a *Service) UpdateDlrCore(ctx context.Context, dlr me.Dlr) (me.Dlr, error) {
	if user.Id.String() == "" || user.Id == (uuid.UUID{}) {
		err := mo.ErrorUserIdIsEmpty
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return me.User{}, err
	}
	var userFilter me.UserFilter
	userFilter.Id = user.Id
	users, err := a.GetUsersByFilter(ctx, userFilter)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return me.User{}, err
	}
	if users.TotalRows > 0 {
		dbUser := users.Users[0]
		dbUser.DlrData = user.DlrData
		if err := a.ValidateUser(&dbUser); err != nil {
			userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
			go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
			return me.User{}, err
		}
		return a.DbPort.SaveUser(ctx, dbUser)
	} else {
		return user, mo.ErrorUserNotFound
	}
}

// UpdateDlrStatus sends the given status value of the DLR entry to the repository of the infrastructure layer for updating the status of DLR data.
func (a *Service) UpdateDlrStatus(ctx context.Context, dlr me.Dlr) (me.Dlr, error) {
	if dlr.Id.String() == "" || dlr.Id == (uuid.UUID{}) {
		err := mo.ErrorDlrIdIsEmpty
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateDlrStatus", dlrId, err.Error()))
		return me.Dlr{}, err
	}

	var dlrFilter me.DlrFilter
	dlrFilter.Id = dlr.Id
	dlrs, err := a.GetDlrsByFilter(ctx, dlrFilter)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateDlrStatus", dlrId, err.Error()))
		return me.Dlr{}, err
	}

	if dlrs.TotalRows > 0 {
		dbDlr := dlrs.Dlrs[0]
		dbDlr.DlrStatus = dlr.DlrStatus
		if err := a.ValidateDlr(&dbDlr); err != nil {
			userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
			go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateDlrStatus", dlrId, err.Error()))
			return me.Dlr{}, err
		}
		return a.DbPort.SaveDlr(ctx, dbDlr)
	} else {
		return dlr, mo.ErrorDlrNotFound
	}
}

// DeleteDlr sends the given dlr to the repository of the infrastructure layer for deleting data.
func (a *Service) DeleteDlr(ctx context.Context, dlr me.Dlr) (me.Dlr, error) {
	var err error
	dlr, err = a.DbPort.DeleteDlr(ctx, dlr)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "DeleteDlr", dlrId, err.Error()))
		return me.Dlr{}, err
	}

	return dlr, err
}
