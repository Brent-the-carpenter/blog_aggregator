package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}
type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {

	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	cmdHandler, ok := c.registeredCommands[cmd.name]
	if !ok {
		return fmt.Errorf("command:%s does not exists. Please register command", cmd.name)
	}
	return cmdHandler(s, cmd)

}
