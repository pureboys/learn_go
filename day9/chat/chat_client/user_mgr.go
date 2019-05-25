package main

import (
	"demo/day9/chat/common"
	"demo/day9/chat/proto"
	"fmt"
)

var onlineUserMap = make(map[int]*common.User, 16)

func outPutUserOnline() {
	fmt.Println("Online User List:")
	for id := range onlineUserMap {
		if id == userId {
			continue
		}
		fmt.Println("user:", id)
	}
}

func updateUserStatus(userStatus proto.UserStatusNotify) {
	user, ok := onlineUserMap[userStatus.UserId]
	if !ok {
		user = &common.User{}
		user.UserId = userStatus.UserId
	}

	user.Status = userStatus.Status
	onlineUserMap[user.UserId] = user

	outPutUserOnline()
}
