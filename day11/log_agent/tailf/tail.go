package tailf

import (
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"time"
)

type CollectConf struct {
	LogPath string `json:"log_path"`
	Topic   string `json:"topic"`
}

type TailObj struct {
	tail *tail.Tail
	conf CollectConf
}

type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan  chan *TextMsg
}

type TextMsg struct {
	Msg   string
	Topic string
}

var (
	tailObjMgr *TailObjMgr
)

func InitTail(conf []CollectConf, chanSize int) (err error) {

	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize),
	}

	if len(conf) == 0 {
		logs.Error("invalid config for log collect, conf:%v", conf)
		return
	}

	for _, v := range conf {
		obj := &TailObj{
			conf: v,
		}

		tails, errTail := tail.TailFile(v.LogPath, tail.Config{
			ReOpen: true,
			Follow: true,
			//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
			MustExist: false,
			Poll:      true,
		})

		if errTail != nil {
			logs.Error("init tail file failed, conf:%v, err:%v", conf, errTail)
			continue
		}

		obj.tail = tails

		tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj)

		go readFromTail(obj)
	}

	return
}

func readFromTail(tailObj *TailObj) {
	for {
		line, ok := <-tailObj.tail.Lines
		if !ok {
			logs.Warn("tail file close reopen, filename: %s\n", tailObj.tail.Filename)
			time.Sleep(time.Millisecond * 100)
			continue
		}

		textMsg := &TextMsg{
			Msg:   line.Text,
			Topic: tailObj.conf.Topic,
		}

		tailObjMgr.msgChan <- textMsg
	}
}

func GetOneLine() (msg *TextMsg) {
	msg = <-tailObjMgr.msgChan
	return
}

func UpdateConfig(conf []CollectConf) (err error) {
	for _, oneConf := range conf {
		var isRunning = false
		for _, obj := range tailObjMgr.tailObjs {
			if oneConf.LogPath == obj.conf.LogPath {
				isRunning = true
				break
			}
		}

		if isRunning {
			continue
		}

		createNewTask(oneConf)
	}
	return
}

func createNewTask(conf CollectConf) {
	obj := &TailObj{
		conf: conf,
	}

	tails, errTail := tail.TailFile(conf.LogPath, tail.Config{
		ReOpen: true,
		Follow: true,
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if errTail != nil {
		logs.Error("collect filename[%s] failed, err:%v", conf.LogPath, errTail)
		return
	}

	obj.tail = tails

	tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj)

	go readFromTail(obj)
}
