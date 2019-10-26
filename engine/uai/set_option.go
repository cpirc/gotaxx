package uai

import (
	"fmt"
	"strconv"
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

	valueInt, err := strconv.Atoi(value)
	if err == nil {
		err = options.SetSpin(name, valueInt)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}
