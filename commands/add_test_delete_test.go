package commands

import (
	"fmt"
	"strings"
	"testing"

	"github.com/francescocolleoni/go-ipset/set"
	"github.com/francescocolleoni/go-ipset/utilities"
)

func TestAddSetTranslateToCommandLine(t *testing.T) {
	type test struct {
		command *AddTestDeleteEntry
		args    []string
	}

	const setName = "testset"
	tests := []test{
		{NewAddEntry(setName, set.SetTypeBitmapIP, "1.1.1.1"), []string{"1.1.1.1"}},
		{NewAddEntry(setName, set.SetTypeBitmapIP, "1.1.1.1-2.2.2.2"), []string{"1.1.1.1-2.2.2.2"}},
		{NewAddEntry(setName, set.SetTypeBitmapIP, "1.1.1.1/10"), []string{"1.1.1.1/10"}},

		{NewAddEntry(setName, set.SetTypeBitmapIPMAC, "1.1.1.1"), []string{"1.1.1.1"}},
		{NewAddEntry(setName, set.SetTypeBitmapIPMAC, "1.1.1.1,aa:bb:cc:11:22:33"), []string{"1.1.1.1,aa:bb:cc:11:22:33"}},

		{NewAddEntry(setName, set.SetTypeBitmapPort, "12345"), []string{"12345"}},
		{NewAddEntry(setName, set.SetTypeBitmapPort, "tcp:56789"), []string{"tcp:56789"}},
		{NewAddEntry(setName, set.SetTypeBitmapPort, "12345-56789"), []string{"12345-56789"}},
		{NewAddEntry(setName, set.SetTypeBitmapPort, "tcp:12345-56789"), []string{"tcp:12345-56789"}},

		{NewAddEntry(setName, set.SetTypeHashIP, "1.1.1.1"), []string{"1.1.1.1"}},

		{NewAddEntry(setName, set.SetTypeHashIPMAC, "1.1.1.1,aa:bb:cc:11:22:33"), []string{"1.1.1.1,aa:bb:cc:11:22:33"}},

		{NewAddEntry(setName, set.SetTypeHashIPPort, "1.1.1.1,56789"), []string{"1.1.1.1,56789"}},
		{NewAddEntry(setName, set.SetTypeHashIPPort, "1.1.1.1,tcp:56789"), []string{"1.1.1.1,tcp:56789"}},

		{NewAddEntry(setName, set.SetTypeHashIPPortIP, "1.1.1.1,56789,2.2.2.2"), []string{"1.1.1.1,56789,2.2.2.2"}},
		{NewAddEntry(setName, set.SetTypeHashIPPortIP, "1.1.1.1,tcp:56789,2.2.2.2"), []string{"1.1.1.1,tcp:56789,2.2.2.2"}},

		{NewAddEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,56789,2.2.2.2"), []string{"1.1.1.1,56789,2.2.2.2"}},
		{NewAddEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,56789,2.2.2.2/10"), []string{"1.1.1.1,56789,2.2.2.2/10"}},
		{NewAddEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,tcp:56789,2.2.2.2"), []string{"1.1.1.1,tcp:56789,2.2.2.2"}},
		{NewAddEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,tcp:56789,2.2.2.2/10"), []string{"1.1.1.1,tcp:56789,2.2.2.2/10"}},

		{NewAddEntry(setName, set.SetTypeHashIPMark, "1.1.1.1,10"), []string{"1.1.1.1,10"}},

		{NewAddEntry(setName, set.SetTypeHashMAC, "aa:bb:cc:11:22:33"), []string{"aa:bb:cc:11:22:33"}},

		{NewAddEntry(setName, set.SetTypeHashNet, "1.1.1.1"), []string{"1.1.1.1"}},
		{NewAddEntry(setName, set.SetTypeHashNet, "1.1.1.1/10"), []string{"1.1.1.1/10"}},

		{NewAddEntry(setName, set.SetTypeHashNetNet, "1.1.1.1,2.2.2.2"), []string{"1.1.1.1,2.2.2.2"}},
		{NewAddEntry(setName, set.SetTypeHashNetNet, "1.1.1.1,2.2.2.2/10"), []string{"1.1.1.1,2.2.2.2/10"}},
		{NewAddEntry(setName, set.SetTypeHashNetNet, "1.1.1.1/10,2.2.2.2"), []string{"1.1.1.1/10,2.2.2.2"}},
		{NewAddEntry(setName, set.SetTypeHashNetNet, "1.1.1.1/10,2.2.2.2/10"), []string{"1.1.1.1/10,2.2.2.2/10"}},

		{NewAddEntry(setName, set.SetTypeHashNetPort, "1.1.1.1,56789"), []string{"1.1.1.1,56789"}},
		{NewAddEntry(setName, set.SetTypeHashNetPort, "1.1.1.1,tcp:56789"), []string{"1.1.1.1,tcp:56789"}},
		{NewAddEntry(setName, set.SetTypeHashNetPort, "1.1.1.1/10,56789"), []string{"1.1.1.1/10,56789"}},
		{NewAddEntry(setName, set.SetTypeHashNetPort, "1.1.1.1/10,tcp:56789"), []string{"1.1.1.1/10,tcp:56789"}},

		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,56789,2.2.2.2"), []string{"1.1.1.1,56789,2.2.2.2"}},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,56789,2.2.2.2/10"), []string{"1.1.1.1,56789,2.2.2.2/10"}},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,tcp:56789,2.2.2.2"), []string{"1.1.1.1,tcp:56789,2.2.2.2"}},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,tcp:56789,2.2.2.2/10"), []string{"1.1.1.1,tcp:56789,2.2.2.2/10"}},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,56789,2.2.2.2"), []string{"1.1.1.1/10,56789,2.2.2.2"}},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,56789,2.2.2.2/10"), []string{"1.1.1.1/10,56789,2.2.2.2/10"}},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,tcp:56789,2.2.2.2"), []string{"1.1.1.1/10,tcp:56789,2.2.2.2"}},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,tcp:56789,2.2.2.2/10"), []string{"1.1.1.1/10,tcp:56789,2.2.2.2/10"}},

		{NewAddEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1,eth0"), []string{"1.1.1.1,eth0"}},
		{NewAddEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1,physdev:eth0"), []string{"1.1.1.1,physdev:eth0"}},
		{NewAddEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1/10,eth0"), []string{"1.1.1.1/10,eth0"}},
		{NewAddEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1/10,physdev:eth0"), []string{"1.1.1.1/10,physdev:eth0"}},

		{NewAddListEntry(setName), []string{}},
		{NewAddListEntryBefore(setName, "otherset"), []string{"before", "otherset"}},
		{NewAddListEntryAfter(setName, "otherset"), []string{"after", "otherset"}},
	}

	for i, test := range tests {
		rawResult := test.command.TranslateToIPSetArgs()
		result := fmt.Sprintf("%v", rawResult)

		args := []string{"add", setName}
		args = append(args, test.args...)

		expects := fmt.Sprintf("%v", args)

		if result != expects {
			t.Errorf("expectation failed (%d, %s): %s != %s (expected)", i+1, test.command.Type.String(), result, expects)
		}
	}
}

func TestDeleteSetTranslateToCommandLine(t *testing.T) {
	type test struct {
		command *AddTestDeleteEntry
		args    []string
	}

	const setName = "testset"
	tests := []test{
		{NewDeleteEntry(setName, set.SetTypeBitmapIP, "1.1.1.1"), []string{"1.1.1.1"}},
		{NewDeleteEntry(setName, set.SetTypeBitmapIP, "1.1.1.1-2.2.2.2"), []string{"1.1.1.1-2.2.2.2"}},
		{NewDeleteEntry(setName, set.SetTypeBitmapIP, "1.1.1.1/10"), []string{"1.1.1.1/10"}},

		{NewDeleteEntry(setName, set.SetTypeBitmapIPMAC, "1.1.1.1"), []string{"1.1.1.1"}},
		{NewDeleteEntry(setName, set.SetTypeBitmapIPMAC, "1.1.1.1,aa:bb:cc:11:22:33"), []string{"1.1.1.1,aa:bb:cc:11:22:33"}},

		{NewDeleteEntry(setName, set.SetTypeBitmapPort, "12345"), []string{"12345"}},
		{NewDeleteEntry(setName, set.SetTypeBitmapPort, "tcp:56789"), []string{"tcp:56789"}},
		{NewDeleteEntry(setName, set.SetTypeBitmapPort, "12345-56789"), []string{"12345-56789"}},
		{NewDeleteEntry(setName, set.SetTypeBitmapPort, "tcp:12345-56789"), []string{"tcp:12345-56789"}},

		{NewDeleteEntry(setName, set.SetTypeHashIP, "1.1.1.1"), []string{"1.1.1.1"}},

		{NewDeleteEntry(setName, set.SetTypeHashIPMAC, "1.1.1.1,aa:bb:cc:11:22:33"), []string{"1.1.1.1,aa:bb:cc:11:22:33"}},

		{NewDeleteEntry(setName, set.SetTypeHashIPPort, "1.1.1.1,56789"), []string{"1.1.1.1,56789"}},
		{NewDeleteEntry(setName, set.SetTypeHashIPPort, "1.1.1.1,tcp:56789"), []string{"1.1.1.1,tcp:56789"}},

		{NewDeleteEntry(setName, set.SetTypeHashIPPortIP, "1.1.1.1,56789,2.2.2.2"), []string{"1.1.1.1,56789,2.2.2.2"}},
		{NewDeleteEntry(setName, set.SetTypeHashIPPortIP, "1.1.1.1,tcp:56789,2.2.2.2"), []string{"1.1.1.1,tcp:56789,2.2.2.2"}},

		{NewDeleteEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,56789,2.2.2.2"), []string{"1.1.1.1,56789,2.2.2.2"}},
		{NewDeleteEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,56789,2.2.2.2/10"), []string{"1.1.1.1,56789,2.2.2.2/10"}},
		{NewDeleteEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,tcp:56789,2.2.2.2"), []string{"1.1.1.1,tcp:56789,2.2.2.2"}},
		{NewDeleteEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,tcp:56789,2.2.2.2/10"), []string{"1.1.1.1,tcp:56789,2.2.2.2/10"}},

		{NewDeleteEntry(setName, set.SetTypeHashIPMark, "1.1.1.1,10"), []string{"1.1.1.1,10"}},

		{NewDeleteEntry(setName, set.SetTypeHashMAC, "aa:bb:cc:11:22:33"), []string{"aa:bb:cc:11:22:33"}},

		{NewDeleteEntry(setName, set.SetTypeHashNet, "1.1.1.1"), []string{"1.1.1.1"}},
		{NewDeleteEntry(setName, set.SetTypeHashNet, "1.1.1.1/10"), []string{"1.1.1.1/10"}},

		{NewDeleteEntry(setName, set.SetTypeHashNetNet, "1.1.1.1,2.2.2.2"), []string{"1.1.1.1,2.2.2.2"}},
		{NewDeleteEntry(setName, set.SetTypeHashNetNet, "1.1.1.1,2.2.2.2/10"), []string{"1.1.1.1,2.2.2.2/10"}},
		{NewDeleteEntry(setName, set.SetTypeHashNetNet, "1.1.1.1/10,2.2.2.2"), []string{"1.1.1.1/10,2.2.2.2"}},
		{NewDeleteEntry(setName, set.SetTypeHashNetNet, "1.1.1.1/10,2.2.2.2/10"), []string{"1.1.1.1/10,2.2.2.2/10"}},

		{NewDeleteEntry(setName, set.SetTypeHashNetPort, "1.1.1.1,56789"), []string{"1.1.1.1,56789"}},
		{NewDeleteEntry(setName, set.SetTypeHashNetPort, "1.1.1.1,tcp:56789"), []string{"1.1.1.1,tcp:56789"}},
		{NewDeleteEntry(setName, set.SetTypeHashNetPort, "1.1.1.1/10,56789"), []string{"1.1.1.1/10,56789"}},
		{NewDeleteEntry(setName, set.SetTypeHashNetPort, "1.1.1.1/10,tcp:56789"), []string{"1.1.1.1/10,tcp:56789"}},

		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,56789,2.2.2.2"), []string{"1.1.1.1,56789,2.2.2.2"}},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,56789,2.2.2.2/10"), []string{"1.1.1.1,56789,2.2.2.2/10"}},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,tcp:56789,2.2.2.2"), []string{"1.1.1.1,tcp:56789,2.2.2.2"}},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,tcp:56789,2.2.2.2/10"), []string{"1.1.1.1,tcp:56789,2.2.2.2/10"}},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,56789,2.2.2.2"), []string{"1.1.1.1/10,56789,2.2.2.2"}},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,56789,2.2.2.2/10"), []string{"1.1.1.1/10,56789,2.2.2.2/10"}},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,tcp:56789,2.2.2.2"), []string{"1.1.1.1/10,tcp:56789,2.2.2.2"}},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,tcp:56789,2.2.2.2/10"), []string{"1.1.1.1/10,tcp:56789,2.2.2.2/10"}},

		{NewDeleteEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1,eth0"), []string{"1.1.1.1,eth0"}},
		{NewDeleteEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1,physdev:eth0"), []string{"1.1.1.1,physdev:eth0"}},
		{NewDeleteEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1/10,eth0"), []string{"1.1.1.1/10,eth0"}},
		{NewDeleteEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1/10,physdev:eth0"), []string{"1.1.1.1/10,physdev:eth0"}},

		{NewDeleteListEntry(setName), []string{}},
		{NewDeleteListEntryBefore(setName, "otherset"), []string{"before", "otherset"}},
		{NewDeleteListEntryAfter(setName, "otherset"), []string{"after", "otherset"}},
	}

	for _, test := range tests {
		rawResult := test.command.TranslateToIPSetArgs()
		result := fmt.Sprintf("%v", rawResult)

		args := []string{"del", setName}
		args = append(args, test.args...)

		expects := fmt.Sprintf("%v", args)

		if result != expects {
			t.Errorf("expectation failed: %s != %s (expected)", result, expects)
		}
	}
}

func TestAddSetValidate(t *testing.T) {
	type test struct {
		command *AddTestDeleteEntry
		expects bool
	}

	const setName = "testset"
	tests := []test{
		// Valid.
		{NewAddEntry(setName, set.SetTypeBitmapIP, "1.1.1.1"), true},
		{NewAddEntry(setName, set.SetTypeBitmapIP, "1.1.1.1-2.2.2.2"), true},
		{NewAddEntry(setName, set.SetTypeBitmapIP, "1.1.1.1/10"), true},

		{NewAddEntry(setName, set.SetTypeBitmapIPMAC, "1.1.1.1"), true},
		{NewAddEntry(setName, set.SetTypeBitmapIPMAC, "1.1.1.1,aa:bb:cc:11:22:33"), true},

		{NewAddEntry(setName, set.SetTypeBitmapPort, "12345"), true},
		{NewAddEntry(setName, set.SetTypeBitmapPort, "tcp:56789"), true},
		{NewAddEntry(setName, set.SetTypeBitmapPort, "12345-56789"), true},
		{NewAddEntry(setName, set.SetTypeBitmapPort, "tcp:12345-56789"), true},

		{NewAddEntry(setName, set.SetTypeHashIP, "1.1.1.1"), true},

		{NewAddEntry(setName, set.SetTypeHashIPMAC, "1.1.1.1,aa:bb:cc:11:22:33"), true},

		{NewAddEntry(setName, set.SetTypeHashIPPort, "1.1.1.1,56789"), true},
		{NewAddEntry(setName, set.SetTypeHashIPPort, "1.1.1.1,tcp:56789"), true},

		{NewAddEntry(setName, set.SetTypeHashIPPortIP, "1.1.1.1,56789,2.2.2.2"), true},
		{NewAddEntry(setName, set.SetTypeHashIPPortIP, "1.1.1.1,tcp:56789,2.2.2.2"), true},

		{NewAddEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,56789,2.2.2.2"), true},
		{NewAddEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,56789,2.2.2.2/10"), true},
		{NewAddEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,tcp:56789,2.2.2.2"), true},
		{NewAddEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,tcp:56789,2.2.2.2/10"), true},

		{NewAddEntry(setName, set.SetTypeHashIPMark, "1.1.1.1,10"), true},

		{NewAddEntry(setName, set.SetTypeHashMAC, "aa:bb:cc:11:22:33"), true},

		{NewAddEntry(setName, set.SetTypeHashNet, "1.1.1.1"), true},
		{NewAddEntry(setName, set.SetTypeHashNet, "1.1.1.1/10"), true},

		{NewAddEntry(setName, set.SetTypeHashNetNet, "1.1.1.1,2.2.2.2"), true},
		{NewAddEntry(setName, set.SetTypeHashNetNet, "1.1.1.1,2.2.2.2/10"), true},
		{NewAddEntry(setName, set.SetTypeHashNetNet, "1.1.1.1/10,2.2.2.2"), true},
		{NewAddEntry(setName, set.SetTypeHashNetNet, "1.1.1.1/10,2.2.2.2/10"), true},

		{NewAddEntry(setName, set.SetTypeHashNetPort, "1.1.1.1,56789"), true},
		{NewAddEntry(setName, set.SetTypeHashNetPort, "1.1.1.1,tcp:56789"), true},
		{NewAddEntry(setName, set.SetTypeHashNetPort, "1.1.1.1/10,56789"), true},
		{NewAddEntry(setName, set.SetTypeHashNetPort, "1.1.1.1/10,tcp:56789"), true},

		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,56789,2.2.2.2"), true},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,56789,2.2.2.2/10"), true},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,tcp:56789,2.2.2.2"), true},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,tcp:56789,2.2.2.2/10"), true},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,56789,2.2.2.2"), true},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,56789,2.2.2.2/10"), true},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,tcp:56789,2.2.2.2"), true},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,tcp:56789,2.2.2.2/10"), true},

		{NewAddEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1,eth0"), true},
		{NewAddEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1,physdev:eth0"), true},
		{NewAddEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1/10,eth0"), true},
		{NewAddEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1/10,physdev:eth0"), true},
		{NewAddListEntry(setName), true},
		{NewAddListEntryBefore(setName, "otherset"), true},
		{NewAddListEntryAfter(setName, "otherset"), true},

		// Not valid.
		{NewAddEntry(setName, set.SetTypeBitmapIP, "invalid"), false},
		{NewAddEntry(setName, set.SetTypeBitmapIP, "invalid-invalid"), false},
		{NewAddEntry(setName, set.SetTypeBitmapIP, "invalid/invalidport"), false},

		{NewAddEntry(setName, set.SetTypeBitmapIPMAC, "invalid"), false},
		{NewAddEntry(setName, set.SetTypeBitmapIPMAC, "invalid,invalidmac"), false},

		{NewAddEntry(setName, set.SetTypeBitmapPort, "invalidport"), false},
		{NewAddEntry(setName, set.SetTypeBitmapPort, "tcp:invalidport"), false},
		{NewAddEntry(setName, set.SetTypeBitmapPort, "invalidport-invalidport"), false},
		{NewAddEntry(setName, set.SetTypeBitmapPort, "tcp:invalidport-invalidport"), false},

		{NewAddEntry(setName, set.SetTypeHashIP, "invalid"), false},

		{NewAddEntry(setName, set.SetTypeHashIPMAC, "invalid,invalidmac"), false},

		{NewAddEntry(setName, set.SetTypeHashIPPort, "invalid,invalidport"), false},
		{NewAddEntry(setName, set.SetTypeHashIPPort, "invalid,tcp:invalidport"), false},

		{NewAddEntry(setName, set.SetTypeHashIPPortIP, "invalid,invalidport,invalid"), false},
		{NewAddEntry(setName, set.SetTypeHashIPPortIP, "invalid,tcp:invalidport,invalid"), false},

		{NewAddEntry(setName, set.SetTypeHashIPPortNet, "invalid,invalidport,invalid"), false},
		{NewAddEntry(setName, set.SetTypeHashIPPortNet, "invalid,invalidport,invalid"), false},
		{NewAddEntry(setName, set.SetTypeHashIPPortNet, "invalid,tcp:invalidport,invalid"), false},
		{NewAddEntry(setName, set.SetTypeHashIPPortNet, "invalid,tcp:invalidport,invalid"), false},

		{NewAddEntry(setName, set.SetTypeHashIPMark, "invalid,10"), false},

		{NewAddEntry(setName, set.SetTypeHashMAC, "invalidmac"), false},

		{NewAddEntry(setName, set.SetTypeHashNet, "invalid"), false},
		{NewAddEntry(setName, set.SetTypeHashNet, "invalid/invalidport"), false},

		{NewAddEntry(setName, set.SetTypeHashNetNet, "invalid,invalid"), false},
		{NewAddEntry(setName, set.SetTypeHashNetNet, "invalid,invalid"), false},
		{NewAddEntry(setName, set.SetTypeHashNetNet, "invalid/invalidport,invalid"), false},
		{NewAddEntry(setName, set.SetTypeHashNetNet, "invalid/invalidport,invalid"), false},

		{NewAddEntry(setName, set.SetTypeHashNetPort, "invalid,invalidport"), false},
		{NewAddEntry(setName, set.SetTypeHashNetPort, "invalid,tcp:invalidport"), false},
		{NewAddEntry(setName, set.SetTypeHashNetPort, "invalid/invalidport,invalidport"), false},
		{NewAddEntry(setName, set.SetTypeHashNetPort, "invalid/invalidport,tcp:invalidport"), false},

		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "invalid,invalidport,invalid"), false},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "invalid,invalidport,invalid"), false},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "invalid,tcp:invalidport,invalid"), false},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "invalid,tcp:invalidport,invalid"), false},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "invalid/invalidport,invalidport,invalid"), false},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "invalid/invalidport,invalidport,invalid"), false},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "invalid/invalidport,tcp:invalidport,invalid"), false},
		{NewAddEntry(setName, set.SetTypeHashNetPortNet, "invalid/invalidport,tcp:invalidport,invalid"), false},

		{NewAddEntry(setName, set.SetTypeHashNetIFace, "invalid,eth0"), false},
		{NewAddEntry(setName, set.SetTypeHashNetIFace, "invalid,physdev:eth0"), false},
		{NewAddEntry(setName, set.SetTypeHashNetIFace, "invalid/invalidport,eth0"), false},
		{NewAddEntry(setName, set.SetTypeHashNetIFace, "invalid/invalidport,physdev:eth0"), false},

		{NewAddListEntry(""), false},
		{NewAddListEntryBefore(setName, ""), true},
		{NewAddListEntryBefore("", "otherset"), false},
		{NewAddListEntryBefore(setName, ""), true},
		{NewAddListEntryBefore("", "otherset"), false},
	}

	for i, test := range tests {
		result := test.command.IncludesMandatoryOptions()
		if result != test.expects {
			t.Errorf("expectation %d failed: %v != %v (expected)", i+1, result, test.expects)
		}
	}
}

