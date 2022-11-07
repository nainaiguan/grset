package grset

import (
	"context"
)

type grset interface {
	Register(id any) (context.Context, error)
	Shut(id any)
	Total() int
}

func (grs *grSet) Register(id any) (context.Context, error) {
	defer grs.total.add()
	c, err := grs.set.regist(id)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (grs *grSet) Shut(id any) {
	defer grs.total.done()
	grs.set.shut(id)
}

func (grs *grSet) Total() int {
	return grs.total.load()
}
