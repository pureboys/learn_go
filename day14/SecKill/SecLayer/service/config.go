package service

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/gomodule/redigo/redis"
)

var (
	secLayerContext = &SecLayerContext{}
)

type RedisConf struct {
	RedisAddr        string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout int
}

type EtcdConf struct {
	EtcdAddr          string
	Timeout           int
	EtcdSecKeyPrefix  string
	EtcdSecProductKey string
}

type SecLayerConf struct {
	Proxy2LayerRedis RedisConf
	Layer2ProxyRedis RedisConf
	EtcdConfig       EtcdConf
	LogPath          string
	LogLevel         string

	WriteGoroutineNum int
	ReadGoroutineNum  int
	SecProductInfoMap map[int]*SecProductInfoConf
}

type SecLayerContext struct {
	proxy2LayerRedisPool *redis.Pool
	layer2ProxyRedisPool *redis.Pool
	etcdClient           *clientv3.Client
}

type SecProductInfoConf struct {
	ProductId int
	StartTime int64
	EndTime   int64
	Status    int
	Total     int
	Left      int
}
