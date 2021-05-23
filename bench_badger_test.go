package benchmark

import (
	"github.com/dgraph-io/badger/v3"
	"log"
	"testing"
)

var (
	badgerDB *badger.DB
	err      error
)

func init() {
	dir := "bench/badgerDB"
	badgerDB, err = badger.Open(badger.DefaultOptions(dir))
	if err != nil {
		log.Fatal("open badger err.", err)
	}
}

func initBadgerData() {
	for i := 0; i < 10000; i++ {
		key := GetKey(i)
		val := GetValue()

		if err = badgerDB.Update(
			func(txn *badger.Txn) error {
				return txn.Set(key, val)
			}); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkPutValue_BadgerDB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := GetKey(n)
		val := GetValue()

		if err = badgerDB.Update(
			func(txn *badger.Txn) error {
				return txn.Set(key, val)
			}); err != nil {
			b.Fatal("badger write data err.", err)
		}
	}
}

func BenchmarkGetValue_BadgerDB(b *testing.B) {
	initBadgerData()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := badgerDB.View(func(tx *badger.Txn) error {
			key := GetKey(i)
			_, _ = tx.Get(key)
			return nil
		})
		if err != nil {
			log.Fatal("badger get data err.", err)
		}
	}
}