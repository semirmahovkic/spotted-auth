package main

import (
	_ "github.com/lib/pq"
	l "github.com/x64puzzle/spotted-common/log"
	"github.com/x64puzzle/spotted-common/storage"
)

func main() {
	l.Log.Info("Spotted Auth Service")

	if err := storage.Init(storage.PGBitMask); err != nil {
		l.Log.Fatal("Storage engine initialization failed: ", err)
	}
}
