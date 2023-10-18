package service

type FakeStore struct {
	data map[string]Value
}

func NewFakeStore() *FakeStore {
	return &FakeStore{
		data: make(map[string]Value),
	}
}

func (s *FakeStore) Get(key string) (Value, error) {
	return s.data[key], nil
}
