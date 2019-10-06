package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/cpirc/gotaxx/libs/ataxx"
)

func main() {
	position := ataxx.NewPosition("x5o/7/7/7/7/7/o5x x 0")

	printer := message.NewPrinter(language.English)
	printer.Println("Depth: Leaves")
	for depth := 1; depth <= 6; depth++ {
		numLeaves := position.Perft(depth)
		printer.Printf("%d: %d\n", depth, numLeaves)
	}
}
