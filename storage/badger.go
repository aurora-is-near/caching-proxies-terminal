package storage

import "github.com/dgraph-io/badger/v4"

type BadgerStorage struct {
	bdg *badger.DB
}

func (s *BadgerStorage) Store(bearerToken string, jwt string) error {
	return s.bdg.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(bearerToken), []byte(jwt))
	})
}

func (s *BadgerStorage) Get(bearerToken string) (string, error) {
	var jwt string
	err := s.bdg.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(bearerToken))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			jwt = string(val)
			return nil
		})
	})
	return jwt, err
}

func NewBadgerStorage(path string) (*BadgerStorage, error) {
	bdg, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}

	return &BadgerStorage{bdg: bdg}, nil
}
