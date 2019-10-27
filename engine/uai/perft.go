package uai

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gotaxx/libs/ataxx"
)

func Perft(pos ataxx.Position, input string) {
	parts := strings.Split(input, " ")
	if len(parts) < 2 {
		fmt.Println("usage: perft <depth>")
		return
	}

	depth, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	numLeaves := uint64(0)
	for d := 1; d <= depth; d++ {
		startTime := time.Now()
		numLeaves = pos.Perft(d)
		timeTaken := time.Since(startTime)
		fmt.Printf("info depth %d time %s nodes %d\n", d, timeTaken, numLeaves)
	}
	fmt.Println("nodes", numLeaves)
}

func HPerft(pos ataxx.Position, input string) {
	parts := strings.Split(input, " ")
	if len(parts) < 2 {
		fmt.Println("usage: hperft <depth>")
		return
	}

	depth, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	numLeaves := uint64(0)
	for d := 1; d <= depth; d++ {
		startTime := time.Now()
		numLeaves = pos.HashPerft(d)
		timeTaken := time.Since(startTime)
		fmt.Printf("info depth %d time %s nodes %d\n", d, timeTaken, numLeaves)
	}
	fmt.Println("nodes", numLeaves)
}
