package builder

import (
	"errors"
	"strings"
)

type StreamCommand struct {
	streams map[string]string
}

func NewStreamCommand() *StreamCommand {
	return &StreamCommand{
		streams: make(map[string]string),
	}
}

func (s *StreamCommand) Build(cmd string) (string, error) {
	var ends []string

	cmd = strings.TrimSpace(cmd)
	if len(cmd) == 0 {
		return "", errors.New("no command provided")
	}

	if cmd[len(cmd)-1] == ';' {
		cmd = cmd[:len(cmd)-1]
		ends = append(ends, ";")
	}

	cmds := strings.Fields(cmd)
	stream := ""

	if len(cmds) > 1 {
		if cmds[len(cmds)-2] == "into" {
			stream = cmds[len(cmds)-1]
			cmds = cmds[:len(cmds)-2]
		}
	}

	if v, ok := s.streams[cmds[0]]; ok {
		cmds[0] = v
	} else {
		// Check if cmds[0] is a simple variable string
		// Error if it is.
	}

	output := strings.Join(cmds, " ")
	if stream != "" {
		s.streams[stream] = output
	}

	outEnds := strings.Join(ends, " ")
	return output + " " + outEnds, nil
}
