package options

import "errors"

const ComboOptionType = "combo"

type ComboOptioner interface {
	Option
	Value() string
	Choices() []string
	SetValue(string)
}

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

func (s ComboOption) Value() string {
	return s.value
}

func (s ComboOption) Choices() []string {
	return s.choices
}

func (s ComboOption) SetValue(value string) error {
	for _, choice := range s.choices {
		if choice == value {
			s.value = value
			return nil
		}
	}
	return errors.New("value not in choices")
}
