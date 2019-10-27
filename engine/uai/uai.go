package uai

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"gotaxx/engine/options"
	"gotaxx/libs/ataxx"
)

func PrintIdentity() {
	fmt.Println("id name Gotaxx")
	fmt.Println("id author kz04px mkchan")
}

func PrintOptions() {
	for _, option := range options.ComboOptions {
		fmt.Print("option name ", option.Name, " type combo default ", option.DefaultValue, " ")
		for _, choice := range option.Choices {
			fmt.Print("var ", choice, " ")
		}
		fmt.Println()
	}

	for _, option := range options.SpinOptions {
		fmt.Println("option name", option.Name, "type spin default", option.DefaultValue, "min", option.Min, "max", option.Max)
	}
}

func Loop() {
	PrintIdentity()
	fmt.Println()
	PrintOptions()
	fmt.Println()
	fmt.Println("uaiok")

	var stop chan struct{}
	var stopController sync.Once
	searchStopper := func() {
		stopController.Do(func() {
			close(stop)
		})
	}

	pos, _ := ataxx.NewPosition(STARTPOS)
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		input = input[:len(input)-1]

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
			stop = make(chan struct{})
			stopController = sync.Once{}
			go func() {
				Go(*pos, input, stop, searchStopper)
			}()
		} else if strings.HasPrefix(input, "setoption") {
			SetOption(input)
		} else if input == "options" {
			PrintOptions()
		} else if strings.HasPrefix(input, "perft") {
			Perft(*pos, input)
		} else if strings.HasPrefix(input, "hperft") {
			HPerft(*pos, input)
		} else if input == "d" || input == "print" {
			pos.Print()
		} else if input == "stop" {
			searchStopper()
		} else {
			fmt.Println("Unrecognized command!")
		}
	}
}
