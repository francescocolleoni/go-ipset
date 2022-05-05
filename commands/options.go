package commands

import (
	"fmt"
	"regexp"
	"strconv"
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

// netmaskOption returns formatted ipset option netmask.
func netmaskOption(value int) []string {
	if value >= 1 && value <= 32 {
		return intOption("netmask", value)
	} else {
		return []string{}
	}
}

// markmaskOption returns formatted ipset option markmask.
func markmaskOption(value int) []string {
	if value >= 0 && value <= 4294967295 {
		return intOption("markmask", value)
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

// sizeOption returns formatted ipset option size.
func sizeOption(value int) []string {
	return intOption("size", value)
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
	if protocol == ProtocolFamilyDefault {
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

// rangeIPOption returns formatted ipset option range.
// Parameter rangeDef format must be either <ip>-<ip> or <ip>/<cidr>.
// IPs must be expressed as IPv4 and CIDR must be a numeric value in [1-32].
// Any other format will return an empty array of arguments.
func rangeIPOption(rangeDef string) []string {
	if regexp.MustCompile(ipRangeMatch).Match([]byte(rangeDef)) {
		return []string{"range", rangeDef}
	} else if regexp.MustCompile(ipCidrRangeMatch).Match([]byte(rangeDef)) {
		return []string{"range", rangeDef}
	} else {
		return []string{}
	}
}

// rangePortOption returns formatted ipset option range.
// Parameter rangeDef format must be <port>-<port>, where port numbers are defined in [0, 65535] range.
// Any other format will return an empty array of arguments.
func rangePortOption(rangeDef string) []string {
	parsePort := func(raw string) int {
		if val, err := strconv.Atoi(raw); err != nil {
			return -1
		} else {
			return val
		}
	}

	if regexp.MustCompile(portRangeMatch).Match([]byte(rangeDef)) {
		ports := strings.Split(rangeDef, "-")

		portA := parsePort(ports[0])
		portB := parsePort(ports[1])

		if portA < 0 || portA > 65535 || portB < 0 || portB > 65535 {
			return []string{}
		} else {
			return []string{"range", rangeDef}
		}
	} else {
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

// matchesTarget returns true if the regex built by joining matchComponents with separator matches target.
func matchesTarget(target, separator string, matchComponents ...string) bool {
	if len(matchComponents) <= 0 {
		return false
	}

	return regexp.MustCompile(
		"^" + strings.Join(matchComponents, separator) + "$",
	).MatchString(target)
}

// Create matches.
const ipRangeMatch = `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)-(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
const ipCidrRangeMatch = `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)/([1-9]|1[0-9]|2[0-9]|3[0-2])$`
const portRangeMatch = `^\d+-\d+$`

// Add / Delete / Test matches.
const ipMatch = `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`
const ipCidrMatch = `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)/([1-9]|1[0-9]|2[0-9]|3[0-2])`
const macMatch = `[a-zA-Z0-9]{2}:[a-zA-Z0-9]{2}:[a-zA-Z0-9]{2}:[a-zA-Z0-9]{2}:[a-zA-Z0-9]{2}:[a-zA-Z0-9]{2}`
const portMatch = `\d+`
const protoMatch = `.+`
const markMatch = `\d+`
const ifaceMatch = `.+`
const physdevMatch = `physdev`
