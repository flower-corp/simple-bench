package benchmark

import (
	"github.com/xujiajun/nutsdb"
	"testing"
)

var nutsDB *nutsdb.DB

func init() {
	opts := nutsdb.DefaultOptions
	opts.Dir = "bench/nutsdb"
	opts.SyncEnable = false
	opts.EntryIdxMode = nutsdb.HintKeyAndRAMIdxMode
	var err error
	nutsDB, err = nutsdb.Open(opts)
	if err != nil {
		panic(err)
	}
}

func initNutsDBData() {
	for i := 0; i < 500000; i++ {
		nutsDB.Update(func(tx *nutsdb.Tx) error {
			err := tx.Put("test-bucket", GetKey(i), GetValue(), 0)
			if err != nil {
				panic(err)
			}
			return nil
		})
	}
}

func BenchmarkPutValue_NutsDB(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		nutsDB.Update(func(tx *nutsdb.Tx) error {
			err := tx.Put("test-bucket", GetKey(i), GetValue(), 0)
			if err != nil {
				panic(err)
			}
			return nil
		})
	}
}

func BenchmarkGetValue_NutsDB(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		nutsDB.View(func(tx *nutsdb.Tx) error {
			_, err := tx.Get("test-bucket", GetKey(i))
			if err != nil {
				panic(err)
			}
			return nil
		})
	}
}
