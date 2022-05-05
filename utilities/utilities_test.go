package utilities

import "testing"

func TestVersion(t *testing.T) {
	// Important: this test relies on ipset being available on the environment running tests.

	if version, err := Version(); err != nil {
		t.Errorf("expectation failed, version returned an error: %v", err)
	} else if version == "" {
		t.Errorf("expectation failed, version returned an empty string")
	}
}
func TestIPSetIsAvailable(t *testing.T) {
	// Important: this test relies on ipset being available on the environment running tests.

	if isAvailable := IPSetIsAvailable(); !isAvailable {
		t.Errorf("expectation failed: ipset is not available")
	}
}

func TestIPSetError(t *testing.T) {
	// Important: this test relies on ipset being available on the environment running tests.

	if out, err := RunIPSet("dummycommand"); err == nil {
		t.Error("expectation failed: ipset should return an error")
	} else if out.Error.Error() != `ipset returned error "No command specified: unknown argument dummycommand"` {
		t.Errorf(`unexpected error message: received %s`, out.Error.Error())
	} else if out.In != "ipset dummycommand" {
		t.Errorf("unexpected input: received %s", out.In)
	}
}

func TestRunCommand(t *testing.T) {
	type test struct {
		cmd          string
		args         []string
		expectsError bool
	}

	tests := []test{
		{"dummycommand", []string{}, true},
		{"echo", []string{"test"}, false},
	}

	for _, test := range tests {
		_, err := runCommand(test.cmd)
		if err != nil && !test.expectsError {
			t.Errorf("expectation failed for command %s: err != nil, %v", test.cmd, err)
		} else if err == nil && test.expectsError {
			t.Errorf("expectation failed for command %s: err == nil", test.cmd)
		}
	}
}
