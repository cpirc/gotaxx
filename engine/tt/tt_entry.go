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

type Entry struct {
	data uint64
	key  uint64
}

func (tt Entry) Move() ataxx.Move {
	data := tt.data ^ tt.key
	return ataxx.Move{
		From: ataxx.Square{Data: uint8(data & 0x00ff)},
		To:   ataxx.Square{Data: uint8((data & 0xff00) >> 8)},
	}
}

func (tt Entry) Flag() Flag {
	return Flag(((tt.data ^ tt.key) & 0x30000) >> 16)
}

func (tt Entry) Depth() int {
	return int(((tt.data ^ tt.key) & 0x1fc0000) >> 18)
}

func (tt Entry) Score() int {
	return int((tt.data ^ tt.key) >> 25)
}

func (tt Entry) Key() uint64 {
	return tt.key
}

func (tt *Entry) Set(move ataxx.Move, flag Flag, depth uint8, score int, key uint64) {
	tt.key = key
	tt.data = uint64(move.From.Data) | (uint64(move.To.Data) << 8) | (uint64(flag) << 16) | (uint64(depth) << 18) | (uint64(score) << 25)
	tt.data ^= key
}
