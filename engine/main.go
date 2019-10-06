package main

import (
	"fmt"

	"github.com/cpirc/gotaxx/engine/uai"
)

func main() {
	fmt.Println("Gotaxx Ataxx Engine")

	var input string
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
			return
		}

		if input == "uai" {
			uai.Loop()
			break
		} else {
			fmt.Println("Unrecognized protocol!")
		}
	}
}
