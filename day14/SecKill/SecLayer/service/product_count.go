package service

import "sync"

type ProductCountMgr struct {
	productCount map[int]int
	lock         sync.RWMutex
}
