package options

import "errors"

type Option interface {
	Name() string
	Type() string
	Default() string
}

func FindByName(name string) (Option, error) {
	for _, option := range OPTIONS {
		if name == option.Name() {
			return option, nil
		}
	}
	return nil, errors.New("unsupported option " + name)
}

var OPTIONS = [...]Option{
	ComboOption{
		name:         "SearchType",
		value:        "Random",
		defaultValue: "Random",
		choices:      []string{"Random", "MostCaptures"},
	},
}
