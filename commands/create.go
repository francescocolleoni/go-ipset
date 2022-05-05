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
	IPRange        string // Used only for bitmap:ip and bitmap:ip,mac sets.
	PortRange      string // Used only for bitmap:port sets.
	NetMask        int    // Used only for bitmap:ip, hash:ip sets.
	MarkMask       int    // Used only for hash:ip,mark sets.
	HashSize       int    // Only for hash sets.
	MaxElements    int    // Only for hash sets.
	Size           int    // Only for list sets.
	Timeout        int
	UseCounters    bool
	UseSKBInfo     bool
	ForceAdd       bool
	AllowsComments bool
	ProtocolFamily ProtocolFamily // Only for hash sets, excluding hash:mac.
}

// TranslateToCommandLine support functions for Create*type* sets.

func (c *CreateSet) translateCreateBitmapIPToCommandLine() []string {
	// bitmap:ip
	// range fromip-toip|ip/cidr [ netmask cidr ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]
	out := rangeIPOption(c.IPRange)
	out = append(out, netmaskOption(c.NetMask)...)
	return out
}
func (c *CreateSet) translateCreateBitmapIPMACToCommandLine() []string {
	// bitmap:ip,mac
	// range fromip-toip|ip/cidr [ timeout value ] [ counters ] [ comment ] [ skbinfo ]
	return rangeIPOption(c.IPRange)
}
func (c *CreateSet) translateCreateBitmapPortToCommandLine() []string {
	// bitmap:port
	// range fromport-toport [ timeout value ] [ counters ] [ comment ] [ skbinfo ]
	return rangePortOption(c.PortRange)
}
func (c *CreateSet) translateCreateHashIPToCommandLine() []string {
	// hash:ip
	// [ family { inet | inet6 } ] | [ hashsize value ] [ maxelem value ] [ netmask cidr ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]
	out := protocolFamilyOption(c.ProtocolFamily, c.Type)
	out = append(out, hashSizeOption(c.HashSize)...)
	out = append(out, maxElementsOption(c.MaxElements)...)
	out = append(out, netmaskOption(c.NetMask)...)
	return out
}
func (c *CreateSet) translateCreateHashMACToCommandLine() []string {
	// hash:ip,mac
	// [ hashsize value ] [ maxelem value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]
	out := hashSizeOption(c.HashSize)
	out = append(out, maxElementsOption(c.MaxElements)...)
	return out
}
func (c *CreateSet) translateCreateHashIPMarkToCommandLine() []string {
	// hash:ip,mark
	// [ family { inet | inet6 } ] | [ markmask value ] [ hashsize value ] [ maxelem value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]
	out := protocolFamilyOption(c.ProtocolFamily, c.Type)
	out = append(out, markmaskOption(c.MarkMask)...)
	out = append(out, hashSizeOption(c.HashSize)...)
	out = append(out, maxElementsOption(c.MaxElements)...)
	return out
}
func (c *CreateSet) translateCreateListToCommandLine() []string {
	// list:set
	// [ size value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]
	return sizeOption(c.Size)
}
func (c *CreateSet) translateCreateHashOtherToCommandLine() []string {
	// [ family { inet | inet6 } ] | [ hashsize value ] [ maxelem value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]
	out := protocolFamilyOption(c.ProtocolFamily, c.Type)
	out = append(out, hashSizeOption(c.HashSize)...)
	out = append(out, maxElementsOption(c.MaxElements)...)
	return out
}

