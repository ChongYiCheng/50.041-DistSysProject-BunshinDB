package Utils

import (
	"fmt"
	"errors"
    "os"
	)

func ParseCommandLine(command string) ([]string, error) {
	//Finite state machine to handle arguments with white spaces enclosed within quotes
	//Handles escaped stuff too
	var args []string
	state := "start"
	current := ""
	quote := "\""
	escapeNext := true
	for i := 0; i < len(command); i++ {
		c := command[i]

		if state == "quotes" {
			if string(c) != quote {
				current += string(c)
			} else {
				args = append(args, current)
				current = ""
				state = "start"
			}
			continue
		}

		if (escapeNext) {
			current += string(c)
			escapeNext = false
			continue
		}

		if (c == '\\') {
			escapeNext = true
			continue
		}

		if c == '"' || c == '\'' {
			state = "quotes"
			quote = string(c)
			continue
		}

		if state == "arg" {
			if c == ' ' || c == '\t' {
				args = append(args, current)
				current = ""
				state = "start"
			} else {
				current += string(c)
			}
			continue
		}

		if c != ' ' && c != '\t' {
			state = "arg"
			current += string(c)
		}
	}

	if state == "quotes" {
		return []string{}, errors.New(fmt.Sprintf("Unclosed quote in command line: %s", command))


	}

	if current != "" {
		args = append(args, current)
	}

	return args, nil
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}
