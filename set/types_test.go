package set

import (
	"testing"
)

func TestSetType_String(t *testing.T) {
	tests := map[SetType]string{
		Bitmap_IP:         "bitmap:ip",
		Bitmap_IP_MAC:     "bitmap:ip,mac",
		Bitmap_Port:       "bitmap:port",
		Hash_IP:           "hash:ip",
		Hash_MAC:          "hash:mac",
		Hash_IP_MAC:       "hash:ip,mac",
		Hash_Net:          "hash:net",
		Hash_Net_Net:      "hash:net,net",
		Hash_IP_Port:      "hash:ip,port",
		Hash_Net_PORT:     "hash:net,port",
		Hash_IP_Port_IP:   "hash:ip,port,ip",
		Hash_IP_Port_Net:  "hash:ip,port,net",
		Hash_IP_Mark:      "hash:ip,mark",
		Hash_Net_Port_Net: "hash:net,port,net",
		Hash_Net_IFace:    "hash:net,iface",
		List_Set:          "list:set",
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
