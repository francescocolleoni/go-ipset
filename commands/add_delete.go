package commands

import (
	"errors"

	"github.com/francescocolleoni/go-ipset/set"
)

// AddDeleteEntry defines ipset commands add or delete.
// This struct serves multiple purposes because ipset arguments for these
// operations are always the same for all set types.
type AddDeleteEntry struct {
	Command CommandName
	Name    string
	Type    set.SetType

	// Options.
	Entry     string // Parsing depends on the set type.
	BeforeSet string // Used only for list:set sets.
	AfterSet  string // Used only for list:set sets.
}

// NewAddEntry returns an add entry command.
// In contrast with create command, adding entries always require one parameter, that will
// be validated and parsed differently according to the source set type.
func NewAddEntry(name string, setType set.SetType, entry string) *AddDeleteEntry {
	return &AddDeleteEntry{Command: CommandNameAdd, Name: name, Type: setType, Entry: entry}
}

// NewAddListEntry returns an add entry command for a list:set set.
func NewAddListEntry(name string) *AddDeleteEntry {
	return &AddDeleteEntry{Command: CommandNameAdd, Name: name, Type: set.SetTypeListSet}
}

// NewAddListEntry returns an add entry command for a list:set set, setting "before" option.
func NewAddListEntryBefore(name, beforeSet string) *AddDeleteEntry {
	return &AddDeleteEntry{Command: CommandNameAdd, Name: name, Type: set.SetTypeListSet, BeforeSet: beforeSet}
}

// NewAddListEntry returns an add entry command for a list:set set, setting "after" option.
func NewAddListEntryAfter(name, afterSet string) *AddDeleteEntry {
	return &AddDeleteEntry{Command: CommandNameAdd, Name: name, Type: set.SetTypeListSet, AfterSet: afterSet}
}

// NewDeleteEntry returns an delete entry command.
// In contrast with create command, removing entries always require one parameter, that will
// be validated and parsed differently according to the source set type.
func NewDeleteEntry(name string, setType set.SetType, entry string) *AddDeleteEntry {
	return &AddDeleteEntry{Command: CommandNameDelete, Name: name, Type: setType, Entry: entry}
}

// NewDeleteListEntry returns an add entry command for a list:set set.
func NewDeleteListEntry(name string) *AddDeleteEntry {
	return &AddDeleteEntry{Command: CommandNameDelete, Name: name, Type: set.SetTypeListSet}
}

// NewDeleteListEntryBefore returns an add entry command for a list:set set, setting "before" option.
func NewDeleteListEntryBefore(name, beforeSet string) *AddDeleteEntry {
	return &AddDeleteEntry{Command: CommandNameDelete, Name: name, Type: set.SetTypeListSet, BeforeSet: beforeSet}
}

// NewDeleteListEntryAfter returns an add entry command for a list:set set, setting "after" option.
func NewDeleteListEntryAfter(name, afterSet string) *AddDeleteEntry {
	return &AddDeleteEntry{Command: CommandNameDelete, Name: name, Type: set.SetTypeListSet, AfterSet: afterSet}
}

