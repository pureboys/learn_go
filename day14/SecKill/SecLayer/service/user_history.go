package service

import "sync"

type UserBuyHistory struct {
	history map[int]int
	lock    sync.RWMutex
}
