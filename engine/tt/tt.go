package tt

import (
	"fmt"
	"unsafe"
)

type TT struct {
	Entries []TTEntry
}

func (tt *TT) Resize(sizeMB int) {
	fmt.Println("DesiredSize", sizeMB)
	var ttEntry TTEntry
	entrySize := int(unsafe.Sizeof(ttEntry))
	fmt.Println("EntrySize", entrySize)
	numEntries := sizeMB * 1024 * 1024 / entrySize
	fmt.Println("NumEntries", numEntries)
	tt.Entries = make([]TTEntry, numEntries)
	fmt.Println(len(tt.Entries))
}

func (tt *TT) Clear() {
	for i := 0; i < len(tt.Entries); i++ {
		tt.Entries[i].data = 0
		tt.Entries[i].key = 0
	}
}

func (tt *TT) Insert(entry TTEntry) {
	index := entry.key % uint64(len(tt.Entries))
	tt.Entries[index] = entry
}

func (tt *TT) Probe(key uint64) TTEntry {
	index := key & uint64(len(tt.Entries))
	return tt.Entries[index]
}

var TranspositionTable = TT{
	Entries: make([]TTEntry, 128),
}
