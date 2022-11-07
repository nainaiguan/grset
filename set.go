package grset

import "context"

func (s *set) regist(id interface{}) context.Context {
	s.mu.Lock()
	defer s.mu.Unlock()
	c, cf := context.WithCancel(context.Background())
	s.item[id] = setItem{
		cf: cf,
	}
	return c
}

func (s *set) shut(id interface{}) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.item[id].cf()
}
