package commands

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/francescocolleoni/go-ipset/utilities"
)

func TestDestroyExistsSetTranslateToCommandLine(t *testing.T) {
	expectOK := 10 + rand.Int()%5
	expectNoName := 10 + rand.Int()%5
	tests := make([]string, expectOK+expectNoName)

	// Generate tests with positive expectations.
	for i := 0; i < expectOK; i++ {
		tests[i] = fakeSetName(1 + rand.Int()%10)
	}

	// Generate tests with negative expectations.
	for i := 0; i < expectNoName; i++ {
		tests[expectOK+i] = ""
	}

	for i, test := range tests {
		destroyArgs := fmt.Sprintf("%v", NewDestroySet(test).TranslateToIPSetArgs())
		existsArgs := fmt.Sprintf("%v", NewExistsSet(test).TranslateToIPSetArgs())

		rawExpectDestroyArgs := []string{"destroy"}
		rawExpectExistsArgs := []string{"-L"}

		if test != "" {
			rawExpectDestroyArgs = append(rawExpectDestroyArgs, test)
			rawExpectExistsArgs = append(rawExpectExistsArgs, test)
		}

		expectDestroyArgs := fmt.Sprintf("%v", rawExpectDestroyArgs)
		if destroyArgs != expectDestroyArgs {
			t.Errorf("expectation failed (%d, destroy): %s != %s (expected)", i+1, destroyArgs, expectDestroyArgs)
		}

		expectExistsArgs := fmt.Sprintf("%v", rawExpectExistsArgs)
		if existsArgs != expectExistsArgs {
			t.Errorf("expectation failed (%d, exists): %s != %s (expected)", i+1, existsArgs, expectExistsArgs)
		}
	}
}
func TestDestroyExistsSetValidate(t *testing.T) {
	expectTrue := 10 + rand.Int()%5
	expectFalse := 10 + rand.Int()%5
	tests := make([]string, expectTrue+expectFalse)

	// Generate tests with positive expectations.
	for i := 0; i < expectTrue; i++ {
		tests[i] = fakeSetName(1 + rand.Int()%10)
	}

	// Generate tests with negative expectations.
	for i := 0; i < expectFalse; i++ {
		tests[expectTrue+i] = ""
	}

	for i, test := range tests {
		expects := (test != "")
		destroyResult := NewDestroySet(test).IncludesMandatoryOptions()
		existsResult := NewExistsSet(test).IncludesMandatoryOptions()

		if destroyResult != expects {
			t.Errorf("expectation %d failed (destroy): %v != %v (expected)", i+1, destroyResult, expects)
		}

		if existsResult != expects {
			t.Errorf("expectation %d failed (exists): %v != %v (expected)", i+1, existsResult, expects)
		}
	}
}
func TestDestroyExistsSet(t *testing.T) {
	const setName = "testset"
	utilities.RunIPSet("destroy", setName)
	defer utilities.RunIPSet("destroy", setName)

	// Create a new set: exists command must return true.
	if err := NewCreateBitmapIP(setName, "1.1.1.1-1.1.1.2", 0, 0, false, false, false).Run(); err != nil {
		t.Errorf("cannot setup test (create set): %v", err)
		return
	}

	if exists := NewExistsSet(setName).Run(); !exists {
		t.Errorf("test set %s should exist, but command returned false", setName)
		return
	}

	// Destroy the set: exists command must return false.
	if err := NewDestroySet(setName).Run(); err != nil {
		t.Errorf("cannot destroy set %s: %v", setName, err)
		return
	}

	if exists := NewExistsSet(setName).Run(); exists {
		t.Errorf("test set %s should not exist, but command returned true", setName)
		return
	}
}
