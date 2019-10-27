package options

import (
	"errors"

	"gotaxx/engine/perft_tt"
	"gotaxx/engine/tt"
)

type SpinOption struct {
	Name         string
	Value        int
	DefaultValue int
	Min          int
	Max          int
}

func GetSpin(name string) (int, error) {
	for _, option := range SpinOptions {
		if name == option.Name {
			return option.Value, nil
		}
	}
	return 0, errors.New("could not find option " + name)
}

func SetSpin(name string, value int) error {
	for i := 0; i < len(SpinOptions); i++ {
		option := &SpinOptions[i]
		if name == option.Name {
			if value > option.Min && value < option.Max {
				option.Value = value
				if option.Name == "Hash" {
					tt.TranspositionTable.Resize(option.Value)
					tt.TranspositionTable.Clear()
				} else if option.Name == "PerftHash" {
					perft_tt.PerftTranspositionTable.Resize(option.Value)
					perft_tt.PerftTranspositionTable.Clear()
				}

				return nil
			} else {
				return errors.New("value " + string(value) + " not in range")
			}
		}
	}
	return nil
}

var SpinOptions = [...]SpinOption{
	{
		Name:         "Hash",
		Value:        128,
		DefaultValue: 128,
		Min:          1,
		Max:          1048576,
	}, {
		Name:         "PerftHash",
		Value:        1,
		DefaultValue: 1,
		Min:          1,
		Max:          1048576,
	},
}
