package commands

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/francescocolleoni/go-ipset/set"
	"github.com/francescocolleoni/go-ipset/utilities"
)

func TestFlushSetTranslateToCommandLine(t *testing.T) {
	type test struct {
		command *FlushSet
		setName string
	}

	expectOK := 10 + rand.Int()%5
	expectNoName := 10 + rand.Int()%5
	tests := make([]test, expectOK+expectNoName)

	// Generate tests with positive expectations.
	for i := 0; i < expectOK; i++ {
		setName := fakeSetName(1 + rand.Int()%10)
		tests[i] = test{command: NewFlushSet(setName), setName: setName}
	}

	// Generate tests with negative expectations.
	for i := 0; i < expectNoName; i++ {
		tests[expectOK+i] = test{command: NewFlushSet("")}
	}

	for i, test := range tests {
		rawResult := test.command.TranslateToIPSetArgs()
		result := fmt.Sprintf("%v", rawResult)

		args := []string{"flush"}
		if test.setName != "" {
			args = append(args, test.setName)
		}

		expects := fmt.Sprintf("%v", args)

		if result != expects {
			t.Errorf("expectation failed (%d): %s != %s (expected)", i+1, result, expects)
		}
	}
}
func TestFlushSetValidate(t *testing.T) {
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
func TestFlushSet(t *testing.T) {
	setName := "testset"
	utilities.RunIPSet("destroy", setName)
	defer utilities.RunIPSet("destroy", setName)

	// Create a new set, then add some addresses.
	if err := NewCreateHashMAC(setName, 0, 0, 0, false, false, false).Run(); err != nil {
		t.Errorf("cannot setup test, failed to create set: %v", err)
		return
	}

	// Add some entries.
	macAddresses := []string{"01:02:03:04:05:06", "01:02:03:04:05:07"}
	for i, entry := range macAddresses {
		if err := NewAddEntry(setName, set.SetTypeHashMAC, entry).Run(); err != nil {
			t.Errorf("cannot setup test (%d), failed to add required addresses: %v", i+1, err)
			return
		}
	}

	// List addresses.
	out, err := NewListSet(setName).Run()
	if err != nil {
		t.Errorf("list set failed: %v", err)
		return
	}

	// Test list before flush.
	sort.Strings(macAddresses)
	sort.Strings(out)

	expect := fmt.Sprintf("%v", macAddresses)
	result := fmt.Sprintf("%v", out)
	if result != expect {
		t.Errorf("unexpected list content")
		t.Logf("received %s", result)
		t.Logf("expected %s", expect)
	}

	// Flush the set.
	if err := NewFlushSet(setName).Run(); err != nil {
		t.Errorf("flush set failed: %v", err)
		return
	}

	// List addresses again.
	if out, err := NewListSet(setName).Run(); err != nil {
		t.Errorf("list set failed: %v", err)
	} else if len(out) > 0 {
		t.Errorf("flush command did not flush target set")
	}
}
