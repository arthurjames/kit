package config

import (
	"fmt"
)

type Option interface {
	fmt.Stringer
}

type EnabledOption struct {
	Name  string
	Value bool
}

type StringOption struct {
	Name  string
	Value string
}

func (eo EnabledOption) String() string {
	if eo.Value {
		return "enable"
	} else {
		return "disable"
	}
}

func (so StringOption) String() string {
	return so.Value
}
