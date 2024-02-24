package service

type Value []byte

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

type SetVCmd struct {
	Key     string
	Version int
	Value   Value
}

func (s *Service) SetV(cmd SetVCmd) error {
	err := s.store.Set(cmd.Key, cmd.Value)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Set(key string, value Value) error {
	err := s.store.Set(key, value)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Get(key string) (Value, error) {
	v, err := s.store.Get(key)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (s *Service) Del(key string) error {
	// err := s.store.Del(key)
	// if err != nil {
	// 	return err
	// }
	return nil
}
