package service

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"sync"
)

type SecLimitMgr struct {
	UserLimitMap map[int]*Limit
	IPLimitMap   map[string]*Limit
	lock         sync.Mutex
}

func antiSpam(req *SecRequest) (err error) {
	_, ok := secKillConf.idBlackMap[req.UserId]
	if ok {
		err = fmt.Errorf("invalid request")
		logs.Error("useId[%v] is block by id black", req.UserId)
		return
	}

	_, ok = secKillConf.ipBlackMap[req.ClientAddr]
	if ok {
		err = fmt.Errorf("invalid request")
		logs.Error("useId[%v] ip[%v] is block by ip black", req.UserId, req.ClientAddr)
		return
	}

	secKillConf.secLimitMgr.lock.Lock()

	// 用户速度限制
	limit, ok := secKillConf.secLimitMgr.UserLimitMap[req.UserId]
	if !ok {
		limit = &Limit{
			secLimit: &SecLimit{},
			minLimit: &MinLimit{},
		}
		secKillConf.secLimitMgr.UserLimitMap[req.UserId] = limit
	}

	secIdCount := limit.secLimit.Count(req.AccessTime.Unix())
	minIdCount := limit.minLimit.Check(req.AccessTime.Unix())

	// ip速度限制
	limit, ok = secKillConf.secLimitMgr.IPLimitMap[req.ClientAddr]
	if !ok {
		limit = &Limit{
			secLimit: &SecLimit{},
			minLimit: &MinLimit{},
		}
		secKillConf.secLimitMgr.IPLimitMap[req.ClientAddr] = limit
	}
	secIpCount := limit.secLimit.Count(req.AccessTime.Unix())
	minIpCount := limit.minLimit.Count(req.AccessTime.Unix())

	secKillConf.secLimitMgr.lock.Unlock()

	if secIpCount > secKillConf.AccessLimitConf.IPSecAccessLimit {
		err = fmt.Errorf("UserSecAccessLimit invalid requeset")
		return
	}

	if minIpCount > secKillConf.AccessLimitConf.IPMinAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}

	if secIdCount > secKillConf.AccessLimitConf.UserSecAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}

	if minIdCount > secKillConf.AccessLimitConf.UserMinAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}

	return
}
