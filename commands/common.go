// Package commands contains definitions of all ipset commands supported by go-ipset.
package commands

import "encoding/xml"

// CommandName defines a command of ipset supported by go-ipset.
type CommandName int

const (
	CommandNameCreate = iota
	CommandNameAdd
	CommandNameDelete
	CommandNameTest
	CommandNameList
	CommandNameFlush
)

// String returns the underlying command name of a given CommandName c.
func (c CommandName) String() string {
	switch c {
	case CommandNameCreate:
		return "create"
	case CommandNameAdd:
		return "add"
	case CommandNameDelete:
		return "del"
	case CommandNameTest:
		return "test"
	case CommandNameList:
		return "list"
	case CommandNameFlush:
		return "flush"

	default:
		return "" // Unsupported command
	}
}

// ProtocolFamily defines the protocol family of IP addresses that will be stored in a set.
type ProtocolFamily int

const (
	ProtocolFamilyDefault = iota // Must be treated as FamilyINet.
	ProtocolFamilyINet
	ProtocolFamilyINet6
)

// String returns the string argument representing Family type.
func (f ProtocolFamily) String() string {
	switch f {
	case ProtocolFamilyINet6:
		return "inet6"
	default:
		return "inet"
	}
}

// Command defines the interface that all go-ipset commands must implement so that they can be run.
type Command interface {
	// TranslateToIPSetArgs returns the list of arguments that will actually be sent to ipset.
	TranslateToIPSetArgs() []string

	// IncludesMandatoryOptions returns true if mandatory options of the command are defined.
	IncludesMandatoryOptions() bool
}

// XML output support.
type OxmlIPSets struct {
	XMLName xml.Name    `xml:"ipsets"`
	Sets    []OxmlIPSet `xml:"ipset"`
}
type OxmlIPSet struct {
	XMLName xml.Name     `xml:"ipset"`
	Name    string       `xml:"name,attr"`
	Type    string       `xml:"type,attr"`
	Members []OxmlMember `xml:"members>member"`
}
type OxmlMember struct {
	XMLName xml.Name `xml:"member"`
	Element string   `xml:"elem"`
}
