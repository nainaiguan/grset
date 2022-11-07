package grset

import "context"

type grset interface {
	Register(id any) context.Context
	Shut(id any)
	Total() int
}

func (grs *grSet) Register(id any) context.Context {
	defer grs.total.add()
	return grs.set.regist(id)
}

func (grs *grSet) Shut(id any) {
	defer grs.total.done()
	grs.set.shut(id)
}

func (grs *grSet) Total() int {
	return grs.total.load()
}
