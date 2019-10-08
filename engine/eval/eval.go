package eval

import "github.com/cpirc/gotaxx/libs/ataxx"

func Eval(pos *ataxx.Position) int {
	return pos.Us().Count() - pos.Them().Count()
}
