// +build int

package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/x64puzzle/spotted-common/storage"
)

func StorageEngineTest(t *testing.T) {
	storage.Init(storage.PGBitMask)

	assert.NotNil(t, storage.PGClient, "PG client should not be nil")
}

func GRPCServerTest(t *testing.T) {

}
