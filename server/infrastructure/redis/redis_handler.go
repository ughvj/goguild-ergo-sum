package redis

import (
	"os"

	"github.com/gomodule/redigo/redis"
)

type RedisHandler struct {
	Conn redis.Conn
}

func NewHandler() (RedisHandler, error) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	conn, err := redis.Dial("tcp", host + ":" + port)
	if err != nil {
		return RedisHandler{}, err
	}

	return RedisHandler{
		Conn: conn,
	}, nil
}

func (h RedisHandler) Set(key, value string) (string, error) {
	return redis.String(h.Conn.Do("SET", key, value))
}

func (h RedisHandler) Get(key string) (string, error) {
	return redis.String(h.Conn.Do("GET", key))
}