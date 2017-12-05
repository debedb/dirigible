package dirigible

import (
	"errors"
	"strings"

	"github.com/docker/libkv"
	libkvStore "github.com/docker/libkv/store"
	libkvEtcd "github.com/docker/libkv/store/etcd"
)

type StoreConfig struct {
	Prefix    string
	Endpoints []string
	Type      string
}

type Store struct {
	config StoreConfig
	libkvStore.Store
}

func NewStoreConfig(storeType string, endpoints string, prefix string) StoreConfig {
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	return StoreConfig{Prefix: prefix,
		Endpoints: strings.Split(endpoints, ","),
		Type:      storeType,
	}
}

func newStore(conf StoreConfig) (*Store, error) {
	var err error
	if conf.Type != "etcd" {
		return nil, errors.New("Only etcd is supported.")
	}

	store := &Store{config: conf}
	store.Store, err = libkv.NewStore(libkvStore.ETCD, conf.Endpoints, nil)
	if err != nil {
		return nil, err
	}

	// Test connection
	_, err = store.Exists(conf.Prefix + "test")
	if err != nil {
		return nil, err
	}

	return store, nil
}

func init() {
	// Register etcd store to libkv
	libkvEtcd.Register()
}
