package service

import "sync"

const (
	ProductStatusNormal = iota
	ProductStatusSaleOut
	ProductStatusForceSaleOut
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

type SecSkillConf struct {
	RedisConf         RedisConf
	EtcdConf          EtcdConf
	LogPath           string
	LogLevel          string
	SecProductInfoMap map[int]*SecProductInfoConf
	RwSecProductLock  sync.RWMutex
}

type SecProductInfoConf struct {
	ProductId int
	StartTime int64
	EndTime   int64
	Status    int
	Total     int
	Left      int
}
