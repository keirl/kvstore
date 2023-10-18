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

func (s *FakeStore) Set(key string, value Value) error {
	s.data[key] = value
	return nil
}
