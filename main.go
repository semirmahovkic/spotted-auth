package main

import (
	_ "github.com/lib/pq"
	"github.com/x64integer/go-common/storage"
	"github.com/x64integer/spotted-auth/service"
	l "github.com/x64integer/spotted-common/log"
)

func main() {
	l.Log.Info("Spotted Auth Service")

	if err := storage.Init(storage.PGBitMask | storage.RedisBitMask | storage.CacheBitMask); err != nil {
		l.Log.Fatal("storage engine initialization failed: ", err)
	}

	service.ListenGRPC()
}
