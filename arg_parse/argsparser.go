package arg_parse

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type ArgParser struct {
	args         []string
	indexMapping map[string]int
	argsCount    int
	help         string
	currentIndex int
}

func NewArgParser(help string) *ArgParser {
	parser := &ArgParser{
		args:         os.Args,
		indexMapping: map[string]int{},
		argsCount:    len(os.Args),
		help:         strings.TrimSpace(help),
		currentIndex: 1,
	}
	for i, item := range parser.args {
		parser.indexMapping[item] = i
	}

	if parser.indexMapping["--help"] != 0 || parser.indexMapping["-h"] != 0 {
		fmt.Println(parser.help)
		os.Exit(1)
	}
	return parser
}


func arrayContains(array []any, item any) bool {
	for _, element := range array {
		if item == element {
			return true
		}
	}
	return false
}

func (parser *ArgParser) GetArgument(arg Arg) *ArgValue {
	
	value := &ArgValue{}

	if arg.isPositional() {

		value.Value = parser.args[parser.currentIndex]
		parser.currentIndex += 1

		if arg.Choices != nil && len(arg.Choices) > 0 && !arrayContains(arg.Choices, value.Value) {
			log.Fatalf("Invalid value %v for %s. Allowed values are: %v", value.Value, arg.Name, arg.Choices)
		}
		return value
	}

	if parser.indexMapping[arg.Name] != 0 {

		if arg.StoreTrue {
			value.Value = true

		} else {
			value.Value = parser.args[parser.indexMapping[arg.Name]+1]
		}

	} else if parser.indexMapping[arg.ShortName] != 0 {
		
		if arg.StoreTrue {
			value.Value = true

		} else {
			value.Value = parser.args[parser.indexMapping[arg.ShortName]+1]
		}

	} else if arg.Default != nil {
		value.Value = arg.Default

	} else if arg.StoreTrue {
		value.Value = false

	} else if arg.Required {

		if arg.ErrorMessage != "" {
			log.Fatal(arg.ErrorMessage)

		} else {
			log.Fatal("Missing required field: ", arg.Name, " or ", arg.ShortName)
		}
	}

	return value

}
