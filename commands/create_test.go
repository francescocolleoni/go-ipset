package commands

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/francescocolleoni/go-ipset/set"
	"github.com/francescocolleoni/go-ipset/utilities"
)

func TestCreateSetTranslateToCommandLine(t *testing.T) {
	type test struct {
		command *CreateSet
		setType set.SetType
		args    []string
	}

	const setName = "testset"
	tests := []test{
		{NewCreateBitmapIP(setName, "1.1.1.1-2.2.2.2", 0, 0, false, false, false), set.SetTypeBitmapIP, []string{"range", "1.1.1.1-2.2.2.2"}},
		{
			NewCreateBitmapIP(setName, "1.1.1.1-2.2.2.2", 10, 10, true, true, true), set.SetTypeBitmapIP,
			[]string{"range", "1.1.1.1-2.2.2.2", "netmask", "10", "timeout", "10", "counters", "comment", "skbinfo"},
		},

		{NewCreateBitmapIPMAC(setName, "1.1.1.1-2.2.2.2", 0, false, false, false), set.SetTypeBitmapIPMAC, []string{"range", "1.1.1.1-2.2.2.2"}},
		{
			NewCreateBitmapIPMAC(setName, "1.1.1.1-2.2.2.2", 10, true, true, true), set.SetTypeBitmapIPMAC,
			[]string{"range", "1.1.1.1-2.2.2.2", "timeout", "10", "counters", "comment", "skbinfo"},
		},

		{NewCreateBitmapPort(setName, "1-2", 0, false, false, false), set.SetTypeBitmapPort, []string{"range", "1-2"}},
		{
			NewCreateBitmapPort(setName, "1-2", 10, true, true, true), set.SetTypeBitmapPort,
			[]string{"range", "1-2", "timeout", "10", "counters", "comment", "skbinfo"},
		},

		{NewCreateHashMAC(setName, 0, 0, 0, false, false, false), set.SetTypeHashMAC, []string{}},
		{
			NewCreateHashMAC(setName, 10, 10, 10, true, true, true), set.SetTypeHashMAC,
			[]string{"hashsize", "10", "maxelem", "10", "timeout", "10", "counters", "comment", "skbinfo"},
		},

		{NewCreateHashIP(setName, ProtocolFamilyINet, 0, 0, 0, 0, false, false, false), set.SetTypeHashIP, []string{"family", "inet"}},
		{NewCreateHashIP(setName, ProtocolFamilyINet, 10, 10, 10, 10, true, true, true), set.SetTypeHashIP, []string{"family", "inet"}},
		{NewCreateHashIP(setName, ProtocolFamilyINet6, 0, 0, 0, 0, false, false, false), set.SetTypeHashIP, []string{"family", "inet6"}},
		{NewCreateHashIP(setName, ProtocolFamilyINet6, 10, 10, 10, 10, true, true, true), set.SetTypeHashIP, []string{"family", "inet6"}},
		{NewCreateHashIP(setName, ProtocolFamilyDefault, 0, 0, 0, 0, false, false, false), set.SetTypeHashIP, []string{}},
		{
			NewCreateHashIP(setName, ProtocolFamilyDefault, 10, 10, 10, 10, true, true, true), set.SetTypeHashIP,
			[]string{"hashsize", "10", "maxelem", "10", "netmask", "10", "timeout", "10", "counters", "comment", "skbinfo"},
		},

		{NewCreateHashIPMark(setName, ProtocolFamilyINet, 0, 0, 0, 0, false, false, false), set.SetTypeHashIPMark, []string{"family", "inet"}},
		{NewCreateHashIPMark(setName, ProtocolFamilyINet, 10, 10, 10, 10, true, true, true), set.SetTypeHashIPMark, []string{"family", "inet"}},
		{NewCreateHashIPMark(setName, ProtocolFamilyINet6, 0, 0, 0, 0, false, false, false), set.SetTypeHashIPMark, []string{"family", "inet6"}},
		{NewCreateHashIPMark(setName, ProtocolFamilyINet6, 10, 10, 10, 10, true, true, true), set.SetTypeHashIPMark, []string{"family", "inet6"}},
		{NewCreateHashIPMark(setName, ProtocolFamilyDefault, 0, 0, 0, 0, false, false, false), set.SetTypeHashIPMark, []string{}},
		{
			NewCreateHashIPMark(setName, ProtocolFamilyDefault, 10, 10, 10, 10, true, true, true), set.SetTypeHashIPMark,
			[]string{"markmask", "10", "hashsize", "10", "maxelem", "10", "timeout", "10", "counters", "comment", "skbinfo"},
		},

		{NewCreateHashIPMAC(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashIPMAC, []string{"family", "inet"}},
		{NewCreateHashIPMAC(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashIPMAC, []string{"family", "inet"}},
		{NewCreateHashIPMAC(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashIPMAC, []string{"family", "inet6"}},
		{NewCreateHashIPMAC(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashIPMAC, []string{"family", "inet6"}},
		{NewCreateHashIPMAC(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashIPMAC, []string{}},
		{
			NewCreateHashIPMAC(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashIPMAC,
			[]string{"hashsize", "10", "maxelem", "10", "timeout", "10", "counters", "comment", "skbinfo"},
		},

		{NewCreateList(setName, 0, 0, false, false, false), set.SetTypeListSet, []string{}},
		{
			NewCreateList(setName, 10, 10, true, true, true), set.SetTypeListSet,
			[]string{"size", "10", "timeout", "10", "counters", "comment", "skbinfo"},
		},

		{NewCreateHashIPPort(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashIPPort, []string{"family", "inet"}},
		{NewCreateHashIPPort(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashIPPort, []string{"family", "inet"}},
		{NewCreateHashIPPort(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashIPPort, []string{"family", "inet6"}},
		{NewCreateHashIPPort(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashIPPort, []string{"family", "inet6"}},
		{NewCreateHashIPPort(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashIPPort, []string{}},
		{
			NewCreateHashIPPort(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashIPPort,
			[]string{"hashsize", "10", "maxelem", "10", "timeout", "10", "counters", "comment", "skbinfo"},
		},

		{NewCreateHashIPPortIP(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashIPPortIP, []string{"family", "inet"}},
		{NewCreateHashIPPortIP(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashIPPortIP, []string{"family", "inet"}},
		{NewCreateHashIPPortIP(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashIPPortIP, []string{"family", "inet6"}},
		{NewCreateHashIPPortIP(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashIPPortIP, []string{"family", "inet6"}},
		{NewCreateHashIPPortIP(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashIPPortIP, []string{}},
		{
			NewCreateHashIPPortIP(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashIPPortIP,
			[]string{"hashsize", "10", "maxelem", "10", "timeout", "10", "counters", "comment", "skbinfo"},
		},

		{NewCreateHashIPPortNet(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashIPPortNet, []string{"family", "inet"}},
		{NewCreateHashIPPortNet(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashIPPortNet, []string{"family", "inet"}},
		{NewCreateHashIPPortNet(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashIPPortNet, []string{"family", "inet6"}},
		{NewCreateHashIPPortNet(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashIPPortNet, []string{"family", "inet6"}},
		{NewCreateHashIPPortNet(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashIPPortNet, []string{}},
		{
			NewCreateHashIPPortNet(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashIPPortNet,
			[]string{"hashsize", "10", "maxelem", "10", "timeout", "10", "counters", "comment", "skbinfo"},
		},

		{NewCreateHashNet(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashNet, []string{"family", "inet"}},
		{NewCreateHashNet(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashNet, []string{"family", "inet"}},
		{NewCreateHashNet(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashNet, []string{"family", "inet6"}},
		{NewCreateHashNet(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashNet, []string{"family", "inet6"}},
		{NewCreateHashNet(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashNet, []string{}},
		{
			NewCreateHashNet(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNet,
			[]string{"hashsize", "10", "maxelem", "10", "timeout", "10", "counters", "comment", "skbinfo"},
		},

		{NewCreateHashNetNet(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashNetNet, []string{"family", "inet"}},
		{NewCreateHashNetNet(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashNetNet, []string{"family", "inet"}},
		{NewCreateHashNetNet(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashNetNet, []string{"family", "inet6"}},
		{NewCreateHashNetNet(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashNetNet, []string{"family", "inet6"}},
		{NewCreateHashNetNet(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashNetNet, []string{}},
		{
			NewCreateHashNetNet(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNetNet,
			[]string{"hashsize", "10", "maxelem", "10", "timeout", "10", "counters", "comment", "skbinfo"},
		},

		{NewCreateHashNetPort(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashNetPort, []string{"family", "inet"}},
		{NewCreateHashNetPort(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashNetPort, []string{"family", "inet"}},
		{NewCreateHashNetPort(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashNetPort, []string{"family", "inet6"}},
		{NewCreateHashNetPort(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashNetPort, []string{"family", "inet6"}},
		{NewCreateHashNetPort(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashNetPort, []string{}},
		{
			NewCreateHashNetPort(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNetPort,
			[]string{"hashsize", "10", "maxelem", "10", "timeout", "10", "counters", "comment", "skbinfo"},
		},

		{NewCreateHashNetPortNet(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashNetPortNet, []string{"family", "inet"}},
		{NewCreateHashNetPortNet(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashNetPortNet, []string{"family", "inet"}},
		{NewCreateHashNetPortNet(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashNetPortNet, []string{"family", "inet6"}},
		{NewCreateHashNetPortNet(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashNetPortNet, []string{"family", "inet6"}},
		{NewCreateHashNetPortNet(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashNetPortNet, []string{}},
		{
			NewCreateHashNetPortNet(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNetPortNet,
			[]string{"hashsize", "10", "maxelem", "10", "timeout", "10", "counters", "comment", "skbinfo"},
		},

		{NewCreateHashNetIFace(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashNetIFace, []string{"family", "inet"}},
		{NewCreateHashNetIFace(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashNetIFace, []string{"family", "inet"}},
		{NewCreateHashNetIFace(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashNetIFace, []string{"family", "inet6"}},
		{NewCreateHashNetIFace(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashNetIFace, []string{"family", "inet6"}},
		{NewCreateHashNetIFace(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashNetIFace, []string{}},
		{
			NewCreateHashNetIFace(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNetIFace,
			[]string{"hashsize", "10", "maxelem", "10", "timeout", "10", "counters", "comment", "skbinfo"},
		},
	}

	for _, test := range tests {
		rawResult := test.command.TranslateToIPSetArgs()
		result := fmt.Sprintf("%v", rawResult)

		args := []string{"create", setName, test.setType.String()}
		args = append(args, test.args...)

		expects := fmt.Sprintf("%v", args)

		if result != expects {
			t.Errorf("expectation failed: %s != %s (expected)", result, expects)
		}
	}
}

func TestIncludesMandatoryOptions(t *testing.T) {
	type test struct {
		command *CreateSet
		setType set.SetType
		expects bool
	}

	const setName = "testset"
	tests := []test{
		{NewCreateBitmapIP(setName, "invalid range", 0, 0, false, false, false), set.SetTypeBitmapIP, false},
		{NewCreateBitmapIP(setName, "1.1.1.1-2.2.2.2", 0, 0, false, false, false), set.SetTypeBitmapIP, true},
		{NewCreateBitmapIP(setName, "invalid range", 10, 10, true, true, true), set.SetTypeBitmapIP, false},
		{NewCreateBitmapIP(setName, "1.1.1.1-2.2.2.2", 10, 10, true, true, true), set.SetTypeBitmapIP, true},

		{NewCreateBitmapIPMAC(setName, "invalid range", 0, false, false, false), set.SetTypeBitmapIPMAC, false},
		{NewCreateBitmapIPMAC(setName, "1.1.1.1-2.2.2.2", 0, false, false, false), set.SetTypeBitmapIPMAC, true},
		{NewCreateBitmapIPMAC(setName, "invalid range", 10, true, true, true), set.SetTypeBitmapIPMAC, false},
		{NewCreateBitmapIPMAC(setName, "1.1.1.1-2.2.2.2", 10, true, true, true), set.SetTypeBitmapIPMAC, true},

		{NewCreateBitmapPort(setName, "invalid range", 0, false, false, false), set.SetTypeBitmapPort, false},
		{NewCreateBitmapPort(setName, "1-2", 0, false, false, false), set.SetTypeBitmapPort, true},
		{NewCreateBitmapPort(setName, "invalid range", 10, true, true, true), set.SetTypeBitmapPort, false},
		{NewCreateBitmapPort(setName, "1-2", 10, true, true, true), set.SetTypeBitmapPort, true},

		{NewCreateHashMAC(setName, 0, 0, 0, false, false, false), set.SetTypeHashMAC, true},
		{NewCreateHashMAC(setName, 10, 10, 10, true, true, true), set.SetTypeHashMAC, true},

		{NewCreateHashIP(setName, ProtocolFamilyDefault, 0, 0, 0, 0, false, false, false), set.SetTypeHashIP, true},
		{NewCreateHashIP(setName, ProtocolFamilyDefault, 10, 10, 10, 10, true, true, true), set.SetTypeHashIP, true},
		{NewCreateHashIP(setName, ProtocolFamilyINet, 0, 0, 0, 0, false, false, false), set.SetTypeHashIP, true},
		{NewCreateHashIP(setName, ProtocolFamilyINet6, 0, 0, 0, 0, false, false, false), set.SetTypeHashIP, true},
		{NewCreateHashIP(setName, ProtocolFamilyDefault, 10, 10, 10, 10, true, true, true), set.SetTypeHashIP, true},
		{NewCreateHashIP(setName, ProtocolFamilyINet, 10, 10, 10, 10, true, true, true), set.SetTypeHashIP, true},
		{NewCreateHashIP(setName, ProtocolFamilyINet6, 10, 10, 10, 10, true, true, true), set.SetTypeHashIP, true},

		{NewCreateHashIPMark(setName, ProtocolFamilyDefault, 0, 0, 0, 0, false, false, false), set.SetTypeHashIPMark, true},
		{NewCreateHashIPMark(setName, ProtocolFamilyDefault, 10, 10, 10, 10, true, true, true), set.SetTypeHashIPMark, true},
		{NewCreateHashIPMark(setName, ProtocolFamilyINet, 0, 0, 0, 0, false, false, false), set.SetTypeHashIPMark, true},
		{NewCreateHashIPMark(setName, ProtocolFamilyINet6, 0, 0, 0, 0, false, false, false), set.SetTypeHashIPMark, true},
		{NewCreateHashIPMark(setName, ProtocolFamilyDefault, 10, 10, 10, 10, true, true, true), set.SetTypeHashIPMark, true},
		{NewCreateHashIPMark(setName, ProtocolFamilyINet, 10, 10, 10, 10, true, true, true), set.SetTypeHashIPMark, true},
		{NewCreateHashIPMark(setName, ProtocolFamilyINet6, 10, 10, 10, 10, true, true, true), set.SetTypeHashIPMark, true},

		{NewCreateHashIPMAC(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashIPMAC, true},
		{NewCreateHashIPMAC(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashIPMAC, true},
		{NewCreateHashIPMAC(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashIPMAC, true},
		{NewCreateHashIPMAC(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashIPMAC, true},
		{NewCreateHashIPMAC(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashIPMAC, true},
		{NewCreateHashIPMAC(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashIPMAC, true},
		{NewCreateHashIPMAC(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashIPMAC, true},

		{NewCreateHashIPPort(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashIPPort, true},
		{NewCreateHashIPPort(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashIPPort, true},
		{NewCreateHashIPPort(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashIPPort, true},
		{NewCreateHashIPPort(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashIPPort, true},
		{NewCreateHashIPPort(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashIPPort, true},
		{NewCreateHashIPPort(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashIPPort, true},
		{NewCreateHashIPPort(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashIPPort, true},

		{NewCreateHashIPPortIP(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashIPPortIP, true},
		{NewCreateHashIPPortIP(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashIPPortIP, true},
		{NewCreateHashIPPortIP(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashIPPortIP, true},
		{NewCreateHashIPPortIP(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashIPPortIP, true},
		{NewCreateHashIPPortIP(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashIPPortIP, true},
		{NewCreateHashIPPortIP(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashIPPortIP, true},
		{NewCreateHashIPPortIP(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashIPPortIP, true},

		{NewCreateHashIPPortNet(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashIPPortNet, true},
		{NewCreateHashIPPortNet(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashIPPortNet, true},
		{NewCreateHashIPPortNet(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashIPPortNet, true},
		{NewCreateHashIPPortNet(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashIPPortNet, true},
		{NewCreateHashIPPortNet(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashIPPortNet, true},
		{NewCreateHashIPPortNet(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashIPPortNet, true},
		{NewCreateHashIPPortNet(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashIPPortNet, true},

		{NewCreateHashNet(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashNet, true},
		{NewCreateHashNet(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNet, true},
		{NewCreateHashNet(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashNet, true},
		{NewCreateHashNet(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashNet, true},
		{NewCreateHashNet(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNet, true},
		{NewCreateHashNet(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashNet, true},
		{NewCreateHashNet(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashNet, true},

		{NewCreateHashNetNet(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashNetNet, true},
		{NewCreateHashNetNet(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNetNet, true},
		{NewCreateHashNetNet(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashNetNet, true},
		{NewCreateHashNetNet(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashNetNet, true},
		{NewCreateHashNetNet(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNetNet, true},
		{NewCreateHashNetNet(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashNetNet, true},
		{NewCreateHashNetNet(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashNetNet, true},

		{NewCreateHashNetPort(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashNetPort, true},
		{NewCreateHashNetPort(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNetPort, true},
		{NewCreateHashNetPort(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashNetPort, true},
		{NewCreateHashNetPort(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashNetPort, true},
		{NewCreateHashNetPort(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNetPort, true},
		{NewCreateHashNetPort(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashNetPort, true},
		{NewCreateHashNetPort(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashNetPort, true},

		{NewCreateHashNetPortNet(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashNetPortNet, true},
		{NewCreateHashNetPortNet(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNetPortNet, true},
		{NewCreateHashNetPortNet(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashNetPortNet, true},
		{NewCreateHashNetPortNet(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashNetPortNet, true},
		{NewCreateHashNetPortNet(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNetPortNet, true},
		{NewCreateHashNetPortNet(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashNetPortNet, true},
		{NewCreateHashNetPortNet(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashNetPortNet, true},

		{NewCreateHashNetIFace(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashNetIFace, true},
		{NewCreateHashNetIFace(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNetIFace, true},
		{NewCreateHashNetIFace(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashNetIFace, true},
		{NewCreateHashNetIFace(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashNetIFace, true},
		{NewCreateHashNetIFace(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNetIFace, true},
		{NewCreateHashNetIFace(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashNetIFace, true},
		{NewCreateHashNetIFace(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashNetIFace, true},
	}

	for _, test := range tests {
		if result := test.command.IncludesMandatoryOptions(); result != test.expects {
			t.Errorf("expectation failed: %v != %v (expected)", result, test.expects)
		}
	}
}

func TestCreateSet(t *testing.T) {
	type test struct {
		command        *CreateSet
		setType        set.SetType
		expectsHeaders []string
	}

	const setName = "testset"
	tests := []test{
		{NewCreateBitmapIP(setName, "1.1.1.1-1.1.1.2", 0, 0, false, false, false), set.SetTypeBitmapIP,
			[]string{"range 1.1.1.1-1.1.1.2"}},
		{NewCreateBitmapIP(setName, "1.1.1.1/1", 2, 10, true, true, true), set.SetTypeBitmapIP,
			[]string{"range 0.0.0.0-127.255.255.255", "netmask 2", "timeout 10", "counters", "comment", "skbinfo"}},

		{NewCreateBitmapIPMAC(setName, "1.1.1.1-1.1.1.2", 0, false, false, false), set.SetTypeBitmapIPMAC,
			[]string{"range 1.1.1.1-1.1.1.2"}},
		{NewCreateBitmapIPMAC(setName, "1.1.1.1-1.1.1.2", 10, true, true, true), set.SetTypeBitmapIPMAC,
			[]string{"range 1.1.1.1-1.1.1.2", "timeout 10", "counters", "comment", "skbinfo"}},

		{NewCreateBitmapPort(setName, "1-2", 0, false, false, false), set.SetTypeBitmapPort,
			[]string{"range 1-2"}},
		{NewCreateBitmapPort(setName, "1-2", 10, true, true, true), set.SetTypeBitmapPort,
			[]string{"range 1-2", "timeout 10", "counters", "comment", "skbinfo"}},

		{NewCreateHashMAC(setName, 0, 0, 0, false, false, false), set.SetTypeHashMAC,
			[]string{"hashsize 1024", "maxelem 65536"}},
		{NewCreateHashMAC(setName, 10, 10, 10, true, true, true), set.SetTypeHashMAC,
			[]string{"hashsize 64", "maxelem 10", "timeout 10", "counters", "comment", "skbinfo"}},

		{NewCreateHashIP(setName, ProtocolFamilyDefault, 0, 0, 0, 0, false, false, false), set.SetTypeHashIP,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIP(setName, ProtocolFamilyDefault, 10, 10, 10, 10, true, true, true), set.SetTypeHashIP,
			[]string{"family inet", "hashsize 64", "maxelem 10", "netmask 10", "timeout 10", "counters", "comment", "skbinfo"}},
		{NewCreateHashIP(setName, ProtocolFamilyINet, 0, 0, 0, 0, false, false, false), set.SetTypeHashIP,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIP(setName, ProtocolFamilyINet, 10, 10, 10, 10, true, true, true), set.SetTypeHashIP,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIP(setName, ProtocolFamilyINet6, 0, 0, 0, 0, false, false, false), set.SetTypeHashIP,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIP(setName, ProtocolFamilyINet6, 10, 10, 10, 10, true, true, true), set.SetTypeHashIP,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},

		{NewCreateHashIPMark(setName, ProtocolFamilyDefault, 0, 0, 0, 0, false, false, false), set.SetTypeHashIPMark,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPMark(setName, ProtocolFamilyDefault, 10, 10, 10, 10, true, true, true), set.SetTypeHashIPMark,
			[]string{"family inet", "markmask 0x0000000a", "hashsize 64", "maxelem 10", "timeout 10", "counters", "comment", "skbinfo"}},
		{NewCreateHashIPMark(setName, ProtocolFamilyINet, 0, 0, 0, 0, false, false, false), set.SetTypeHashIPMark,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPMark(setName, ProtocolFamilyINet, 10, 10, 10, 10, true, true, true), set.SetTypeHashIPMark,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPMark(setName, ProtocolFamilyINet6, 0, 0, 0, 0, false, false, false), set.SetTypeHashIPMark,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPMark(setName, ProtocolFamilyINet6, 10, 10, 10, 10, true, true, true), set.SetTypeHashIPMark,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},

		{NewCreateHashIPMAC(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashIPMAC,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPMAC(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashIPMAC,
			[]string{"family inet", "hashsize 64", "maxelem 10", "timeout 10", "counters", "comment", "skbinfo"}},
		{NewCreateHashIPMAC(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashIPMAC,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPMAC(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashIPMAC,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPMAC(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashIPMAC,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPMAC(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashIPMAC,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},

		{NewCreateHashIPPort(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashIPPort,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPPort(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashIPPort,
			[]string{"family inet", "hashsize 64", "maxelem 10", "timeout 10", "counters", "comment", "skbinfo"}},
		{NewCreateHashIPPort(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashIPPort,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPPort(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashIPPort,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPPort(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashIPPort,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPPort(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashIPPort,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},

		{NewCreateHashIPPortIP(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashIPPortIP,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPPortIP(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashIPPortIP,
			[]string{"family inet", "hashsize 64", "maxelem 10", "timeout 10", "counters", "comment", "skbinfo"}},
		{NewCreateHashIPPortIP(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashIPPortIP,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPPortIP(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashIPPortIP,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPPortIP(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashIPPortIP,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPPortIP(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashIPPortIP,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},

		{NewCreateHashIPPortNet(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashIPPortNet,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPPortNet(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashIPPortNet,
			[]string{"family inet", "hashsize 64", "maxelem 10", "timeout 10", "counters", "comment", "skbinfo"}},
		{NewCreateHashIPPortNet(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashIPPortNet,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPPortNet(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashIPPortNet,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPPortNet(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashIPPortNet,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashIPPortNet(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashIPPortNet,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},

		{NewCreateHashNet(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashNet,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNet(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNet,
			[]string{"family inet", "hashsize 64", "maxelem 10", "timeout 10", "counters", "comment", "skbinfo"}},
		{NewCreateHashNet(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashNet,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNet(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashNet,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNet(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashNet,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNet(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashNet,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},

		{NewCreateHashNetNet(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashNetNet,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNetNet(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNetNet,
			[]string{"family inet", "hashsize 64", "maxelem 10", "timeout 10", "counters", "comment", "skbinfo"}},
		{NewCreateHashNetNet(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashNetNet,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNetNet(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashNetNet,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNetNet(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashNetNet,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNetNet(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashNetNet,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},

		{NewCreateHashNetPort(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashNetPort,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNetPort(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNetPort,
			[]string{"family inet", "hashsize 64", "maxelem 10", "timeout 10", "counters", "comment", "skbinfo"}},
		{NewCreateHashNetPort(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashNetPort,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNetPort(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashNetPort,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNetPort(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashNetPort,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNetPort(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashNetPort,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},

		{NewCreateHashNetPortNet(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashNetPortNet,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNetPortNet(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNetPortNet,
			[]string{"family inet", "hashsize 64", "maxelem 10", "timeout 10", "counters", "comment", "skbinfo"}},
		{NewCreateHashNetPortNet(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashNetPortNet,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNetPortNet(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashNetPortNet,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNetPortNet(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashNetPortNet,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNetPortNet(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashNetPortNet,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},

		{NewCreateHashNetIFace(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false), set.SetTypeHashNetIFace,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNetIFace(setName, ProtocolFamilyDefault, 10, 10, 10, true, true, true), set.SetTypeHashNetIFace,
			[]string{"family inet", "hashsize 64", "maxelem 10", "timeout 10", "counters", "comment", "skbinfo"}},
		{NewCreateHashNetIFace(setName, ProtocolFamilyINet, 0, 0, 0, false, false, false), set.SetTypeHashNetIFace,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNetIFace(setName, ProtocolFamilyINet, 10, 10, 10, true, true, true), set.SetTypeHashNetIFace,
			[]string{"family inet", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNetIFace(setName, ProtocolFamilyINet6, 0, 0, 0, false, false, false), set.SetTypeHashNetIFace,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},
		{NewCreateHashNetIFace(setName, ProtocolFamilyINet6, 10, 10, 10, true, true, true), set.SetTypeHashNetIFace,
			[]string{"family inet6", "hashsize 1024", "maxelem 65536"}},
	}

	for i, test := range tests {
		utilities.RunIPSet("destroy", test.command.Name)
		defer utilities.RunIPSet("destroy", test.command.Name)

		command := fmt.Sprintf("%v", test.command.TranslateToIPSetArgs())
		command = strings.ReplaceAll(command, "[", "")
		command = strings.ReplaceAll(command, "]", "")

		t.Logf("[%02d] running ipset %s", i+1, command)
		if err := test.command.Run(); err != nil {
			t.Error(err)
		} else if result, err := utilities.RunIPSet("list", test.command.Name); err != nil {
			t.Error(err)
		} else {
			nameRegex := regexp.MustCompile("Name:.+" + test.command.Name)
			if !nameRegex.Match([]byte(result.Out)) {
				t.Errorf("set name did not match")
				continue
			}

			typeRegex := regexp.MustCompile("Type:.+" + test.setType.String())
			if !typeRegex.Match([]byte(result.Out)) {
				t.Errorf("set type did not match")
				continue
			}

			for _, pattern := range test.expectsHeaders {
				pattern := "Header:.+" + pattern
				if !regexp.MustCompile(pattern).Match([]byte(result.Out)) {
					t.Errorf(`header does not contain pattern "%s"`, pattern)
				}
			}
		}
	}
}
