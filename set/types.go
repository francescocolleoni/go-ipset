// Package set contains structs, types and functions related to ipset sets.
package set

// SetType defines a set type supported by ipset.
type SetType int

const (
	SetTypeBitmapIP = iota
	SetTypeBitmapIPMAC
	SetTypeBitmapPort
	SetTypeHashIP
	SetTypeHashIPMAC
	SetTypeHashIPPort
	SetTypeHashIPPortIP
	SetTypeHashIPPortNet
	SetTypeHashIPMark
	SetTypeHashMAC
	SetTypeHashNet
	SetTypeHashNetNet
	SetTypeHashNetPort
	SetTypeHashNetPortNet
	SetTypeHashNetIFace
	SetTypeListSet
	SetTypeUnsupported
)

// String returns the string representation of a given SetType s.
func (s SetType) String() string {
	switch s {
	// Bitmap.
	case SetTypeBitmapIP:
		return "bitmap:ip"
	case SetTypeBitmapIPMAC:
		return "bitmap:ip,mac"
	case SetTypeBitmapPort:
		return "bitmap:port"

		// Hash.
	case SetTypeHashIP:
		return "hash:ip"
	case SetTypeHashMAC:
		return "hash:mac"
	case SetTypeHashIPMAC:
		return "hash:ip,mac"
	case SetTypeHashNet:
		return "hash:net"
	case SetTypeHashNetNet:
		return "hash:net,net"
	case SetTypeHashIPPort:
		return "hash:ip,port"
	case SetTypeHashNetPort:
		return "hash:net,port"
	case SetTypeHashIPPortIP:
		return "hash:ip,port,ip"
	case SetTypeHashIPPortNet:
		return "hash:ip,port,net"
	case SetTypeHashIPMark:
		return "hash:ip,mark"
	case SetTypeHashNetPortNet:
		return "hash:net,port,net"
	case SetTypeHashNetIFace:
		return "hash:net,iface"

		// List.
	case SetTypeListSet:
		return "list:set"

	default:
		return "" // Unsupported type
	}
}
func SetTypeWithString(in string) SetType {
	switch in {
	// Bitmap.
	case "bitmap:ip":
		return SetTypeBitmapIP
	case "bitmap:ip,mac":
		return SetTypeBitmapIPMAC
	case "bitmap:port":
		return SetTypeBitmapPort

		// Hash.
	case "hash:ip":
		return SetTypeHashIP
	case "hash:mac":
		return SetTypeHashMAC
	case "hash:ip,mac":
		return SetTypeHashIPMAC
	case "hash:net":
		return SetTypeHashNet
	case "hash:net,net":
		return SetTypeHashNetNet
	case "hash:ip,port":
		return SetTypeHashIPPort
	case "hash:net,port":
		return SetTypeHashNetPort
	case "hash:ip,port,ip":
		return SetTypeHashIPPortIP
	case "hash:ip,port,net":
		return SetTypeHashIPPortNet
	case "hash:ip,mark":
		return SetTypeHashIPMark
	case "hash:net,port,net":
		return SetTypeHashNetPortNet
	case "hash:net,iface":
		return SetTypeHashNetIFace

		// List.
	case "list:set":
		return SetTypeListSet

	default:
		return SetTypeUnsupported
	}
}
