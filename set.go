package grset

import (
	"context"
	"errors"
	"log"
)

func (s *set) regist(id any) (context.Context, error) {
	s.mu.RLock()
	if _, ok := s.item[id]; ok {
		err := errors.New("duplicate id")
		log.Printf("regist err: %s", err)
		return nil, err
	}
	s.mu.RUnlock()

	s.mu.Lock()
	defer s.mu.Unlock()
	c, cf := context.WithCancel(context.Background())
	s.item[id] = setItem{
		cf: cf,
	}
	return c, nil
}

func (s *set) shut(id any) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.item[id].cf()
}
