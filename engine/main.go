package main

import (
	"bufio"
	"fmt"
	"os"

	"gotaxx/engine/uai"
	"gotaxx/libs/ataxx"
)

func main() {
	fmt.Println("Gotaxx Ataxx Engine")
	ataxx.InitZobristKeys()

	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		input = input[:len(input)-1]

		if input == "uai" {
			uai.Loop()
			break
		} else {
			fmt.Println("Unrecognized protocol!")
		}
	}
}
