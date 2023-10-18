package service

type Store interface {
	Get(key string) (Value, error)
	Set(key string, value Value) error
}