// CreateSet implementation of TranslateToIPSetArgs.
// CreateSet with bitmap:ip, bitmap:ip,mac and bitmap:port will use the appropriate IP or port range option.
// For sets like:
// - hash:ip
// - hash:ip,mac
// - hash:net
// - hash:net,net
// - hash:ip,port
// - hash:net,port
// - hash:ip,port,ip
// - hash:ip,port,net
// - hash:net,port,net
// - hash:net,iface
// only option family will be returned if not default.
func (c *CreateSet) TranslateToIPSetArgs() []string {
	out := []string{c.Command.String(), c.Name, c.Type.String()}

	switch c.Type {
	case set.SetTypeBitmapIP:
		out = append(out, c.translateCreateBitmapIPToCommandLine()...)
	case set.SetTypeBitmapIPMAC:
		out = append(out, c.translateCreateBitmapIPMACToCommandLine()...)
	case set.SetTypeBitmapPort:
		out = append(out, c.translateCreateBitmapPortToCommandLine()...)
	case set.SetTypeHashIP:
		out = append(out, c.translateCreateHashIPToCommandLine()...)
	case set.SetTypeHashMAC:
		out = append(out, c.translateCreateHashMACToCommandLine()...)
	case set.SetTypeHashIPMark:
		out = append(out, c.translateCreateHashIPMarkToCommandLine()...)
	case set.SetTypeListSet:
		out = append(out, c.translateCreateListToCommandLine()...)
	case set.SetTypeHashIPPort,
		set.SetTypeHashIPPortIP,
		set.SetTypeHashIPPortNet,
		set.SetTypeHashIPMAC,
		set.SetTypeHashNet,
		set.SetTypeHashNetNet,
		set.SetTypeHashNetPort,
		set.SetTypeHashNetPortNet,
		set.SetTypeHashNetIFace:
		out = append(out, c.translateCreateHashOtherToCommandLine()...)
	}

	out = append(out, timeoutOption(c.Timeout)...)
	out = append(out, countersOption(c.UseCounters)...)
	out = append(out, commentFlagOption(c.AllowsComments)...)
	out = append(out, skbInfoOption(c.UseSKBInfo)...)

	return out
}

// CreateSet implementation of ValidateOptions.
// CreateSet with bitmap:ip, bitmap:ip,mac and bitmap:port require either IP or port ranges.
// All other variants will return true (all options are optional).
// This function does NOT validate argument format.
func (c *CreateSet) IncludesMandatoryOptions() bool {
	switch c.Type {
	case set.SetTypeBitmapIP, set.SetTypeBitmapIPMAC:
		rangeDef := rangeIPOption(c.IPRange)
		return len(rangeDef) > 0
	case set.SetTypeBitmapPort:
		rangeDef := rangePortOption(c.PortRange)
		return len(rangeDef) > 0
	default:
		return true
	}
}

// Run executes a CreateSet command.
func (c *CreateSet) Run() error {
	if out, err := utilities.RunIPSet(c.TranslateToIPSetArgs()...); err != nil {
		return out.Error
	}

	return nil
}

// New Create*type* set.
// NewCreateBitmapIP returns a create command for a SetTypeBitmapIP set.
func NewCreateBitmapIP(name, ipRange string, netMask, timeout int, useCounters, allowsComments, useSKBInfo bool) *CreateSet {
	// bitmap:ip
	// range fromip-toip|ip/cidr [ netmask cidr ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]
	out := newCreateCommand(name, set.SetTypeBitmapIP)
	out.IPRange = ipRange
	out.NetMask = netMask
	out.Timeout = timeout
	out.UseCounters = useCounters
	out.AllowsComments = allowsComments
	out.UseSKBInfo = useSKBInfo
	return out
}

// NewCreateBitmapIPMAC returns a create command for a SetTypeBitmapIPMAC set.
func NewCreateBitmapIPMAC(name, ipRange string, timeout int, useCounters, allowsComments, useSKBInfo bool) *CreateSet {
	// bitmap:ip,mac
	// range fromip-toip|ip/cidr [ timeout value ] [ counters ] [ comment ] [ skbinfo ]
	out := newCreateCommand(name, set.SetTypeBitmapIPMAC)
	out.IPRange = ipRange
	out.Timeout = timeout
	out.UseCounters = useCounters
	out.AllowsComments = allowsComments
	out.UseSKBInfo = useSKBInfo
	return out
}

// NewCreateBitmapPort returns a create command for a SetTypeBitmapPort set.
func NewCreateBitmapPort(name, portRange string, timeout int, useCounters, allowsComments, useSKBInfo bool) *CreateSet {
	// bitmap:port
	// range fromport-toport [ timeout value ] [ counters ] [ comment ] [ skbinfo ]
	out := newCreateCommand(name, set.SetTypeBitmapPort)
	out.PortRange = portRange
	out.Timeout = timeout
	out.UseCounters = useCounters
	out.AllowsComments = allowsComments
	out.UseSKBInfo = useSKBInfo
	return out
}

