package sled

type GetSeter interface {
	Set(key string, v interface{}) error
	Get(key string, v interface{}) error
}

type Iterater interface {
	Iterate(<-chan struct{}) <-chan Element
}

type CRUD interface {
	GetSeter
	SetIfNil(string, interface{}) bool
	Delete(string) (interface{}, bool)
}

type OpenCloser interface {
	Open(*string) error
	Close() error
}

type Snapshoter interface {
	Snapshot() CRUD
}

// KV is a generalized interface to any key/value store
type KV interface {
	CRUD
	Iterater
	OpenCloser
	Snapshoter
}