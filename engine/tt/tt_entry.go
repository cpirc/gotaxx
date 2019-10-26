package tt

import (
	"gotaxx/libs/ataxx"
)

type Flag int8

const (
	FLAG_LOWER = 0
	FLAG_UPPER = 1
	FLAG_EXACT = 2
)

type TTEntry struct {
	data uint64
	key  uint64
}

func (tt TTEntry) Move() ataxx.Move {
	data := tt.data ^ tt.key
	return ataxx.Move{
		From: ataxx.Square{Data: uint8(data & 0x00ff)},
		To:   ataxx.Square{Data: uint8((data & 0xff00) >> 8)},
	}
}

func (tt TTEntry) Flag() Flag {
	return Flag(((tt.data ^ tt.key) & 0x30000) >> 16)
}

func (tt TTEntry) Depth() uint8 {
	return uint8(((tt.data ^ tt.key) & 0x1fc0000) >> 18)
}

func (tt TTEntry) Score() int {
	return int((tt.data ^ tt.key) >> 25)
}

func (tt TTEntry) Key() uint64 {
	return tt.key
}

func (tt *TTEntry) Set(move ataxx.Move, flag Flag, depth uint8, score int, key uint64) {
	tt.key = key
	tt.data = uint64(move.From.Data) | (uint64(move.To.Data) << 8) | (uint64(flag) << 16) | (uint64(depth) << 18) | (uint64(score) << 25)
	tt.data ^= key
}
