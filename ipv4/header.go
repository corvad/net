// Copyright (c) 2026 David Corvaglia. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package ipv4

// A simplified representation of an IPv4 header.
type Header []byte

type Flags uint8

const (
	MoreFragments Flags = 1 << iota
	DontFragment
	Reserved
)

func (h Header) Version() uint8 {
	return h[0] >> 4
}

func (h Header) IHL() uint8 {
	return h[0] & 0b1111
}

func (h Header) DSCP() uint8 {
	return h[1] >> 2
}

func (h Header) ECN() uint8 {
	return h[1] & 0b11
}

func (h Header) TotalLen() uint16 {
	return uint16(h[2])<<8 | uint16(h[3])
}

func (h Header) ID() uint16 {
	return uint16(h[4])<<8 | uint16(h[5])
}

func (h Header) Flags() Flags {
	return Flags(h[6] >> 5)
}

func (h Header) FragmentOffset() uint16 {
	return (uint16(h[6])&0b00011111)<<8 | uint16(h[7])
}

func (h Header) TTL() uint8 {
	return h[8]
}

func (h Header) Protocol() uint8 {
	return h[9]
}

func (h Header) Checksum() uint16 {
	return uint16(h[10])<<8 | uint16(h[11])
}

func (h Header) SrcIP() [4]byte {
	return [4]byte(h[12:16])
}

func (h Header) DstIP() [4]byte {
	return [4]byte(h[16:20])
}
