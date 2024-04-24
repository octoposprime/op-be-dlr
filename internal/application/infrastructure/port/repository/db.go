package application

import (
	"context"

	me "github.com/octoposprime/op-be-dlr/internal/domain/model/entity"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// DbPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the database.
type DbPort interface {
	// SetLogger sets logging call-back function
	SetLogger(LogFunc func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error))

	// GetDlrsByFilter returns the dlrs that match the given filter.
	GetDlrsByFilter(ctx context.Context, dlrFilter me.DlrFilter) (me.Dlrs, error)

	// SaveDlr insert a new dlr or update the existing one in the database.
	SaveDlr(ctx context.Context, dlr me.Dlr) (me.Dlr, error)

	// DeleteDlr soft-deletes the given dlr in the database.
	DeleteDlr(ctx context.Context, dlr me.Dlr) (me.Dlr, error)
}
