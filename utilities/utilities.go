// Package utilities includes all utility functions like
// "get ipset version" or "check if ipset is available".
package utilities

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	liberrors "github.com/francescocolleoni/go-ipset/errors"
)

// IPSetOutput describes the output of the execution of ipset.
type IPSetOutput struct {
	Error error

	In  string
	Out string
}

// Version returns a string describing the current ipset version or an error if the tool is not accessible.
func Version() (string, error) {
	out, err := RunIPSet("-v")
	if err != nil {
		return "", liberrors.ErrIPSetDidFail
	}

	versionString := string(out.Out)
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
func RunIPSet(args ...string) (IPSetOutput, error) {
	if out, err := runCommand("ipset", args...); err != nil {
		return newIPSetErrorOutput(out, err, args...), err
	} else {
		return newIPSetOutput(out, args...), nil
	}
}

// Support functions.

// newIPSetOutput returns an IPSetOutput instance representing a successful run of ipset command.
func newIPSetOutput(out []byte, args ...string) IPSetOutput {
	in := fmt.Sprintf("ipset %v", args)
	in = strings.ReplaceAll(in, "[", "")
	in = strings.ReplaceAll(in, "]", "")

	if out == nil {
		return IPSetOutput{In: in}
	} else {
		return IPSetOutput{In: in, Out: string(out)}
	}
}

// newIPSetErrorOutput returns an IPSetOutput instance representing a run of ipset command that returned an error.
func newIPSetErrorOutput(out []byte, err error, args ...string) IPSetOutput {
	result := newIPSetOutput(out, args...)
	result.Error = rewriteIPSetErrorFromCombinedOutput(out, err)
	return result
}

// rewriteIPSetErrorFromCombinedOutput formats ipset output received along with an error as a human readable error.
func rewriteIPSetErrorFromCombinedOutput(out []byte, err error) error {
	if out == nil {
		return err
	}

	reason := regexp.MustCompile(`^ipset v[0-9]+[\.0-9]*[\.0-9]*:\s`).ReplaceAll(out, []byte{})
	reason = regexp.MustCompile("Try `ipset help' for more information.").ReplaceAll(reason, []byte{})
	return errors.New(fmt.Sprintf(`ipset returned error "%s"`, strings.Trim(string(reason), "\n")))
}

// runCommand runs a generic command followed by a list of arguments.
func runCommand(name string, args ...string) ([]byte, error) {
	return exec.Command(name, args...).CombinedOutput()
}
