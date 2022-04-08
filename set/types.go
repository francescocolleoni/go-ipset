// Package set contains structs, types and functions related to ipset sets.
package set

// SetType defines a set type supported by ipset.
type SetType int

const (
	Bitmap_IP = iota
	Bitmap_IP_MAC
	Bitmap_Port
	Hash_IP
	Hash_MAC
	Hash_IP_MAC
	Hash_Net
	Hash_Net_Net
	Hash_IP_Port
	Hash_Net_PORT
	Hash_IP_Port_IP
	Hash_IP_Port_Net
	Hash_IP_Mark
	Hash_Net_Port_Net
	Hash_Net_IFace
	List_Set
)

// String returns the string representation of a given SetType s.
func (s SetType) String() string {
	switch s {
	// Bitmap.
	case Bitmap_IP:
		return "bitmap:ip"
	case Bitmap_IP_MAC:
		return "bitmap:ip,mac"
	case Bitmap_Port:
		return "bitmap:port"

		// Hash.
	case Hash_IP:
		return "hash:ip"
	case Hash_MAC:
		return "hash:mac"
	case Hash_IP_MAC:
		return "hash:ip,mac"
	case Hash_Net:
		return "hash:net"
	case Hash_Net_Net:
		return "hash:net,net"
	case Hash_IP_Port:
		return "hash:ip,port"
	case Hash_Net_PORT:
		return "hash:net,port"
	case Hash_IP_Port_IP:
		return "hash:ip,port,ip"
	case Hash_IP_Port_Net:
		return "hash:ip,port,net"
	case Hash_IP_Mark:
		return "hash:ip,mark"
	case Hash_Net_Port_Net:
		return "hash:net,port,net"
	case Hash_Net_IFace:
		return "hash:net,iface"

		// List.
	case List_Set:
		return "list:set"

	default:
		return "" // Unsupported type
	}
}
