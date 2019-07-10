package model

import (
	etcdclient "github.com/coreos/etcd/clientv3"
	"github.com/jmoiron/sqlx"
)

var (
	Db             *sqlx.DB
	EtcdClient     *etcdclient.Client
	EtcdPrefix     string
	EtcdProductKey string
)

func Init(db *sqlx.DB, etcClient *etcdclient.Client, prefix, productKey string) (err error) {
	Db = db
	EtcdClient = etcClient
	EtcdPrefix = prefix
	EtcdProductKey = productKey
	return
}
