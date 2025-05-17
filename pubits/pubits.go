package pubits

// Sets the bit at pos in the integer n.
func SetBit(n uint8, pos uint8) uint8 {
	n |= (1 << pos)
	return n
}

// Clears the bit at pos in n.
func ClearBit(n uint8, pos uint8) uint8 {
	mask := ^(1 << pos)
	return n & uint8(mask)
}

// Checks if bit at pos is set
func HasBit(n uint8, pos uint8) bool {
	val := n & (1 << pos)
	return (val > 0)
}

// Toggles the bit at pos in n.
func ToggleBit(n uint8, pos uint8) uint8 {
	return n ^ (1 << pos)
}
