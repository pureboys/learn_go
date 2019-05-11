package balance

type Balance interface {
	DoBalance(insts []*Instance, key ...string) (inst *Instance, err error)
}
