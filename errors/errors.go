// Package errors defines all errors exported by go-ipset.
package errors

import "errors"

var ErrIPSetDidFail = errors.New("ipset command did fail")
var ErrIPSetVersionIsNil = errors.New("ipset version is nil")
