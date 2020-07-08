package builder

import (
	"errors"
	"strings"
)

type Command struct {
	str strings.Builder
}

func NewCommand() *Command {
	return &Command{
		str: strings.Builder{},
	}
}

// Add adds a command to the builder and returns build command after
// it's addition and a flag to denote if the command is finished.
func (cmd *Command) Add(str string) (string, bool, error) {
	str = strings.TrimSpace(str)

	if cmd.str.Len() != 0 && str[0] != '|' {
		return "", false, errors.New("mistake command")
	}

	cmd.str.WriteString(" ")

	if str[len(str)-1:] == ";" {
		// The command building is finished
		cmd.str.WriteString(str[:len(str)-1])

		s := cmd.str.String()
		cmd.str.Reset()
		return s, true, nil
	} else {
		cmd.str.WriteString(str)
	}

	return cmd.str.String(), false, nil
}

func (cmd *Command) Reset() {
	cmd.str.Reset()
}
