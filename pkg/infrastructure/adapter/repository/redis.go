package infrastructure

import (
	"context"

	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tredis "github.com/octoposprime/op-be-shared/tool/redis"
)

type RedisAdapter struct {
	*tredis.RedisClient
	Log func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error)
}

func NewRedisAdapter(redisClient *tredis.RedisClient) RedisAdapter {
	adapter := RedisAdapter{
		redisClient,
		Log,
	}

	return adapter
}

// SetLogger sets logging call-back function
func (a *RedisAdapter) SetLogger(LoggerFunc func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error)) {
	a.Log = LoggerFunc
}
