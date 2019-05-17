package main

import "demo/day9/example4/chat_server/model"

var (
	mgr *model.UserMgr
)

func initUserMgr() {
	mgr = model.NewUserMgr(pool)
}
