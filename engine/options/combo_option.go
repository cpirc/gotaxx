package options

import "errors"

type ComboOption struct {
	Name         string
	Value        string
	DefaultValue string
	Choices      []string
}

func (s ComboOption) IsValidChoice(value string) bool {
	for _, choice := range s.Choices {
		if choice == value {
			return true
		}
	}
	return false
}

func GetCombo(name string) (string, error) {
	for _, option := range ComboOptions {
		if name == option.Name {
			return option.Value, nil
		}
	}
	return "", errors.New("could not find option " + name)
}

func SetCombo(name string, value string) error {
	for i := 0; i < len(ComboOptions); i++ {
		option := &ComboOptions[i]
		if name == option.Name {
			if option.IsValidChoice(value) {
				option.Value = value
				return nil
			} else {
				return errors.New("value " + value + " not a valid choice")
			}
		}
	}
	return nil
}

var ComboOptions = [...]ComboOption{
	{
		Name:         "SearchType",
		Value:        "AlphaBeta",
		DefaultValue: "AlphaBeta",
		Choices:      []string{"AlphaBeta", "MostCaptures", "Random"},
	},
}
