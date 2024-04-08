package arg_parse

import (
	"fmt"
	"strconv"
	"strings"
)

type ArgType int

const (
	NAMED ArgType = iota
	POSTIONAL
)

type Arg struct {
	Default      any
	ShortName    string
	ErrorMessage string
	Name         string
	Choices      []any
	Required     bool
	Help         string
	// ArgsCount    uint
	StoreTrue bool
}

func (arg *Arg) isPositional() bool {
	return !(strings.HasPrefix(arg.Name, "-") || strings.HasPrefix(arg.ShortName, "-"))
}

type ArgValue struct {
	Value any
	error error
}


func (v *ArgValue) IntoString() (string, error) {
	if v.error != nil {
		return "", nil
	}
	return fmt.Sprint(v.Value), nil
}

func (v *ArgValue) IntoInt() (int, error) {
	val, err := v.IntoString()
	if err != nil {
		return 0, nil
	}
	return strconv.Atoi(val)
}

func (v *ArgValue) IntoFloat() (float64, error) {
	val, err := v.IntoString()
	if err != nil {
		return 0, nil
	}
	return strconv.ParseFloat(val, 64)
}

func (v *ArgValue) IntoBool() (bool, error) {
	val, err := v.IntoString()
	if err != nil {
		return false, nil
	}
	return strconv.ParseBool(val)
}
