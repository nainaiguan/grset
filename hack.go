package grset

import "context"

type grset interface {
	Register(id interface{}) context.Context
	Shut(id interface{})
	Total() int
}

func (grs *grSet) Register(id interface{}) context.Context {
	defer grs.total.add()
	return grs.set.regist(id)
}

func (grs *grSet) Shut(id interface{}) {
	defer grs.total.done()
	grs.set.shut(id)
}

func (grs *grSet) Total() int {
	return grs.total.load()
}
