package uai

import "fmt"

func PrintIdentity() {
	fmt.Println("id name Gotaxx")
	fmt.Println("id author kz04px mkchan")
}

func PrintOptions() {
}

func Loop() {
	PrintIdentity()
	fmt.Println()
	PrintOptions()
	fmt.Println("uciok")

	var input string
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
			return
		}

		if input == "quit" {
			break
		} else {
			fmt.Println("Unrecognized command!")
		}
	}
}