// AddDeleteEntry implementation of TranslateToIPSetArgs..
func (c *AddDeleteEntry) TranslateToIPSetArgs() []string {
	makeArgs := func(args ...string) []string {
		out := []string{c.Command.String(), c.Name}
		return append(out, args...)
	}

	switch c.Type {
	case set.SetTypeBitmapIP:
		if matchesTarget(c.Entry, "", ipMatch) ||
			matchesTarget(c.Entry, "-", ipMatch, ipMatch) ||
			matchesTarget(c.Entry, "", ipCidrMatch) {
			return makeArgs(c.Entry)
		}

	case set.SetTypeBitmapIPMAC:
		if matchesTarget(c.Entry, "", ipMatch) ||
			matchesTarget(c.Entry, ",", ipMatch, macMatch) {
			return makeArgs(c.Entry)
		}

	case set.SetTypeBitmapPort:
		if matchesTarget(c.Entry, "", portMatch) ||
			matchesTarget(c.Entry, ":", protoMatch, portMatch) ||
			matchesTarget(c.Entry, "-", portMatch, portMatch) ||
			matchesTarget(c.Entry, ":", protoMatch, portMatch+"-"+portMatch) {
			return makeArgs(c.Entry)
		}

	case set.SetTypeHashIP:
		if matchesTarget(c.Entry, "", ipMatch) {
			return makeArgs(c.Entry)
		}

	case set.SetTypeHashIPMAC:
		if matchesTarget(c.Entry, ",", ipMatch, macMatch) {
			return makeArgs(c.Entry)
		}

	case set.SetTypeHashIPPort:
		if matchesTarget(c.Entry, ",", ipMatch, portMatch) ||
			matchesTarget(c.Entry, ",", ipMatch, protoMatch+":"+portMatch) {
			return makeArgs(c.Entry)
		}

	case set.SetTypeHashIPPortIP:
		if matchesTarget(c.Entry, ",", ipMatch, portMatch, ipMatch) ||
			matchesTarget(c.Entry, ",", ipMatch, protoMatch+":"+portMatch, ipMatch) {
			return makeArgs(c.Entry)
		}

	case set.SetTypeHashIPPortNet:
		if matchesTarget(c.Entry, ",", ipMatch, portMatch, ipMatch) ||
			matchesTarget(c.Entry, ",", ipMatch, portMatch, ipCidrMatch) ||
			matchesTarget(c.Entry, ",", ipMatch, protoMatch+":"+portMatch, ipMatch) ||
			matchesTarget(c.Entry, ",", ipMatch, protoMatch+":"+portMatch, ipCidrMatch) {
			return makeArgs(c.Entry)
		}

	case set.SetTypeHashIPMark:
		if matchesTarget(c.Entry, ",", ipMatch, markMatch) {
			return makeArgs(c.Entry)
		}

	case set.SetTypeHashMAC:
		if matchesTarget(c.Entry, ",", macMatch) {
			return makeArgs(c.Entry)
		}

	case set.SetTypeHashNet:
		if matchesTarget(c.Entry, ",", ipMatch) ||
			matchesTarget(c.Entry, ",", ipCidrMatch) {
			return makeArgs(c.Entry)
		}

	case set.SetTypeHashNetNet:
		if matchesTarget(c.Entry, ",", ipMatch, ipMatch) ||
			matchesTarget(c.Entry, ",", ipMatch, ipCidrMatch) ||
			matchesTarget(c.Entry, ",", ipCidrMatch, ipMatch) ||
			matchesTarget(c.Entry, ",", ipCidrMatch, ipCidrMatch) {
			return makeArgs(c.Entry)
		}

	case set.SetTypeHashNetPort:
		if matchesTarget(c.Entry, ",", ipMatch, portMatch) ||
			matchesTarget(c.Entry, ",", ipMatch, protoMatch+":"+portMatch) ||
			matchesTarget(c.Entry, ",", ipCidrMatch, portMatch) ||
			matchesTarget(c.Entry, ",", ipCidrMatch, protoMatch+":"+portMatch) {
			return makeArgs(c.Entry)
		}

	case set.SetTypeHashNetPortNet:
		if matchesTarget(c.Entry, ",", ipMatch, portMatch, ipMatch) ||
			matchesTarget(c.Entry, ",", ipMatch, portMatch, ipCidrMatch) ||
			matchesTarget(c.Entry, ",", ipMatch, protoMatch+":"+portMatch, ipMatch) ||
			matchesTarget(c.Entry, ",", ipMatch, protoMatch+":"+portMatch, ipCidrMatch) ||
			matchesTarget(c.Entry, ",", ipCidrMatch, portMatch, ipMatch) ||
			matchesTarget(c.Entry, ",", ipCidrMatch, portMatch, ipCidrMatch) ||
			matchesTarget(c.Entry, ",", ipCidrMatch, protoMatch+":"+portMatch, ipMatch) ||
			matchesTarget(c.Entry, ",", ipCidrMatch, protoMatch+":"+portMatch, ipCidrMatch) {
			return makeArgs(c.Entry)
		}

	case set.SetTypeHashNetIFace:
		if matchesTarget(c.Entry, ",", ipMatch, ifaceMatch) ||
			matchesTarget(c.Entry, ",", ipMatch, physdevMatch+":"+ifaceMatch) ||
			matchesTarget(c.Entry, ",", ipCidrMatch, ifaceMatch) ||
			matchesTarget(c.Entry, ",", ipCidrMatch, physdevMatch+":"+ifaceMatch) {
			return makeArgs(c.Entry)
		}

	case set.SetTypeListSet:
		if c.Name == "" {
			return []string{} // Name is always required.
		}

		if c.BeforeSet != "" {
			return makeArgs("before", c.BeforeSet)
		} else if c.AfterSet != "" {
			return makeArgs("after", c.AfterSet)
		} else {
			return makeArgs() // No "before" or "after" args, defaults to plain "add" option.
		}
	}

	return []string{}
}

// AddDeleteEntry implementation of ValidateOptions.
// This function will return true iif result of TranslateToIPSetArgs returns a non-empty array of arguments.
func (c *AddDeleteEntry) IncludesMandatoryOptions() bool {
	return len(c.TranslateToIPSetArgs()) > 0
}

// AddDeleteEntry implementation of Run.
func (c *AddDeleteEntry) Run() error {
	return errors.New("not implemented")
}
