package storage

import (
	"errors"
	"sync"
)

type MemoryStorage struct {
	*sync.Mutex
	data map[string]string
}

func (s *MemoryStorage) Store(bearerToken string, jwt string) error {
	s.Lock()
	defer s.Unlock()

	s.data[bearerToken] = jwt
	return nil
}

func (s *MemoryStorage) Get(bearerToken string) (string, error) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.data[bearerToken]; !ok {
		return "", errors.New("token not found")
	}
	return s.data[bearerToken], nil
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:  make(map[string]string),
		Mutex: &sync.Mutex{},
	}
}
