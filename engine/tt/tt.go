package tt

import (
	"unsafe"
)

type TT struct {
	Entries []Entry
}

func (tt *TT) Resize(sizeMB int) {
	var ttEntry Entry
	entrySize := int(unsafe.Sizeof(ttEntry))
	numEntries := sizeMB * 1024 * 1024 / entrySize
	tt.Entries = make([]Entry, numEntries)
}

func (tt *TT) Clear() {
	for i := 0; i < len(tt.Entries); i++ {
		tt.Entries[i].data = 0
		tt.Entries[i].key = 0
	}
}

func (tt *TT) Insert(entry Entry) {
	index := entry.key % uint64(len(tt.Entries))
	tt.Entries[index] = entry
}

func (tt *TT) Probe(key uint64) Entry {
	index := key % uint64(len(tt.Entries))
	return tt.Entries[index]
}

var TranspositionTable = TT{
	Entries: make([]Entry, 128*1024*1024/int(unsafe.Sizeof(Entry{}))),
}
