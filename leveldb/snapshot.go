package leveldb

import (
	"github.com/jmhodges/levigo"
)

type Snapshot struct {
	db           *DB
	s            *levigo.Snapshot
	readOpts     *levigo.ReadOptions
	iteratorOpts *levigo.ReadOptions
}

func newSnapshot(db *DB) *Snapshot {
	s := new(Snapshot)
	s.db = db
	s.s = db.db.NewSnapshot()

	s.readOpts = levigo.NewReadOptions()
	s.readOpts.SetSnapshot(s.s)

	s.iteratorOpts = levigo.NewReadOptions()
	s.iteratorOpts.SetSnapshot(s.s)
	s.iteratorOpts.SetFillCache(false)

	return s
}

func (s *Snapshot) Close() {
	s.db.db.ReleaseSnapshot(s.s)

	s.iteratorOpts.Close()
	s.readOpts.Close()
}

func (s *Snapshot) Get(key []byte) ([]byte, error) {
	return s.db.db.Get(s.readOpts, key)
}

//same as db iterator and reverse iterator
func (s *Snapshot) Iterator(r *Range, limit int) *Iterator {
	return newIterator(s.db, s.iteratorOpts, r, limit, forward)
}

func (s *Snapshot) ReverseIterator(r *Range, limit int) *Iterator {
	return newIterator(s.db, s.iteratorOpts, r, limit, backward)
}

func (s *Snapshot) GetInt(key []byte) (int64, error) {
	return Int(s.Get(key))
}

func (s *Snapshot) GetUInt(key []byte) (uint64, error) {
	return Uint(s.Get(key))
}

func (s *Snapshot) GetFloat(key []byte) (float64, error) {
	return Float(s.Get(key))
}

func (s *Snapshot) GetString(key []byte) (string, error) {
	return String(s.Get(key))
}

func (s *Snapshot) GetSlice(key []byte) ([]byte, error) {
	return Slice(s.Get(key))
}
