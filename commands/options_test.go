package commands

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/francescocolleoni/go-ipset/set"
)

func TestOptions(t *testing.T) {
	type test struct {
		handler func() []string
		expects []string
	}

	tests := []test{
		{handler: func() []string { return timeoutOption(0) }},
		{func() []string { return timeoutOption(10) }, []string{"timeout", "10"}},

		{handler: func() []string { return hashSizeOption(0) }},
		{func() []string { return hashSizeOption(10) }, []string{"hashsize", "10"}},

		{handler: func() []string { return maxElementsOption(0) }},
		{func() []string { return maxElementsOption(10) }, []string{"maxelem", "10"}},

		{handler: func() []string { return countersOption(false) }},
		{func() []string { return countersOption(true) }, []string{"counters"}},

		{handler: func() []string { return forceAddOption(false) }},
		{func() []string { return forceAddOption(true) }, []string{"forceadd"}},

		{handler: func() []string { return skbInfoOption(false) }},
		{func() []string { return skbInfoOption(true) }, []string{"skbinfo"}},

		{handler: func() []string { return commentFlagOption(false) }},
		{func() []string { return commentFlagOption(true) }, []string{"comment"}},

		{handler: func() []string { return commentOption("") }},
		{func() []string { return commentOption(`this is a \"comment"`) }, []string{"comment", `"this is a comment"`}},
	}

	for _, test := range tests {
		rawResult := test.handler()
		result := fmt.Sprintf("%v", rawResult)
		expects := fmt.Sprintf("%v", test.expects)

		if result != expects {
			t.Errorf("expectation failed: \"%s\" != \"%s\" (expected)", result, expects)
		}
	}
}

func TestProtocolFamilyOption(t *testing.T) {
	anyProtocol := func() ProtocolFamily {
		all := []ProtocolFamily{ProtocolFamilyDefault, ProtocolFamilyINet, ProtocolFamilyINet6}
		return all[rand.Int()%len(all)]
	}

	tests := map[set.SetType]bool{
		set.SetTypeBitmapIP:       false,
		set.SetTypeBitmapIPMAC:    false,
		set.SetTypeBitmapPort:     false,
		set.SetTypeHashMAC:        false,
		set.SetTypeListSet:        false,
		set.SetTypeHashIP:         true,
		set.SetTypeHashIPMAC:      true,
		set.SetTypeHashNet:        true,
		set.SetTypeHashNetNet:     true,
		set.SetTypeHashNetPort:    true,
		set.SetTypeHashNetPortNet: true,
		set.SetTypeHashNetIFace:   true,
		set.SetTypeHashIPPort:     true,
		set.SetTypeHashIPPortIP:   true,
		set.SetTypeHashIPPortNet:  true,
		set.SetTypeHashIPMark:     true,
	}

	for key, expectsOption := range tests {
		expectation := "[]"
		protocol := anyProtocol()
		result := fmt.Sprintf("%v", protocolFamilyOption(protocol, key))

		if expectsOption && protocol == ProtocolFamilyINet6 {
			expectation = fmt.Sprintf("%v", []string{"family", protocol.String()})
		}

		if result != expectation {
			t.Errorf("expectation failed for key %v: \"%s\" != \"%s\" (expected)", key, result, expectation)
		}
	}
}
