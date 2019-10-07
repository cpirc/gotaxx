package options

import "errors"

const ComboOptionType = "combo"

type ComboOption struct {
	name         string
	value        string
	defaultValue string
	choices      []string
}

func (s ComboOption) Name() string {
	return s.name
}

func (s ComboOption) Type() string {
	return ComboOptionType
}

func (s ComboOption) Default() string {
	return s.defaultValue
}

func (s *ComboOption) Value() string {
	return s.value
}

func (s ComboOption) Choices() []string {
	return s.choices
}

func (s ComboOption) IsValidChoice(value string) bool {
	for _, choice := range s.choices {
		if choice == value {
			return true
		}
	}
	return false
}

func GetCombo(name string) (string, error) {
	for _, option := range ComboOptions {
		if name == option.Name() {
			return option.value, nil
		}
	}
	return "", errors.New("could not find option " + name)
}

func SetCombo(name string, value string) error {
	for i := 0; i < len(ComboOptions); i++ {
		option := &ComboOptions[i]
		if name == option.Name() {
			if option.IsValidChoice(value) {
				option.value = value
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
		name:         "SearchType",
		value:        "Random",
		defaultValue: "Random",
		choices:      []string{"Random", "MostCaptures"},
	},
}
