package perft_tt

import (
	"unsafe"
)

type PerftTT struct {
	Entries []PerftEntry
}

func (tt *PerftTT) Resize(sizeMB int) {
	var ttPerftEntry PerftEntry
	entrySize := int(unsafe.Sizeof(ttPerftEntry))
	numEntries := sizeMB * 1024 * 1024 / entrySize
	tt.Entries = make([]PerftEntry, numEntries)
}

func (tt *PerftTT) Clear() {
	for i := 0; i < len(tt.Entries); i++ {
		tt.Entries[i].NodeCount = 0
		tt.Entries[i].Depth = 0
		tt.Entries[i].Key = 0
	}
}

func (tt *PerftTT) Insert(entry PerftEntry) {
	index := entry.Key % uint64(len(tt.Entries))
	tt.Entries[index] = entry
}

func (tt *PerftTT) Probe(key uint64) PerftEntry {
	index := key % uint64(len(tt.Entries))
	return tt.Entries[index]
}

var PerftTranspositionTable = PerftTT{
	Entries: make([]PerftEntry, 128*1024*1024/int(unsafe.Sizeof(PerftEntry{}))),
}
