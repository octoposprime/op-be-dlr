package application

import (
	"context"

	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// RedisPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the redis.
type RedisPort interface {
	// SetLogger sets logging call-back function
	SetLogger(LogFunc func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error))
}
