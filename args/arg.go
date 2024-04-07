package args

import (
	action "cli/args/arg_types"
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
	Action       action.Action
	Help         string
	// ArgsCount    uint
	StoreTrue bool
}

func (arg *Arg) isPositional() bool {
	return !(strings.HasPrefix(arg.Name, "-") || strings.HasPrefix(arg.ShortName, "-"))
}

type ArgValue struct {
	Value any
}

// func NewArg() *ArgParser {
// 	return &ArgParser{

// 	}
// }

func (v *ArgValue) IntoString() string {
	return fmt.Sprint(v.Value)
}

func (v *ArgValue) IntoInt() (int, error) {
	return strconv.Atoi(v.IntoString())
}

func (v *ArgValue) IntoFloat() (float64, error) {
	return strconv.ParseFloat(v.IntoString(), 64)
}

func (v *ArgValue) IntoBool() (bool, error) {
	return strconv.ParseBool(v.IntoString())
}