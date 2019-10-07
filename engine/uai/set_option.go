package uai

import (
	"fmt"
	"strings"

	"github.com/cpirc/gotaxx/engine/options"
)

func SetOption(input string) {
	parts := strings.Split(input, " ")

	if len(parts) < 5 || parts[1] != "name" || parts[3] != "value" {
		fmt.Println("unable to parse option")
		return
	}

	option, err := options.FindByName(parts[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	comboOption, isComboOption := option.(options.ComboOption)
	if isComboOption {
		err := comboOption.SetValue(parts[4])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("value set to", parts[4])
		}
		return
	}
}
