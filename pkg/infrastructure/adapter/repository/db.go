package infrastructure

import (
	"context"

	"github.com/google/uuid"
	me "github.com/octoposprime/op-be-dlr/internal/domain/model/entity"
	map_repo "github.com/octoposprime/op-be-dlr/pkg/infrastructure/mapper/repository"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tgorm "github.com/octoposprime/op-be-shared/tool/gorm"
)

type DbAdapter struct {
	*tgorm.GormClient
	Log func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error)
}

func NewDbAdapter(dbClient *tgorm.GormClient) DbAdapter {
	adapter := DbAdapter{
		dbClient,
		Log,
	}

	err := dbClient.DbClient.AutoMigrate(&map_repo.Dlr{})
	if err != nil {
		panic(err)
	}

	return adapter
}

// SetLogger sets logging call-back function
func (a *DbAdapter) SetLogger(LoggerFunc func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error)) {
	a.Log = LoggerFunc
}

// GetDlrsByFilter returns the dlrs that match the given filter.
func (a DbAdapter) GetDlrsByFilter(ctx context.Context, dlrFilter me.DlrFilter) (me.Dlrs, error) {
	var dlrsDbMapper map_repo.Dlrs
	var filter map_repo.Dlr
	qry := a.DbClient
	if dlrFilter.Id.String() != "" && dlrFilter.Id != (uuid.UUID{}) {
		filter.ID = dlrFilter.Id
	}
	if dlrFilter.DlrData != "" {
		filter.DlrData = dlrFilter.DlrData
	}
	if dlrFilter.DlrType != 0 {
		filter.DlrType = int(dlrFilter.DlrType)
	}
	if dlrFilter.DlrStatus != 0 {
		filter.DlrStatus = int(dlrFilter.DlrStatus)
	}
	if len(dlrFilter.Tags) > 0 {
		filter.Tags = dlrFilter.Tags
	}
	if !dlrFilter.CreatedAtFrom.IsZero() && !dlrFilter.CreatedAtTo.IsZero() {
		qry = qry.Where("created_at between ? and ?", dlrFilter.CreatedAtFrom, dlrFilter.CreatedAtTo)
	}
	if !dlrFilter.UpdatedAtFrom.IsZero() && !dlrFilter.UpdatedAtTo.IsZero() {
		qry = qry.Where("updated_at between ? and ?", dlrFilter.UpdatedAtFrom, dlrFilter.UpdatedAtTo)
	}
	if dlrFilter.SearchText != "" {
		qry = qry.Where(
			qry.Where("UPPER(dlr_name) LIKE UPPER(?)", "%"+dlrFilter.SearchText+"%").
				Or("UPPER(email) LIKE UPPER(?)", "%"+dlrFilter.SearchText+"%").
				Or("UPPER(array_to_string(tags, ',')) LIKE UPPER(?)", "%"+dlrFilter.SearchText+"%"),
		)
	}
	qry = qry.Where(filter)
	var totalRows int64
	result := qry.Model(&map_repo.Dlr{}).Where(filter).Count(&totalRows)
	if result.Error != nil {
		dlrId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetDlrsByFilter", dlrId, result.Error.Error()))
		totalRows = 0
	}
	if dlrFilter.Limit != 0 {
		qry = qry.Limit(dlrFilter.Limit)
	}
	if dlrFilter.Offset != 0 {
		qry = qry.Offset(dlrFilter.Offset)
	}
	if dlrFilter.SortType != "" && dlrFilter.SortField != 0 {
		sortStr := map_repo.DlrSortMap[dlrFilter.SortField]
		if dlrFilter.SortType == "desc" || dlrFilter.SortType == "DESC" {
			sortStr += " desc"
		} else {
			sortStr += " asc"
		}
		qry = qry.Order(sortStr)
	}
	result = qry.Find(&dlrsDbMapper)
	if result.Error != nil {
		dlrId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetDlrsByFilter", dlrId, result.Error.Error()))
		return me.Dlrs{}, result.Error
	}
	return me.Dlrs{
		Dlrs:      dlrsDbMapper.ToEntities(),
		TotalRows: totalRows,
	}, nil
}

// SaveDlr insert a new dlr or update the existing one in the database.
func (a DbAdapter) SaveDlr(ctx context.Context, dlr me.Dlr) (me.Dlr, error) {
	dlrDbMapper := map_repo.NewDlrFromEntity(dlr)
	qry := a.DbClient
	if dlr.Id.String() != "" && dlr.Id != (uuid.UUID{}) {
		qry = qry.Omit("created_at")
	}
	dlrId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	if dlrDbMapper.ID != (uuid.UUID{}) {
		dlrDbMapper.UpdatedBy, _ = uuid.Parse(dlrId)
	} else {
		dlrDbMapper.CreatedBy, _ = uuid.Parse(dlrId)
	}
	result := qry.Save(&dlrDbMapper)
	if result.Error != nil {
		dlrId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "SaveDlr", dlrId, result.Error.Error()))
		return me.Dlr{}, result.Error
	}
	return *dlrDbMapper.ToEntity(), nil
}

// DeleteDlr soft-deletes the given dlr in the database.
func (a DbAdapter) DeleteDlr(ctx context.Context, dlr me.Dlr) (me.Dlr, error) {
	dlrDbMapper := map_repo.NewDlrFromEntity(dlr)
	dlrId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	dlrDbMapper.DeletedBy, _ = uuid.Parse(dlrId)
	result := a.DbClient.Delete(&dlrDbMapper)
	if result.Error != nil {
		dlrId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "DeleteDlr", dlrId, result.Error.Error()))
		return me.Dlr{}, result.Error
	}
	return *dlrDbMapper.ToEntity(), nil
}
