package service

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
	"time"
)

func WriteHandle() {
	for {
		req := <-secKillConf.SecReqChan
		conn := secKillConf.proxy2LayerRedisPool.Get()

		data, err := json.Marshal(req)
		if err != nil {
			logs.Error("json.Marshal failed, error:%v req:%v", err, req)
			conn.Close()
			continue
		}

		_, err = conn.Do("LPUSH", "sec_queue", string(data))
		if err != nil {
			logs.Error("lpush failed, err:%v, req:%v", err, req)
			conn.Close()
			continue
		}
		logs.Debug("LPUSH from redis succ, data:%s", string(data))

		conn.Close()
	}
}

func ReadHandle() {

	for {
		conn := secKillConf.proxy2LayerRedisPool.Get()
		reply, err := conn.Do("RPOP", "recv_queue")
		data, err := redis.String(reply, err)
		if err == redis.ErrNil {
			conn.Close()
			time.Sleep(time.Second)
			continue
		}
		logs.Debug("rpop from redis succ, data:%s", string(data))
		if err != nil {
			logs.Error("rpop failed, err:%v", err)
			conn.Close()
			continue
		}

		var result SecResult
		err = json.Unmarshal([]byte(data), &result)
		if err != nil {
			logs.Error("json.Unmarshal failed, err:%v", err)
			conn.Close()
			continue
		}

		userKey := fmt.Sprintf("%d_%d", result.UserId, result.ProductId)

		secKillConf.UserConnMapLock.Lock()
		resultChan, ok := secKillConf.UserConnMap[userKey]
		secKillConf.UserConnMapLock.Unlock()

		if !ok {
			conn.Close()
			logs.Warn("user not found: %v", userKey)
			continue
		}

		resultChan <- &result
		conn.Close()

	}
}
