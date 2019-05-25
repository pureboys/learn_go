package proto

import (
	"demo/day9/chat/common"
)

type Message struct {
	Cmd  string `json:"cmd"`
	Data string `json:"data"`
}

type LoginCmd struct {
	Id     int    `json:"user_id"`
	Passwd string `json:"passwd"`
}

type RegisterCmd struct {
	User common.User `json:"user"`
}

type LoginCmdRes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	User  []int  `json:"users"`
}
