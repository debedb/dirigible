package dirigible

import (
	libkvStore "github.com/docker/libkv/store"
)

type Executor struct {
	store libkvStore.Store
}

func NewExecutor(storeConfig StoreConfig) (*Executor, error) {
	store, err := newStore(storeConfig)
	if err != nil {
		return nil, err
	}
	return &Executor{store: store}, nil
}