func TestDeleteSetValidate(t *testing.T) {
	type test struct {
		command *AddTestDeleteEntry
		expects bool
	}

	const setName = "testset"
	tests := []test{
		// Valid.
		{NewDeleteEntry(setName, set.SetTypeBitmapIP, "1.1.1.1"), true},
		{NewDeleteEntry(setName, set.SetTypeBitmapIP, "1.1.1.1-2.2.2.2"), true},
		{NewDeleteEntry(setName, set.SetTypeBitmapIP, "1.1.1.1/10"), true},

		{NewDeleteEntry(setName, set.SetTypeBitmapIPMAC, "1.1.1.1"), true},
		{NewDeleteEntry(setName, set.SetTypeBitmapIPMAC, "1.1.1.1,aa:bb:cc:11:22:33"), true},

		{NewDeleteEntry(setName, set.SetTypeBitmapPort, "12345"), true},
		{NewDeleteEntry(setName, set.SetTypeBitmapPort, "tcp:56789"), true},
		{NewDeleteEntry(setName, set.SetTypeBitmapPort, "12345-56789"), true},
		{NewDeleteEntry(setName, set.SetTypeBitmapPort, "tcp:12345-56789"), true},

		{NewDeleteEntry(setName, set.SetTypeHashIP, "1.1.1.1"), true},

		{NewDeleteEntry(setName, set.SetTypeHashIPMAC, "1.1.1.1,aa:bb:cc:11:22:33"), true},

		{NewDeleteEntry(setName, set.SetTypeHashIPPort, "1.1.1.1,56789"), true},
		{NewDeleteEntry(setName, set.SetTypeHashIPPort, "1.1.1.1,tcp:56789"), true},

		{NewDeleteEntry(setName, set.SetTypeHashIPPortIP, "1.1.1.1,56789,2.2.2.2"), true},
		{NewDeleteEntry(setName, set.SetTypeHashIPPortIP, "1.1.1.1,tcp:56789,2.2.2.2"), true},

		{NewDeleteEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,56789,2.2.2.2"), true},
		{NewDeleteEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,56789,2.2.2.2/10"), true},
		{NewDeleteEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,tcp:56789,2.2.2.2"), true},
		{NewDeleteEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,tcp:56789,2.2.2.2/10"), true},

		{NewDeleteEntry(setName, set.SetTypeHashIPMark, "1.1.1.1,10"), true},

		{NewDeleteEntry(setName, set.SetTypeHashMAC, "aa:bb:cc:11:22:33"), true},

		{NewDeleteEntry(setName, set.SetTypeHashNet, "1.1.1.1"), true},
		{NewDeleteEntry(setName, set.SetTypeHashNet, "1.1.1.1/10"), true},

		{NewDeleteEntry(setName, set.SetTypeHashNetNet, "1.1.1.1,2.2.2.2"), true},
		{NewDeleteEntry(setName, set.SetTypeHashNetNet, "1.1.1.1,2.2.2.2/10"), true},
		{NewDeleteEntry(setName, set.SetTypeHashNetNet, "1.1.1.1/10,2.2.2.2"), true},
		{NewDeleteEntry(setName, set.SetTypeHashNetNet, "1.1.1.1/10,2.2.2.2/10"), true},

		{NewDeleteEntry(setName, set.SetTypeHashNetPort, "1.1.1.1,56789"), true},
		{NewDeleteEntry(setName, set.SetTypeHashNetPort, "1.1.1.1,tcp:56789"), true},
		{NewDeleteEntry(setName, set.SetTypeHashNetPort, "1.1.1.1/10,56789"), true},
		{NewDeleteEntry(setName, set.SetTypeHashNetPort, "1.1.1.1/10,tcp:56789"), true},

		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,56789,2.2.2.2"), true},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,56789,2.2.2.2/10"), true},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,tcp:56789,2.2.2.2"), true},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,tcp:56789,2.2.2.2/10"), true},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,56789,2.2.2.2"), true},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,56789,2.2.2.2/10"), true},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,tcp:56789,2.2.2.2"), true},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,tcp:56789,2.2.2.2/10"), true},

		{NewDeleteEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1,eth0"), true},
		{NewDeleteEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1,physdev:eth0"), true},
		{NewDeleteEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1/10,eth0"), true},
		{NewDeleteEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1/10,physdev:eth0"), true},
		{NewDeleteListEntry(setName), true},
		{NewDeleteListEntryBefore(setName, "otherset"), true},
		{NewDeleteListEntryAfter(setName, "otherset"), true},

		// Not valid.
		{NewDeleteEntry(setName, set.SetTypeBitmapIP, "invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeBitmapIP, "invalid-invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeBitmapIP, "invalid/invalidport"), false},

		{NewDeleteEntry(setName, set.SetTypeBitmapIPMAC, "invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeBitmapIPMAC, "invalid,invalidmac"), false},

		{NewDeleteEntry(setName, set.SetTypeBitmapPort, "invalidport"), false},
		{NewDeleteEntry(setName, set.SetTypeBitmapPort, "tcp:invalidport"), false},
		{NewDeleteEntry(setName, set.SetTypeBitmapPort, "invalidport-invalidport"), false},
		{NewDeleteEntry(setName, set.SetTypeBitmapPort, "tcp:invalidport-invalidport"), false},

		{NewDeleteEntry(setName, set.SetTypeHashIP, "invalid"), false},

		{NewDeleteEntry(setName, set.SetTypeHashIPMAC, "invalid,invalidmac"), false},

		{NewDeleteEntry(setName, set.SetTypeHashIPPort, "invalid,invalidport"), false},
		{NewDeleteEntry(setName, set.SetTypeHashIPPort, "invalid,tcp:invalidport"), false},

		{NewDeleteEntry(setName, set.SetTypeHashIPPortIP, "invalid,invalidport,invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeHashIPPortIP, "invalid,tcp:invalidport,invalid"), false},

		{NewDeleteEntry(setName, set.SetTypeHashIPPortNet, "invalid,invalidport,invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeHashIPPortNet, "invalid,invalidport,invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeHashIPPortNet, "invalid,tcp:invalidport,invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeHashIPPortNet, "invalid,tcp:invalidport,invalid"), false},

		{NewDeleteEntry(setName, set.SetTypeHashIPMark, "invalid,10"), false},

		{NewDeleteEntry(setName, set.SetTypeHashMAC, "invalidmac"), false},

		{NewDeleteEntry(setName, set.SetTypeHashNet, "invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeHashNet, "invalid/invalidport"), false},

		{NewDeleteEntry(setName, set.SetTypeHashNetNet, "invalid,invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeHashNetNet, "invalid,invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeHashNetNet, "invalid/invalidport,invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeHashNetNet, "invalid/invalidport,invalid"), false},

		{NewDeleteEntry(setName, set.SetTypeHashNetPort, "invalid,invalidport"), false},
		{NewDeleteEntry(setName, set.SetTypeHashNetPort, "invalid,tcp:invalidport"), false},
		{NewDeleteEntry(setName, set.SetTypeHashNetPort, "invalid/invalidport,invalidport"), false},
		{NewDeleteEntry(setName, set.SetTypeHashNetPort, "invalid/invalidport,tcp:invalidport"), false},

		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "invalid,invalidport,invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "invalid,invalidport,invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "invalid,tcp:invalidport,invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "invalid,tcp:invalidport,invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "invalid/invalidport,invalidport,invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "invalid/invalidport,invalidport,invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "invalid/invalidport,tcp:invalidport,invalid"), false},
		{NewDeleteEntry(setName, set.SetTypeHashNetPortNet, "invalid/invalidport,tcp:invalidport,invalid"), false},

		{NewDeleteEntry(setName, set.SetTypeHashNetIFace, "invalid,eth0"), false},
		{NewDeleteEntry(setName, set.SetTypeHashNetIFace, "invalid,physdev:eth0"), false},
		{NewDeleteEntry(setName, set.SetTypeHashNetIFace, "invalid/invalidport,eth0"), false},
		{NewDeleteEntry(setName, set.SetTypeHashNetIFace, "invalid/invalidport,physdev:eth0"), false},

		{NewDeleteListEntry(""), false},
		{NewDeleteListEntryAfter(setName, ""), true},
		{NewDeleteListEntryAfter("", "otherset"), false},
		{NewDeleteListEntryAfter(setName, ""), true},
		{NewDeleteListEntryAfter("", "otherset"), false},
	}

	for i, test := range tests {
		result := test.command.IncludesMandatoryOptions()
		if result != test.expects {
			t.Errorf("expectation %d failed: %v != %v (expected)", i+1, result, test.expects)
		}
	}
}

func TestTestSetValidate(t *testing.T) {
	type test struct {
		command *AddTestDeleteEntry
		expects bool
	}

	const setName = "testset"
	tests := []test{
		// Valid.
		{NewTestEntry(setName, set.SetTypeBitmapIP, "1.1.1.1"), true},

		{NewTestEntry(setName, set.SetTypeBitmapIPMAC, "1.1.1.1"), true},
		{NewTestEntry(setName, set.SetTypeBitmapIPMAC, "1.1.1.1,aa:bb:cc:11:22:33"), true},

		{NewTestEntry(setName, set.SetTypeBitmapPort, "12345"), true},
		{NewTestEntry(setName, set.SetTypeBitmapPort, "tcp:56789"), true},

		{NewTestEntry(setName, set.SetTypeHashIP, "1.1.1.1"), true},

		{NewTestEntry(setName, set.SetTypeHashIPMAC, "1.1.1.1,aa:bb:cc:11:22:33"), true},

		{NewTestEntry(setName, set.SetTypeHashIPPort, "1.1.1.1,56789"), true},
		{NewTestEntry(setName, set.SetTypeHashIPPort, "1.1.1.1,tcp:56789"), true},

		{NewTestEntry(setName, set.SetTypeHashIPPortIP, "1.1.1.1,56789,2.2.2.2"), true},
		{NewTestEntry(setName, set.SetTypeHashIPPortIP, "1.1.1.1,tcp:56789,2.2.2.2"), true},

		{NewTestEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,56789,2.2.2.2"), true},
		{NewTestEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,56789,2.2.2.2/10"), true},
		{NewTestEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,tcp:56789,2.2.2.2"), true},
		{NewTestEntry(setName, set.SetTypeHashIPPortNet, "1.1.1.1,tcp:56789,2.2.2.2/10"), true},

		{NewTestEntry(setName, set.SetTypeHashIPMark, "1.1.1.1,10"), true},

		{NewTestEntry(setName, set.SetTypeHashMAC, "aa:bb:cc:11:22:33"), true},

		{NewTestEntry(setName, set.SetTypeHashNet, "1.1.1.1"), true},
		{NewTestEntry(setName, set.SetTypeHashNet, "1.1.1.1/10"), true},

		{NewTestEntry(setName, set.SetTypeHashNetNet, "1.1.1.1,2.2.2.2"), true},
		{NewTestEntry(setName, set.SetTypeHashNetNet, "1.1.1.1,2.2.2.2/10"), true},
		{NewTestEntry(setName, set.SetTypeHashNetNet, "1.1.1.1/10,2.2.2.2"), true},
		{NewTestEntry(setName, set.SetTypeHashNetNet, "1.1.1.1/10,2.2.2.2/10"), true},

		{NewTestEntry(setName, set.SetTypeHashNetPort, "1.1.1.1,56789"), true},
		{NewTestEntry(setName, set.SetTypeHashNetPort, "1.1.1.1,tcp:56789"), true},
		{NewTestEntry(setName, set.SetTypeHashNetPort, "1.1.1.1/10,56789"), true},
		{NewTestEntry(setName, set.SetTypeHashNetPort, "1.1.1.1/10,tcp:56789"), true},

		{NewTestEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,56789,2.2.2.2"), true},
		{NewTestEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,56789,2.2.2.2/10"), true},
		{NewTestEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,tcp:56789,2.2.2.2"), true},
		{NewTestEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1,tcp:56789,2.2.2.2/10"), true},
		{NewTestEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,56789,2.2.2.2"), true},
		{NewTestEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,56789,2.2.2.2/10"), true},
		{NewTestEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,tcp:56789,2.2.2.2"), true},
		{NewTestEntry(setName, set.SetTypeHashNetPortNet, "1.1.1.1/10,tcp:56789,2.2.2.2/10"), true},

		{NewTestEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1,eth0"), true},
		{NewTestEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1,physdev:eth0"), true},
		{NewTestEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1/10,eth0"), true},
		{NewTestEntry(setName, set.SetTypeHashNetIFace, "1.1.1.1/10,physdev:eth0"), true},
		{NewTestListEntry(setName), true},
		{NewTestListEntryBefore(setName, "otherset"), true},
		{NewTestListEntryAfter(setName, "otherset"), true},

		// Not valid.
		{NewTestEntry(setName, set.SetTypeBitmapIP, "invalid"), false},

		{NewTestEntry(setName, set.SetTypeBitmapIP, "1.1.1.1-2.2.2.2"), false}, // Ranges are not supported.
		{NewTestEntry(setName, set.SetTypeBitmapIP, "1.1.1.1/10"), false},      // Ranges are not supported.

		{NewTestEntry(setName, set.SetTypeBitmapIPMAC, "invalid"), false},
		{NewTestEntry(setName, set.SetTypeBitmapIPMAC, "invalid,invalidmac"), false},
		{NewTestEntry(setName, set.SetTypeBitmapPort, "invalidport"), false},
		{NewTestEntry(setName, set.SetTypeBitmapPort, "tcp:invalidport"), false},

		{NewTestEntry(setName, set.SetTypeBitmapPort, "12345-56789"), false},                 // Ranges are not supported.
		{NewTestEntry(setName, set.SetTypeBitmapPort, "tcp:12345-56789"), false},             // Ranges are not supported.
		{NewTestEntry(setName, set.SetTypeBitmapPort, "invalidport-invalidport"), false},     // Ranges are not supported.
		{NewTestEntry(setName, set.SetTypeBitmapPort, "tcp:invalidport-invalidport"), false}, // Ranges are not supported.

		{NewTestEntry(setName, set.SetTypeHashIP, "invalid"), false},

		{NewTestEntry(setName, set.SetTypeHashIPMAC, "invalid,invalidmac"), false},

		{NewTestEntry(setName, set.SetTypeHashIPPort, "invalid,invalidport"), false},
		{NewTestEntry(setName, set.SetTypeHashIPPort, "invalid,tcp:invalidport"), false},

		{NewTestEntry(setName, set.SetTypeHashIPPortIP, "invalid,invalidport,invalid"), false},
		{NewTestEntry(setName, set.SetTypeHashIPPortIP, "invalid,tcp:invalidport,invalid"), false},

		{NewTestEntry(setName, set.SetTypeHashIPPortNet, "invalid,invalidport,invalid"), false},
		{NewTestEntry(setName, set.SetTypeHashIPPortNet, "invalid,invalidport,invalid"), false},
		{NewTestEntry(setName, set.SetTypeHashIPPortNet, "invalid,tcp:invalidport,invalid"), false},
		{NewTestEntry(setName, set.SetTypeHashIPPortNet, "invalid,tcp:invalidport,invalid"), false},

		{NewTestEntry(setName, set.SetTypeHashIPMark, "invalid,10"), false},

		{NewTestEntry(setName, set.SetTypeHashMAC, "invalidmac"), false},

		{NewTestEntry(setName, set.SetTypeHashNet, "invalid"), false},
		{NewTestEntry(setName, set.SetTypeHashNet, "invalid/invalidport"), false},

		{NewTestEntry(setName, set.SetTypeHashNetNet, "invalid,invalid"), false},
		{NewTestEntry(setName, set.SetTypeHashNetNet, "invalid,invalid"), false},
		{NewTestEntry(setName, set.SetTypeHashNetNet, "invalid/invalidport,invalid"), false},
		{NewTestEntry(setName, set.SetTypeHashNetNet, "invalid/invalidport,invalid"), false},

		{NewTestEntry(setName, set.SetTypeHashNetPort, "invalid,invalidport"), false},
		{NewTestEntry(setName, set.SetTypeHashNetPort, "invalid,tcp:invalidport"), false},
		{NewTestEntry(setName, set.SetTypeHashNetPort, "invalid/invalidport,invalidport"), false},
		{NewTestEntry(setName, set.SetTypeHashNetPort, "invalid/invalidport,tcp:invalidport"), false},

		{NewTestEntry(setName, set.SetTypeHashNetPortNet, "invalid,invalidport,invalid"), false},
		{NewTestEntry(setName, set.SetTypeHashNetPortNet, "invalid,invalidport,invalid"), false},
		{NewTestEntry(setName, set.SetTypeHashNetPortNet, "invalid,tcp:invalidport,invalid"), false},
		{NewTestEntry(setName, set.SetTypeHashNetPortNet, "invalid,tcp:invalidport,invalid"), false},
		{NewTestEntry(setName, set.SetTypeHashNetPortNet, "invalid/invalidport,invalidport,invalid"), false},
		{NewTestEntry(setName, set.SetTypeHashNetPortNet, "invalid/invalidport,invalidport,invalid"), false},
		{NewTestEntry(setName, set.SetTypeHashNetPortNet, "invalid/invalidport,tcp:invalidport,invalid"), false},
		{NewTestEntry(setName, set.SetTypeHashNetPortNet, "invalid/invalidport,tcp:invalidport,invalid"), false},

		{NewTestEntry(setName, set.SetTypeHashNetIFace, "invalid,eth0"), false},
		{NewTestEntry(setName, set.SetTypeHashNetIFace, "invalid,physdev:eth0"), false},
		{NewTestEntry(setName, set.SetTypeHashNetIFace, "invalid/invalidport,eth0"), false},
		{NewTestEntry(setName, set.SetTypeHashNetIFace, "invalid/invalidport,physdev:eth0"), false},

		{NewTestListEntry(""), false},
		{NewTestListEntryAfter(setName, ""), true},
		{NewTestListEntryAfter("", "otherset"), false},
		{NewTestListEntryAfter(setName, ""), true},
		{NewTestListEntryAfter("", "otherset"), false},
	}

	for i, test := range tests {
		result := test.command.IncludesMandatoryOptions()
		if result != test.expects {
			t.Errorf("expectation %d failed: %v != %v (expected)", i+1, result, test.expects)
		}
	}
}

func TestAddTestDelete(t *testing.T) {
	type test struct {
		create       *CreateSet
		add          []*AddTestDeleteEntry // 1st command to be run.
		mustExist    []*AddTestDeleteEntry // 2nd command to be run.
		mustNotExist []*AddTestDeleteEntry // 2nd command to be run.
		delete       []*AddTestDeleteEntry // 3rd command to be run.
	}

	const setName = "testset"
	tests := []test{
		{
			create: NewCreateBitmapIP(setName, "192.168.0.0/16", 0, 0, false, false, false),
			add:    []*AddTestDeleteEntry{NewAddEntry(setName, set.SetTypeBitmapIP, "192.168.1.0/24")},
			mustExist: []*AddTestDeleteEntry{
				NewTestEntry(setName, set.SetTypeBitmapIP, "192.168.1.1"),
				NewTestEntry(setName, set.SetTypeBitmapIP, "192.168.1.10"),
				NewTestEntry(setName, set.SetTypeBitmapIP, "192.168.1.100"),
			},
			mustNotExist: []*AddTestDeleteEntry{
				NewTestEntry(setName, set.SetTypeBitmapIP, "192.192.1.1"),
				NewTestEntry(setName, set.SetTypeBitmapIP, "192.192.1.10"),
			},
			delete: []*AddTestDeleteEntry{NewDeleteEntry(setName, set.SetTypeBitmapIP, "192.168.1.1")},
		},

		{
			create: NewCreateHashMAC(setName, 0, 0, 0, false, false, false),
			add: []*AddTestDeleteEntry{
				NewAddEntry(setName, set.SetTypeHashMAC, "01:02:03:04:05:06"),
				NewAddEntry(setName, set.SetTypeHashMAC, "01:02:03:04:05:07"),
			},
			mustExist: []*AddTestDeleteEntry{
				NewTestEntry(setName, set.SetTypeHashMAC, "01:02:03:04:05:06"),
				NewTestEntry(setName, set.SetTypeHashMAC, "01:02:03:04:05:07"),
			},
			mustNotExist: []*AddTestDeleteEntry{
				NewTestEntry(setName, set.SetTypeHashMAC, "01:02:03:04:05:08"),
			},
			delete: []*AddTestDeleteEntry{
				NewTestEntry(setName, set.SetTypeHashMAC, "01:02:03:04:05:06"),
				NewTestEntry(setName, set.SetTypeHashMAC, "01:02:03:04:05:07"),
			},
		},

		{
			create: NewCreateHashNetIFace(setName, ProtocolFamilyDefault, 0, 0, 0, false, false, false),
			add: []*AddTestDeleteEntry{
				NewAddEntry(setName, set.SetTypeHashNetIFace, "192.168.0.0/24,eth0"),
			},
			mustExist: []*AddTestDeleteEntry{
				NewTestEntry(setName, set.SetTypeHashNetIFace, "192.168.0.0/24,eth0"),
			},
			mustNotExist: []*AddTestDeleteEntry{
				NewTestEntry(setName, set.SetTypeHashNetIFace, "10.1.0.0/16,eth1"),
			},
			delete: []*AddTestDeleteEntry{
				NewTestEntry(setName, set.SetTypeHashNetIFace, "192.168.0.0/24,eth0"),
			},
		},
	}

	for i, test := range tests {
		utilities.RunIPSet("destroy", test.create.Name)
		defer utilities.RunIPSet("destroy", test.create.Name)

		createCommand := fmt.Sprintf("%v", test.create.TranslateToIPSetArgs())
		createCommand = strings.ReplaceAll(createCommand, "[", "")
		createCommand = strings.ReplaceAll(createCommand, "]", "")

		t.Logf("[%02d] running create, add, test, delete test set", i+1)
		t.Logf("  - [create] running %s", createCommand)
		if err := test.create.Run(); err != nil {
			t.Errorf("create command %d failed: %v", i+1, err)
			continue
		}

		t.Logf("  - [add]")
		for j, command := range test.add {
			logTestAddTestDeleteCommandLine(t, command, j)
			if err := command.Run(); err != nil {
				t.Errorf("add command %d.%d failed: %v", i+1, j+1, err)
			}
		}

		t.Logf("  - [must exist]")
		for j, command := range test.mustExist {
			logTestAddTestDeleteCommandLine(t, command, j)
			if err := command.Run(); err != nil {
				t.Errorf("containment test %d.%d for %s in %s failed: %v", i+1, j+1, command.Entry, setName, err)
			}
		}

		t.Logf("  - [must not exist]")
		for j, command := range test.mustNotExist {
			logTestAddTestDeleteCommandLine(t, command, j)
			if err := command.Run(); err == nil {
				t.Errorf("containment test %d.%d for %s in %s failed: target should not exist", i+1, j+1, command.Entry, setName)
			}
		}

		t.Log("  - [delete]")
		for j, command := range test.delete {
			logTestAddTestDeleteCommandLine(t, command, j)
			if err := command.Run(); err != nil {
				t.Errorf("delete command %d.%d failed: %v", i+1, j+1, err)
			}
		}
	}
}

// Support functions (used only for tests).
func logTestAddTestDeleteCommandLine(t *testing.T, c *AddTestDeleteEntry, i int) {
	command := fmt.Sprintf("%v", c.TranslateToIPSetArgs())
	command = strings.ReplaceAll(command, "[", "")
	command = strings.ReplaceAll(command, "]", "")

	t.Logf("    - [%02d] running ipset %s", i+1, command)
}
