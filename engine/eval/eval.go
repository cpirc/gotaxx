package eval

import "github.com/cpirc/gotaxx/libs/ataxx"

func Eval(pos *ataxx.Position) int {
	score := pos.Us().Count() - pos.Them().Count()
	return score
}
