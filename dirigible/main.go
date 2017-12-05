package main

import (
	"flag"
	"os"

	"github.com/debedb/dirigible"
	log "github.com/romana/rlog"
)

func main() {
	endpointsStr := flag.String("endpoints", "localhost:2379", "Comma-separated list of endpoints.")
	storeType := flag.String("type", "etcd", "Storage type to use")
	prefix := flag.String("prefix", dirigible.DefaultPrefix, "Prefix for the KV store")
	flag.Parse()
	if endpointsStr == nil {
		log.Errorf("No etcd endpoints specified")
		os.Exit(1)
	}

	storeConfig := dirigible.NewStoreConfig(*storeType, *endpointsStr, *prefix)
	executor, err := dirigible.NewExecutor(storeConfig)
	if err != nil {
		log.Error(err)
		os.Exit(3)
	}
	dirigible.RunShell(executor)
}
