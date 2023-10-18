package service

type Store interface {
	Get(key string) (Value, bool, error)
	Set(key string, value Value) error
}
