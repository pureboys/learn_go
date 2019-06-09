package tailf

import (
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"sync"
	"time"
)

const (
	StatusNormal = 1
	StatusDelete = 2
)

type CollectConf struct {
	LogPath string `json:"log_path"`
	Topic   string `json:"topic"`
}

type TailObj struct {
	tail     *tail.Tail
	conf     CollectConf
	status   int
	exitChan chan int
}

type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan  chan *TextMsg
	lock     sync.Mutex
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
		createNewTask(v)
	}

	return
}

func readFromTail(tailObj *TailObj) {
	for {
		select {
		case line, ok := <-tailObj.tail.Lines:
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

		case <-tailObj.exitChan:
			logs.Warn("tail obj will exited, conf: %v", tailObj.conf)
			return
		}

	}
}

func GetOneLine() (msg *TextMsg) {
	msg = <-tailObjMgr.msgChan
	return
}

func UpdateConfig(conf []CollectConf) (err error) {
	// 并发加锁
	tailObjMgr.lock.Lock()
	defer tailObjMgr.lock.Unlock()

	// 新增部分
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

	//删除部分
	var tailObjs []*TailObj
	for _, obj := range tailObjMgr.tailObjs {
		obj.status = StatusDelete

		for _, oneConf := range conf {
			if oneConf.LogPath == obj.conf.LogPath {
				obj.status = StatusNormal
				break
			}
		}

		if StatusDelete == obj.status {
			obj.exitChan <- 1
			continue
		}

		tailObjs = append(tailObjs, obj)
	}

	tailObjMgr.tailObjs = tailObjs
	return
}

func createNewTask(conf CollectConf) {
	obj := &TailObj{
		conf:     conf,
		exitChan: make(chan int, 100),
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
