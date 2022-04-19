package commands

import (
	"fmt"
	"testing"

	"github.com/francescocolleoni/go-ipset/set"
)

func TestAddSetTranslateToCommandLine(t *testing.T) {
	type test struct {
		command *AddDeleteEntry
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
		command *AddDeleteEntry
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

func TestAddDeleteSetValidate(t *testing.T) {
	type test struct {
		command *AddDeleteEntry
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
		{NewDeleteListEntry(""), false},
		{NewAddListEntryBefore(setName, ""), true},
		{NewAddListEntryBefore("", "otherset"), false},
		{NewAddListEntryBefore(setName, ""), true},
		{NewAddListEntryBefore("", "otherset"), false},
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
