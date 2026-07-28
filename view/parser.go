package view

import (
	"errors"
)

// The parse function takes an input and orders it dependig of what the first command is
func Parse(input string) ([]string, error) {
	tokens, err := tokenize(input)
	if err != nil {
		return nil, err
	}

	if len(tokens) == 0 {
		return nil, errors.New("empty input")
	}

	cmd := tokens[0]
	args := []string{cmd}

	switch cmd {

	case "add":

		if len(tokens) < 2 {
			return nil, errors.New("missing title")
		}

		var title = tokens[1]

		valid, err := isValidTitle(title)

		if valid {
			args = append(args, title)
		} else {
			return nil, err
		}

		for i := 1; i < len(tokens); i++ {
			if tokens[i] == "-d" && i+1 < len(tokens) {
				args = append(args, tokens[i+1])
			}
		}

		return args, nil

	case "list":
		for i := 1; i < len(tokens); i++ {
			if tokens[i] == "-c" {
				args = append(args, "completed")
			}
		}

		return args, nil

	case "remove":

		if len(tokens) < 2 {
			return nil, errors.New("missing title")
		}

		var title = tokens[1]

		valid, err := isValidTitle(title)

		if valid {
			args = append(args, title)
		} else {
			return nil, err
		}

		return args, nil

	case "complete":

		if len(tokens) < 2 {
			return nil, errors.New("missing title")
		}

		var title = tokens[1]

		valid, err := isValidTitle(title)

		if valid {
			args = append(args, title)
		} else {
			return nil, err
		}

		return args, nil

	case "help", "exit":
		return args, nil

	default:
		return nil, errors.New("invalid command")
	}
}

// tokenize() is an auxiliary function to Parse()
// Meant to separate the input into a list of strings for easier access while parsing
//
// -> tokenize is "" sensitive, so anything between "" will be considered the same argument
func tokenize(input string) ([]string, error) {
	var tokens []string
	var current []rune
	inQuotes := false

	for _, c := range input {
		switch c {
		case '"':
			inQuotes = !inQuotes

		case ' ':
			if inQuotes {
				current = append(current, c)
			} else if len(current) > 0 {
				tokens = append(tokens, string(current))
				current = []rune{}
			}

		default:
			current = append(current, c)
		}
	}

	if inQuotes {
		return nil, errors.New("unclosed quotes")
	}

	if len(current) > 0 {
		tokens = append(tokens, string(current))
	}

	return tokens, nil
}

func parseAddCommand(tokens []string) (args []string, err error) {
	var title = tokens[1]

	valid, err := isValidTitle(title)

	if valid {
		args = append(args, title)
	} else {
		return nil, err
	}

	for i := 1; i < len(tokens); i++ {
		if tokens[i] == "-d" && i+1 < len(tokens) {
			args = append(args, tokens[i+1])
		}
	}

	return args, nil
}

func parseRemoveCommand(tokens []string) (args []string, err error) {
	var title = tokens[1]

	valid, err := isValidTitle(title)

	if valid {
		args = append(args, title)
	} else {
		return nil, err
	}

	return args, nil
}

func parseListCommand(tokens []string) (args []string, err error) {
	for i := 1; i < len(tokens); i++ {
		if tokens[i] == "-c" {
			args = append(args, "completed")
		}
	}

	return args, nil
}

func parseCompleteCommand(tokens []string) (args []string, err error) {
	var title = tokens[1]

	valid, err := isValidTitle(title)

	if valid {
		args = append(args, title)
	} else {
		return nil, err
	}

	return args, nil
}

// isValidTitle validates if the title is empty or not
func isValidTitle(title string) (bool, error) {
	if title == "" {
		return false, errors.New("missing title")
	}

	return true, nil
}