// NewCreateHashIP returns a create command for a SetTypeHashIP set.
func NewCreateHashIP(name string, protocolFamily ProtocolFamily, hashSize, maxElements, netMask, timeout int, useCounters, allowsComments, useSKBInfo bool) *CreateSet {
	// hash:ip
	// [ family { inet | inet6 } ] | [ hashsize value ] [ maxelem value ] [ netmask cidr ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]
	out := newCreateCommand(name, set.SetTypeHashIP)
	if protocolFamily != ProtocolFamilyDefault {
		out.ProtocolFamily = protocolFamily
	} else {
		out.HashSize = hashSize
		out.MaxElements = maxElements
		out.NetMask = netMask
		out.Timeout = timeout
		out.UseCounters = useCounters
		out.AllowsComments = allowsComments
		out.UseSKBInfo = useSKBInfo
	}
	return out
}

// NewCreateHashMAC returns a create command for a SetTypeHashMAC set.
func NewCreateHashMAC(name string, hashSize, maxElements, timeout int, useCounters, allowsComments, useSKBInfo bool) *CreateSet {
	// hash:ip,mac
	// [ hashsize value ] [ maxelem value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]
	out := newCreateCommand(name, set.SetTypeHashMAC)
	out.HashSize = hashSize
	out.MaxElements = maxElements
	out.Timeout = timeout
	out.UseCounters = useCounters
	out.AllowsComments = allowsComments
	out.UseSKBInfo = useSKBInfo
	return out
}

// NewCreateHashIPMark returns a create command for a SetTypeHashIPMark set.
func NewCreateHashIPMark(name string, protocolFamily ProtocolFamily, markMask, hashSize, maxElements, timeout int, useCounters, allowsComments, useSKBInfo bool) *CreateSet {
	// hash:ip,mark
	// [ family { inet | inet6 } ] | [ markmask value ] [ hashsize value ] [ maxelem value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]
	out := newCreateCommand(name, set.SetTypeHashIPMark)
	if protocolFamily != ProtocolFamilyDefault {
		out.ProtocolFamily = protocolFamily
	} else {
		out.MarkMask = markMask
		out.HashSize = hashSize
		out.MaxElements = maxElements
		out.Timeout = timeout
		out.UseCounters = useCounters
		out.AllowsComments = allowsComments
		out.UseSKBInfo = useSKBInfo
	}
	return out
}

// NewCreateList returns a create command for a SetTypeListSet set.
func NewCreateList(name string, size, timeout int, useCounters, allowsComments, useSKBInfo bool) *CreateSet {
	// list:set
	// [ size value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]
	out := newCreateCommand(name, set.SetTypeListSet)
	out.Size = size
	out.Timeout = timeout
	out.UseCounters = useCounters
	out.AllowsComments = allowsComments
	out.UseSKBInfo = useSKBInfo
	return out
}

// NewCreateHashIPPort returns a create command for a SetTypeHashIPPort set.
func NewCreateHashIPPort(name string, protocolFamily ProtocolFamily, hashSize, maxElements, timeout int, useCounters, allowsComments, useSKBInfo bool) *CreateSet {
	return newCreateHashOther(name, set.SetTypeHashIPPort, protocolFamily, hashSize, maxElements, timeout, useCounters, allowsComments, useSKBInfo)
}

// NewCreateHashIPPortIP returns a create command for a SetTypeHashIPPortIP set.
func NewCreateHashIPPortIP(name string, protocolFamily ProtocolFamily, hashSize, maxElements, timeout int, useCounters, allowsComments, useSKBInfo bool) *CreateSet {
	return newCreateHashOther(name, set.SetTypeHashIPPortIP, protocolFamily, hashSize, maxElements, timeout, useCounters, allowsComments, useSKBInfo)
}

