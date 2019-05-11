package thrid

import (
	"demo/day6/homework1/balance"
	"fmt"
	"hash/crc32"
	"math/rand"
)

type HashBalance struct {
}

func init() {
	balance.RegisterBalance("hash", &HashBalance{})
}

func (p *HashBalance) DoBalance(insts []*balance.Instance, key ...string) (inst *balance.Instance, err error) {

	var defKey = fmt.Sprintf("%d", rand.Intn(100))

	if len(key) > 0 {
		defKey = key[0]
		return
	}

	lens := len(insts)
	if lens == 0 {
		err = fmt.Errorf("no backend instance")
		return
	}

	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	index := int(hashVal) % lens

	inst = insts[index]
	return
}
