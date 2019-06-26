package service

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

var (
	secKillConf *SecSkillConf
)

func InitService(serviceConf *SecSkillConf) {
	secKillConf = serviceConf
	logs.Debug("init service success, conf: %v", secKillConf)
}

func SecInfo(productId int) (data map[string]interface{}, code int, err error) {

	secKillConf.RwSecProductLock.RLock()
	defer secKillConf.RwSecProductLock.RUnlock()

	v, ok := secKillConf.SecProductInfoMap[productId]
	if !ok {
		code = ErrNotFoundProductId
		err = fmt.Errorf("not found product_id:%d", productId)
		return
	}

	data = make(map[string]interface{})
	data["product_id"] = productId
	data["start_time"] = v.StartTime
	data["end_time"] = v.EndTime
	data["status"] = v.Status

	return
}
