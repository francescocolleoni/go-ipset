package commands

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"

	"github.com/francescocolleoni/go-ipset/utilities"
)

// ListSet defines the ipset list command.
type ListSet struct {
	Command CommandName
	Name    string
}

// NewListSet returns a list set command.
func NewListSet(name string) *ListSet {
	return &ListSet{Command: CommandNameList, Name: name}
}

// ListSet implementation of TranslateToIPSetArgs.
func (c *ListSet) TranslateToIPSetArgs() []string {
	name := strings.Trim(c.Name, " \n")
	if name == "" {
		return []string{c.Command.String()}
	} else {
		return []string{c.Command.String(), name}
	}
}

// ListSet implementation of ValidateOptions.
// This function will return true iif name is not empty.
func (c *ListSet) IncludesMandatoryOptions() bool {
	return strings.Trim(c.Name, " \n") != ""
}

// Run executes the list set command and returns ip addresses contained in the target set.
func (c *ListSet) Run() ([]string, error) {
	args := c.TranslateToIPSetArgs()
	args = append(args, "-output", "xml")

	out, err := utilities.RunIPSet(args...)
	if err != nil {
		return nil, err
	}

	var xmlOut OxmlIPSets
	if err := xml.Unmarshal([]byte(out.Out), &xmlOut); err != nil {
		return nil, err // Cannot decode output.
	} else if len(xmlOut.Sets) != 1 {
		return nil, fmt.Errorf("expected 1 set, received %d", len(xmlOut.Sets))
	} else if xmlOut.Sets[0].Name != c.Name {
		return nil, errors.New("unexpected target set")
	}

	set := xmlOut.Sets[0]
	members := make([]string, len(set.Members))
	for i, member := range set.Members {
		members[i] = member.Element
	}

	return members, nil
}
