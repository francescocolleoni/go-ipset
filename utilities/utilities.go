// Package utilities includes all utility functions like
// "get ipset version" or "check if ipset is available".
package utilities

import (
	"os/exec"

	liberrors "github.com/francescocolleoni/go-ipset/errors"
)

// Version returns a string describing the current ipset version or an error if the tool is not accessible.
func Version() (string, error) {
	out, err := RunIPSet("-v")
	if err != nil {
		return "", liberrors.ErrIPSetDidFail
	}

	versionString := string(out)
	if versionString == "" {
		return "", liberrors.ErrIPSetVersionIsNil
	} else {
		return versionString, nil
	}
}

// IPSetIsAvailable returns true if ipset is available.
func IPSetIsAvailable() bool {
	if _, err := Version(); err != nil {
		return false
	} else {
		return true
	}
}

// RunIPSet runs ipset command followed by a list of arguments.
func RunIPSet(args ...string) ([]byte, error) {
	return runCommand("ipset", args...)
}

// runCommand runs a generic command followed by a list of arguments.
func runCommand(name string, args ...string) ([]byte, error) {
	return exec.Command(name, args...).Output()
}
