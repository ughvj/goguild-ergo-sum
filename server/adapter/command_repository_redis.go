package adapter

import (
	_ "github.com/ughvj/goguild-ergo-sum/domain/repository"
	"github.com/ughvj/goguild-ergo-sum/infrastructure/redis"
)

type CommandRepositoryRedis struct {
	handler redis.RedisHandler
}

func NewCommandRepositoryRedis() (CommandRepositoryRedis, error) {
	handler, err := redis.NewHandler()
	if err != nil {
		return CommandRepositoryRedis{}, err
	}
	return CommandRepositoryRedis{
		handler: handler,
	}, nil
}

func (r CommandRepositoryRedis) Get(key string) (string, error) {
	return r.handler.Get(key)
}