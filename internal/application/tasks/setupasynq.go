package tasks

import (
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/marveldo/gogin/internal/config"
)

var redis_opt *asynq.RedisClientOpt

func GetRedisOptions(config *config.Config) *asynq.RedisClientOpt {
	if redis_opt == nil {
	   redis_opt = &asynq.RedisClientOpt{
		Addr:  fmt.Sprintf("%v:%v", config.Redis_Host , config.Redis_Port),
		DB: 0,
	   }
    }
	return redis_opt
}

func CreateAsynqClient(r *asynq.RedisClientOpt) *asynq.Client {

	client := asynq.NewClient(r)

	return client
}

func GetAsynqServer(r *asynq.RedisClientOpt) *asynq.Server {
	return asynq.NewServer(r, asynq.Config{
		Concurrency: 10,
	} )
}


