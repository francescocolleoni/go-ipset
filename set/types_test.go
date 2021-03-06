package set

import (
	"testing"
)

func TestSetType_String(t *testing.T) {
	tests := map[SetType]string{
		SetTypeBitmapIP:       "bitmap:ip",
		SetTypeBitmapIPMAC:    "bitmap:ip,mac",
		SetTypeBitmapPort:     "bitmap:port",
		SetTypeHashIP:         "hash:ip",
		SetTypeHashIPMAC:      "hash:ip,mac",
		SetTypeHashIPPort:     "hash:ip,port",
		SetTypeHashIPPortIP:   "hash:ip,port,ip",
		SetTypeHashIPPortNet:  "hash:ip,port,net",
		SetTypeHashIPMark:     "hash:ip,mark",
		SetTypeHashMAC:        "hash:mac",
		SetTypeHashNet:        "hash:net",
		SetTypeHashNetNet:     "hash:net,net",
		SetTypeHashNetPort:    "hash:net,port",
		SetTypeHashNetPortNet: "hash:net,port,net",
		SetTypeHashNetIFace:   "hash:net,iface",
		SetTypeListSet:        "list:set",
	}

	// Populate tests with invalid values.
	// Any value lower or equal to 0 or higher or equal to 16 must result in "".
	tests[-1] = "" // SetType is an enum whose first value is iota (equal to 0)
	tests[16] = "" // and last value is iota + 15 (the enum contains 16 values).

	for key, expectation := range tests {
		if result := key.String(); result != expectation {
			t.Errorf("expectation failed for key %v: \"%s\" != \"%s\" (expected)", key, result, expectation)
		}
	}
}
func TestSetType_InitWithString(t *testing.T) {
	tests := map[string]SetType{
		"bitmap:ip":         SetTypeBitmapIP,
		"bitmap:ip,mac":     SetTypeBitmapIPMAC,
		"bitmap:port":       SetTypeBitmapPort,
		"hash:ip":           SetTypeHashIP,
		"hash:ip,mac":       SetTypeHashIPMAC,
		"hash:ip,port":      SetTypeHashIPPort,
		"hash:ip,port,ip":   SetTypeHashIPPortIP,
		"hash:ip,port,net":  SetTypeHashIPPortNet,
		"hash:ip,mark":      SetTypeHashIPMark,
		"hash:mac":          SetTypeHashMAC,
		"hash:net":          SetTypeHashNet,
		"hash:net,net":      SetTypeHashNetNet,
		"hash:net,port":     SetTypeHashNetPort,
		"hash:net,port,net": SetTypeHashNetPortNet,
		"hash:net,iface":    SetTypeHashNetIFace,
		"list:set":          SetTypeListSet,
		"":                  SetTypeUnsupported,
	}

	for key, expectation := range tests {
		if result := SetTypeWithString(key); result != expectation {
			t.Errorf("expectation failed for key %v: \"%s\" != \"%s\" (expected)", key, result, expectation)
		}
	}
}
