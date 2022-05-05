package commands

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/francescocolleoni/go-ipset/set"
	"github.com/francescocolleoni/go-ipset/utilities"
)

func TestListSetTranslateToCommandLine(t *testing.T) {
	type test struct {
		command *ListSet
		setName string
	}

	expectOK := 10 + rand.Int()%5
	expectNoName := 10 + rand.Int()%5
	tests := make([]test, expectOK+expectNoName)

	// Generate tests with positive expectations.
	for i := 0; i < expectOK; i++ {
		setName := fakeSetName(1 + rand.Int()%10)
		tests[i] = test{command: NewListSet(setName), setName: setName}
	}

	// Generate tests with negative expectations.
	for i := 0; i < expectNoName; i++ {
		tests[expectOK+i] = test{command: NewListSet("")}
	}

	for i, test := range tests {
		rawResult := test.command.TranslateToIPSetArgs()
		result := fmt.Sprintf("%v", rawResult)

		args := []string{"list"}
		if test.setName != "" {
			args = append(args, test.setName)
		}

		expects := fmt.Sprintf("%v", args)

		if result != expects {
			t.Errorf("expectation failed (%d): %s != %s (expected)", i+1, result, expects)
		}
	}
}
func TestListSetValidate(t *testing.T) {
	type test struct {
		command *ListSet
		expects bool
	}

	expectTrue := 10 + rand.Int()%5
	expectFalse := 10 + rand.Int()%5
	tests := make([]test, expectTrue+expectFalse)

	// Generate tests with positive expectations.
	for i := 0; i < expectTrue; i++ {
		tests[i] = test{command: NewListSet(fakeSetName(1 + rand.Int()%10)), expects: true}
	}

	// Generate tests with negative expectations.
	for i := 0; i < expectFalse; i++ {
		tests[expectTrue+i] = test{command: NewListSet("")}
	}

	for i, test := range tests {
		result := test.command.IncludesMandatoryOptions()
		if result != test.expects {
			t.Errorf("expectation %d failed: %v != %v (expected)", i+1, result, test.expects)
		}
	}
}
func TestListSet(t *testing.T) {
	// creo un set, poi verifico l'elenco degli indirizzi contenuti.
	// I risultati sono sempre restituiti in ordine alfabetico per ragioni di leggibilità.

	setName := "testset"
	utilities.RunIPSet("destroy", setName)
	defer utilities.RunIPSet("destroy", setName)

	// Create a new set, then add some addresses.
	if err := NewCreateBitmapIP(setName, "192.168.0.0/16", 0, 0, false, false, false).Run(); err != nil {
		t.Errorf("cannot setup test, failed to create set: %v", err)
		return
	}

	// Add some entries.
	if err := NewAddEntry(setName, set.SetTypeBitmapIP, "192.168.1.0/24").Run(); err != nil {
		t.Errorf("cannot setup test, failed to add required addresses: %v", err)
		return
	}

	// List addresses.
	out, err := NewListSet(setName).Run()
	if err != nil {
		t.Errorf("list set failed: %v", err)
		return
	}

	// Test first list.
	// Build expectations array.
	expectedAddressesCount := 256
	expectedAddressesRemoveMargin := expectedAddressesCount / 8
	expectedAddresses := make([]string, expectedAddressesCount)
	for i := 0; i < expectedAddressesCount; i++ {
		expectedAddresses[i] = fmt.Sprintf("192.168.1.%d", i)
	}

	expect := fmt.Sprintf("%v", expectedAddresses)
	result := fmt.Sprintf("%v", out)
	if result != expect {
		t.Errorf("unexpected list content")
		t.Logf("received %s", result)
		t.Logf("expected %s", expect)
	}

	// Remove addresses from the existing set, then list its content again.
	removeAtIndex := expectedAddressesRemoveMargin + (rand.Int() % (len(expectedAddresses) - expectedAddressesRemoveMargin))
	removedAddress := expectedAddresses[removeAtIndex]

	if err := NewDeleteEntry(setName, set.SetTypeBitmapIP, removedAddress).Run(); err != nil {
		t.Errorf("cannot continue test, failed to remove required addresses: %v", err)
		return
	}

	// List addresses.
	out, err = NewListSet(setName).Run()
	if err != nil {
		t.Errorf("list set failed: %v", err)
		return
	}

	expect = strings.ReplaceAll(expect, removedAddress, "")
	expect = strings.ReplaceAll(expect, "  ", " ")
	result = fmt.Sprintf("%v", out)
	if result != expect {
		t.Errorf("unexpected list content")
		t.Logf("received %s", result)
		t.Logf("expected %s", expect)
	}
}

// Support.
func fakeSetName(length int) string {
	if length <= 0 {
		return ""
	}

	charsPool := "abcdefghijklmnopqrstuvwxyz"
	charsPoolSize := len(charsPool)
	lastCharIndex := charsPoolSize - 1

	selectChar := func() string {
		index := rand.Int() % charsPoolSize
		if index >= lastCharIndex {
			return charsPool[index:]
		} else {
			return charsPool[index : index+1]
		}
	}

	if length == 1 {
		return selectChar()
	}

	out := ""
	for i := 0; i < length; i++ {
		out += selectChar()
	}

	return out
}
