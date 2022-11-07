package grset

import (
	"context"
	"sync"
)

type grSet struct {
	total total
	set   set
}

type set struct {
	mu   sync.RWMutex
	item map[interface{}]setItem
}

type setItem struct {
	cf context.CancelFunc
}

type total struct {
	mu sync.RWMutex
	n  int
}
