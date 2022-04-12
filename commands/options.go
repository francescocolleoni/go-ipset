package commands

import (
	"fmt"
	"strings"

	"github.com/francescocolleoni/go-ipset/set"
)

// intOption returns the representation of an ipset command numeric option.
// If the value is negative, the result string will be empty.
func intOption(argName string, value int) []string {
	if value <= 0 {
		return []string{}
	} else {
		return []string{argName, fmt.Sprintf("%d", value)}
	}
}

// flagOption returns the representation of an ipset command boolean option.
// If the flag is false, the result string will be empty.
func flagOption(argName string, flag bool) []string {
	if flag {
		return []string{argName}
	} else {
		return []string{}
	}
}

// timeoutOption returns formatted ipset option timeout.
func timeoutOption(value int) []string {
	return intOption("timeout", value)
}

// hashSizeOption returns formatted ipset option hashsize.
func hashSizeOption(value int) []string {
	return intOption("hashsize", value)
}

// maxElementsOption returns formatted ipset option maxelem.
func maxElementsOption(value int) []string {
	return intOption("maxelem", value)
}

// countersOption returns formatted ipset option counters.
func countersOption(flag bool) []string {
	return flagOption("counters", flag)
}

// forceAddOption returns formatted ipset option forceadd.
func forceAddOption(flag bool) []string {
	return flagOption("forceadd", flag)
}

// skbInfoOption returns formatted ipset option skbinfo.
func skbInfoOption(flag bool) []string {
	return flagOption("skbinfo", flag)
}

// protocolFamilyOption returns formatted ipset option skbinfo family.
func protocolFamilyOption(protocol ProtocolFamily, setType set.SetType) []string {
	if protocol != ProtocolFamilyINet6 {
		return []string{} // Default protocol is always inet.
	}

	switch setType {
	case
		set.SetTypeHashIP, set.SetTypeHashIPMAC,
		set.SetTypeHashNet, set.SetTypeHashNetNet, set.SetTypeHashNetPort, set.SetTypeHashNetPortNet, set.SetTypeHashNetIFace,
		set.SetTypeHashIPPort, set.SetTypeHashIPPortIP, set.SetTypeHashIPPortNet, set.SetTypeHashIPMark:
		return []string{"family", protocol.String()}
	default:
		return []string{}
	}
}

// commentFlagOption returns formatted ipset option comment, which should be used with create command.
func commentFlagOption(flag bool) []string {
	return flagOption("comment", flag)
}

// commentOption returns formatted ipset option comment.
func commentOption(comment string) []string {
	comment = strings.Trim(comment, " \n")
	for strings.ContainsAny(comment, `\"`) {
		comment = strings.ReplaceAll(comment, `"`, "")
		comment = strings.ReplaceAll(comment, `\`, "")
	}

	if comment == "" {
		return []string{}
	} else {
		return []string{"comment", fmt.Sprintf(`"%s"`, comment)}
	}
}
