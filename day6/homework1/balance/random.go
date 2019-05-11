package balance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
}

func init() {
	RegisterBalance("random", &RandomBalance{})
}

func (p *RandomBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("no instance")
		return
	}

	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]

	return
}
