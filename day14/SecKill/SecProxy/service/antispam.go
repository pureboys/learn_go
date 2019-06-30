package service

import (
	"fmt"
	"sync"
)

var (
	secLimitMgr = &SecLimitMgr{
		UserLimitMap: make(map[int]*SecLimit, 10000),
	}
)

type SecLimitMgr struct {
	UserLimitMap map[int]*SecLimit
	IPLimitMap   map[string]*SecLimit
	lock         sync.Mutex
}

type SecLimit struct {
	count   int
	curTime int64
}

func antiSpam(req *SecRequest) (err error) {
	secLimitMgr.lock.Lock()

	// 用户速度限制
	secLimit, ok := secLimitMgr.UserLimitMap[req.UserId]
	if !ok {
		secLimit = &SecLimit{}
		secLimitMgr.UserLimitMap[req.UserId] = secLimit
	}
	count := secLimit.Count(req.AccessTime.Unix())

	// ip速度限制
	ipLimit, ok := secLimitMgr.IPLimitMap[req.ClientAddr]
	if !ok {
		ipLimit = &SecLimit{}
		secLimitMgr.IPLimitMap[req.ClientAddr] = ipLimit
	}
	count2 := secLimit.Count(req.AccessTime.Unix())

	secLimitMgr.lock.Unlock()

	// 用户超过限制则进行限流
	if count > secKillConf.UserSecAccessLimit {
		err = fmt.Errorf("UserSecAccessLimit invalid requeset")
		return
	}

	// IP超过限制则进行限流
	if count2 > secKillConf.IPSecAccessLimit {
		err = fmt.Errorf("IPSecAccessLimit invalid requeset")
		return
	}

	return
}

func (p *SecLimit) Count(nowTime int64) (curCount int) {
	if p.curTime != nowTime {
		p.count = 1
		p.curTime = nowTime
		curCount = p.count
		return
	}

	p.count++
	curCount = p.count
	return
}

func (p *SecLimit) Check(nowTime int64) int {
	if p.curTime != nowTime {
		return 0
	}
	return p.count
}
