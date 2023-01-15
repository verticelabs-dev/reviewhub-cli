package badger

import (
	"fmt"
	"log"
	"os"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/badger/options"
)

// BadgerDB struct for storing the database handle
type BadgerDB struct {
	db *badger.DB
}

func GetStoragePath(appendedPath string) string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s%s", path, appendedPath)
}

// NewBadgerDB creates a new BadgerDB instance
func InitDB() (*BadgerDB, error) {
	path := GetStoragePath("")

	// Open the BadgerDB database
	db, err := badger.Open(badger.Options{
		Dir:                 path,
		ValueDir:            path,
		LevelOneSize:        256 << 20,
		LevelSizeMultiplier: 10,
		TableLoadingMode:    options.MemoryMap,
		ValueLogLoadingMode: options.MemoryMap,
		// table.MemoryMap to mmap() the tables.
		// table.Nothing to not preload the tables.
		MaxLevels:               7,
		MaxTableSize:            64 << 20,
		NumCompactors:           2, // Compactions can be expensive. Only run 2.
		NumLevelZeroTables:      5,
		NumLevelZeroTablesStall: 10,
		NumMemtables:            5,
		SyncWrites:              true,
		NumVersionsToKeep:       1,
		CompactL0OnClose:        true,
		VerifyValueChecksum:     false,
		// Nothing to read/write value log using standard File I/O
		// MemoryMap to mmap() the value log files
		// (2^30 - 1)*2 when mmapping < 2^31 - 1, max int32.
		// -1 so 2*ValueLogFileSize won't overflow on 32-bit systems.
		ValueLogFileSize: 1<<30 - 1,

		ValueLogMaxEntries: 1000000,
		ValueThreshold:     32,
		Truncate:           false,
		Logger:             &ZeroLogAdpater{},
		EventLogging:       true,
		LogRotatesToFlush:  2,
	})

	if err != nil {
		return nil, err
	}

	return &BadgerDB{db: db}, nil
}

// SetString sets a key-value pair in the database
func (b *BadgerDB) SetString(key, value string) error {
	return b.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
}

// GetString gets the value for a key from the database
func (b *BadgerDB) GetString(key string) (string, error) {
	var value []byte
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		value, err = item.Value()
		return err
	})
	return string(value), err
}

// Close closes the underlying database
func (b *BadgerDB) Close() error {
	return b.db.Close()
}
