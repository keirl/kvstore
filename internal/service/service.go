package service

type Value []byte

type Service struct {
	store Store
}

func New(store Store) *Service {
	return &Service{store: store}
}

type SetVCmd struct {
	Key     string
	Version int
	Value   Value
}

func (s *Service) SetV(cmd SetVCmd) error {
	return nil
}
