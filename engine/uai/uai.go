package uai

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cpirc/gotaxx/engine/options"
	"github.com/cpirc/gotaxx/libs/ataxx"
)

func PrintIdentity() {
	fmt.Println("id name Gotaxx")
	fmt.Println("id author kz04px mkchan")
}

func PrintOptions() {
	for _, option := range options.ComboOptions {
		fmt.Print("option name ", option.Name(), " type ", option.Type(), " default ", option.Default(), " ")
		for _, choice := range option.Choices() {
			fmt.Print("var ", choice, " ")
		}
		fmt.Println()
	}
}

func Loop() {
	PrintIdentity()
	fmt.Println()
	PrintOptions()
	fmt.Println()
	fmt.Println("uaiok")

	pos, _ := ataxx.NewPosition(STARTPOS)
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		input = input[:len(input) - 1]

		if input == "quit" {
			break
		} else if input == "isready" {
			IsReady()
		} else if strings.HasPrefix(input, "position") {
			pos, err = Position(input)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else if strings.HasPrefix(input, "go") {
			Go(*pos, input)
		} else if strings.HasPrefix(input, "setoption") {
			SetOption(input)
		} else if input == "options" {
			PrintOptions()
		} else if strings.HasPrefix(input, "perft") {
			Perft(*pos, input)
		} else if input == "d" || input == "print" {
			pos.Print()
		} else {
			fmt.Println("Unrecognized command!")
		}
	}
}