// NewCreateHashIPPortNet returns a create command for a SetTypeHashIPPortNet set.
func NewCreateHashIPPortNet(name string, protocolFamily ProtocolFamily, hashSize, maxElements, timeout int, useCounters, allowsComments, useSKBInfo bool) *CreateSet {
	return newCreateHashOther(name, set.SetTypeHashIPPortNet, protocolFamily, hashSize, maxElements, timeout, useCounters, allowsComments, useSKBInfo)
}

// NewCreateHashIPMAC returns a create command for a SetTypeHashIPMAC set.
func NewCreateHashIPMAC(name string, protocolFamily ProtocolFamily, hashSize, maxElements, timeout int, useCounters, allowsComments, useSKBInfo bool) *CreateSet {
	return newCreateHashOther(name, set.SetTypeHashIPMAC, protocolFamily, hashSize, maxElements, timeout, useCounters, allowsComments, useSKBInfo)
}

// NewCreateHashNet returns a create command for a SetTypeHashNet set.
func NewCreateHashNet(name string, protocolFamily ProtocolFamily, hashSize, maxElements, timeout int, useCounters, allowsComments, useSKBInfo bool) *CreateSet {
	return newCreateHashOther(name, set.SetTypeHashNet, protocolFamily, hashSize, maxElements, timeout, useCounters, allowsComments, useSKBInfo)
}

// NewCreateHashNetNet returns a create command for a SetTypeHashNetNet set.
func NewCreateHashNetNet(name string, protocolFamily ProtocolFamily, hashSize, maxElements, timeout int, useCounters, allowsComments, useSKBInfo bool) *CreateSet {
	return newCreateHashOther(name, set.SetTypeHashNetNet, protocolFamily, hashSize, maxElements, timeout, useCounters, allowsComments, useSKBInfo)
}

// NewCreateHashNetPort returns a create command for a SetTypeHashNetPort set.
func NewCreateHashNetPort(name string, protocolFamily ProtocolFamily, hashSize, maxElements, timeout int, useCounters, allowsComments, useSKBInfo bool) *CreateSet {
	return newCreateHashOther(name, set.SetTypeHashNetPort, protocolFamily, hashSize, maxElements, timeout, useCounters, allowsComments, useSKBInfo)
}

// NewCreateHashNetPortNet returns a create command for a SetTypeHashNetPortNet set.
func NewCreateHashNetPortNet(name string, protocolFamily ProtocolFamily, hashSize, maxElements, timeout int, useCounters, allowsComments, useSKBInfo bool) *CreateSet {
	return newCreateHashOther(name, set.SetTypeHashNetPortNet, protocolFamily, hashSize, maxElements, timeout, useCounters, allowsComments, useSKBInfo)
}

// NewCreateHashNetIFace returns a create command for a SetTypeHashNetIFace set.
func NewCreateHashNetIFace(name string, protocolFamily ProtocolFamily, hashSize, maxElements, timeout int, useCounters, allowsComments, useSKBInfo bool) *CreateSet {
	return newCreateHashOther(name, set.SetTypeHashNetIFace, protocolFamily, hashSize, maxElements, timeout, useCounters, allowsComments, useSKBInfo)
}

// Support functions.
func newCreateCommand(name string, setType set.SetType) *CreateSet {
	return &CreateSet{Command: CommandNameCreate, Name: name, Type: setType}
}
func newCreateHashOther(name string, setType set.SetType, protocolFamily ProtocolFamily, hashSize, maxElements, timeout int, useCounters, allowsComments, useSKBInfo bool) *CreateSet {
	// [ family { inet | inet6 } ] | [ hashsize value ] [ maxelem value ] [ timeout value ] [ counters ] [ comment ] [ skbinfo ]
	out := newCreateCommand(name, setType)
	if protocolFamily != ProtocolFamilyDefault {
		out.ProtocolFamily = protocolFamily
	} else {
		out.HashSize = hashSize
		out.MaxElements = maxElements
		out.Timeout = timeout
		out.UseCounters = useCounters
		out.AllowsComments = allowsComments
		out.UseSKBInfo = useSKBInfo
	}
	return out
}
