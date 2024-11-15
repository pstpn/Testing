package inmemory

import (
	"fmt"
	"sync"
)

type Storage struct {
	sync.RWMutex
	data map[string]interface{}
}

func NewStorage() *Storage {
	return &Storage{
		RWMutex: sync.RWMutex{},
		data:    make(map[string]interface{}),
	}
}

func (s *Storage) Insert(key string, value interface{}) error {
	s.Lock()
	defer s.Unlock()

	if _, exist := s.data[key]; exist {
		return fmt.Errorf("already exist")
	}

	s.data[key] = value
	return nil
}

func (s *Storage) Get(key string) (interface{}, bool) {
	s.RLock()
	defer s.RUnlock()

	value, ok := s.data[key]

	return value, ok
}

func (s *Storage) Update(key string, newValue interface{}) {
	s.Lock()
	defer s.Unlock()

	s.data[key] = newValue
}

func (s *Storage) Delete(key string) {
	s.Lock()
	defer s.Unlock()

	delete(s.data, key)
}
