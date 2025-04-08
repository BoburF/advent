package solutions_2015

import (
	"bytes"
	"fmt"
	"maps"
	"strconv"
)

const (
	intialVal int16 = 0
)

func Solution20157(data []byte) {
	linesParsed := bytes.Split(data, []byte("\n"))
	linesParsed = linesParsed[:len(linesParsed)-1]

	parsedCommands := make(map[string][]byte)
	for _, line := range linesParsed {
		variable, command := parseCommand(line)
		parsedCommands[variable] = command
	}

	values := executeCommands(parsedCommands)
	aSignal := values["a"]

	modifiedCommands := make(map[string][]byte)
	maps.Copy(modifiedCommands, parsedCommands)
	modifiedCommands["b"] = fmt.Append(nil, aSignal)

	values2 := executeCommands(modifiedCommands)
	fmt.Println("Result:", values2["a"])
}

func executeCommands(commands map[string][]byte) map[string]uint16 {
	values := make(map[string]uint16)
	var evaluate func(wire string) uint16

	evaluate = func(wire string) uint16 {
		if val, ok := values[wire]; ok {
			return val
		}

		if n, err := strconv.Atoi(wire); err == nil {
			return uint16(n)
		}

		command := string(commands[wire])
		tokens := bytes.Fields([]byte(command))

		var result uint16
		switch len(tokens) {
		case 1:
			result = evaluate(string(tokens[0]))
		case 2:
			result = ^evaluate(string(tokens[1]))
		case 3:
			op := string(tokens[1])
			left := evaluate(string(tokens[0]))
			right := evaluate(string(tokens[2]))

			switch op {
			case "AND":
				result = left & right
			case "OR":
				result = left | right
			case "LSHIFT":
				result = left << right
			case "RSHIFT":
				result = left >> right
			}
		}

		values[wire] = result
		return result
	}

	for wire := range commands {
		evaluate(wire)
	}

	return values
}

func parseCommand(line []byte) (string, []byte) {
	parsedLine := bytes.Fields(line)

	return string(parsedLine[len(parsedLine)-1]), bytes.Join(parsedLine[:len(parsedLine)-2], []byte(" "))
}
