package commands

import (
	"encoding/xml"
	"strings"

	"github.com/francescocolleoni/go-ipset/utilities"
)

// DestroySet defines the ipset destroy command.
type DestroySet struct {
	Command CommandName
	Name    string
}

// ExistsSet defines the ipset exists command.
type ExistsSet struct {
	Command CommandName
	Name    string
}

// NewDestroySet returns a destroy set command.
func NewDestroySet(name string) *DestroySet {
	return &DestroySet{Command: CommandNameDestroy, Name: name}
}

// NewExistsSet returns a "set exists" command.
func NewExistsSet(name string) *ExistsSet {
	return &ExistsSet{Command: CommandNameExists, Name: name}
}

// DestroySet implementation of TranslateToIPSetArgs.
func (c *DestroySet) TranslateToIPSetArgs() []string {
	name := strings.Trim(c.Name, " \n")
	if name == "" {
		return []string{c.Command.String()}
	} else {
		return []string{c.Command.String(), name}
	}
}

// ExistsSet implementation of TranslateToIPSetArgs.
func (c *ExistsSet) TranslateToIPSetArgs() []string {
	name := strings.Trim(c.Name, " \n")
	if name == "" {
		return []string{c.Command.String()}
	} else {
		return []string{c.Command.String(), name}
	}
}

// DestroySet implementation of ValidateOptions.
// This function will return true iif name is not empty.
func (c *DestroySet) IncludesMandatoryOptions() bool {
	return strings.Trim(c.Name, " \n") != ""
}

// ExistsSet implementation of ValidateOptions.
// This function will return true iif name is not empty.
func (c *ExistsSet) IncludesMandatoryOptions() bool {
	return strings.Trim(c.Name, " \n") != ""
}

// Run executes a DestroySet command.
func (c *DestroySet) Run() error {
	if out, err := utilities.RunIPSet(c.TranslateToIPSetArgs()...); err != nil {
		return out.Error
	}

	return nil
}

// Run executes a ExistsSet command.
func (c *ExistsSet) Run() bool {
	args := c.TranslateToIPSetArgs()
	args = append(args, "-output", "xml")

	out, err := utilities.RunIPSet(args...)
	if err != nil {
		return false
	}

	var xmlOut OxmlIPSets
	if err := xml.Unmarshal([]byte(out.Out), &xmlOut); err != nil {
		return false // Cannot decode output, but no error returned.
	}

	for _, set := range xmlOut.Sets {
		if set.Name == c.Name {
			return true
		}
	}

	return false
}
