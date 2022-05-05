package commands

import (
	"strings"

	"github.com/francescocolleoni/go-ipset/utilities"
)

// FlushSet defines the ipset flush command.
type FlushSet struct {
	Command CommandName
	Name    string
}

// NewFlushSet returns a flush set command.
func NewFlushSet(name string) *FlushSet {
	return &FlushSet{Command: CommandNameFlush, Name: name}
}

// FlushSet implementation of TranslateToIPSetArgs.
func (c *FlushSet) TranslateToIPSetArgs() []string {
	name := strings.Trim(c.Name, " \n")
	if name == "" {
		return []string{c.Command.String()}
	} else {
		return []string{c.Command.String(), name}
	}
}

// FlushSet implementation of ValidateOptions.
// This function will return true iif name is not empty.
func (c *FlushSet) IncludesMandatoryOptions() bool {
	return strings.Trim(c.Name, " \n") != ""
}

// Run executes a FlushSet command.
func (c *FlushSet) Run() error {
	if out, err := utilities.RunIPSet(c.TranslateToIPSetArgs()...); err != nil {
		return out.Error
	}

	return nil
}
