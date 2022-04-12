// Package commands contains definitions of all ipset commands supported by go-ipset.
package commands

// CommandName defines a command of ipset supported by go-ipset.
type CommandName int

const (
	CommandNameCreate = iota
	CommandNameAdd
	CommandNameList
)

// String returns the underlying command name of a given CommandName c.
func (c CommandName) String() string {
	switch c {
	case CommandNameCreate:
		return "create"
	case CommandNameAdd:
		return "add"
	case CommandNameList:
		return "list"

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
	// This method returns an array of arrays of arguments, supporting multiple calls to ipset.
	TranslateToIPSetArgs() [][]string

	// Run sends to ipset a command, which may consist of many calls to ipset.
	Run() error
}
