package uai

import (
	"fmt"
	"strings"

	"gotaxx/engine/options"
)

func SetOption(input string) {
	parts := strings.Split(input, " ")

	if len(parts) < 5 || parts[1] != "name" || parts[3] != "value" {
		fmt.Println("unable to parse option")
		return
	}

	name, value := parts[2], parts[4]

	err := options.SetCombo(name, value)
	if err != nil {
		fmt.Println(err)
	}
}
