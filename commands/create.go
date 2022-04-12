// Package commands contains definitions of all ipset commands supported by go-ipset.
package commands

import (
	"github.com/francescocolleoni/go-ipset/set"
	"github.com/francescocolleoni/go-ipset/utilities"
)

// CreateSet defines the ipset create command.
type CreateSet struct {
	Command CommandName
	Name    string
	Type    set.SetType

	// Options.
	Timeout        int
	HashSize       int // Only for hash sets.
	MaxElements    int // Only for hash sets.
	UseCounters    bool
	UseSKBInfo     bool
	ForceAdd       bool
	AllowsComments bool
	ProtocolFamily ProtocolFamily // Only for hash sets, excluding hash:mac.
}

// NewCreateSet returns an instance of a CreateSet command.
func NewCreateSet(
	name string, setType set.SetType,
	timeout, hashSize, maxElements int,
	useCounters, useSKBInfo, forceAdd, allowsComments bool,
	protocolFamily ProtocolFamily,
) *CreateSet {
	return &CreateSet{
		Command: CommandNameCreate, Name: name, Type: setType,
		Timeout: timeout, HashSize: hashSize, MaxElements: maxElements, ProtocolFamily: protocolFamily,
		UseCounters: useCounters, UseSKBInfo: useSKBInfo, ForceAdd: forceAdd, AllowsComments: allowsComments,
	}
}

// CreateSet implementation of TranslateToCommandLine.
func (c *CreateSet) TranslateToCommandLine() [][]string {
	out := []string{c.Command.String(), c.Name, c.Type.String()}

	out = append(out, timeoutOption(c.Timeout)...)
	out = append(out, hashSizeOption(c.HashSize)...)
	out = append(out, maxElementsOption(c.MaxElements)...)
	out = append(out, countersOption(c.UseCounters)...)
	out = append(out, skbInfoOption(c.UseSKBInfo)...)
	out = append(out, forceAddOption(c.ForceAdd)...)
	out = append(out, protocolFamilyOption(c.ProtocolFamily, c.Type)...)
	out = append(out, commentFlagOption(c.AllowsComments)...)

	return [][]string{out}
}

// CreateSet implementation of Run.
func (c *CreateSet) Run() error {
	args := c.TranslateToCommandLine()[0] // Expects just one line of arguments.
	_, err := utilities.RunIPSet(args...)
	return err
}
