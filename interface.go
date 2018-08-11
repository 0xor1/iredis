package iredis

import(
	"github.com/gomodule/redigo/redis"
)

type Pool interface {
	Get() Conn
}

type Conn interface{
	redis.Conn
}