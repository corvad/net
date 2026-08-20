// Copyright (c) 2026 David Corvaglia
// SPDX-License-Identifier: MIT

package ipv4

import "github.com/corvad/net/ip"

type Header []byte

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

func (h Header) Flags() ip.Flags {
	return ip.Flags(h[6] >> 5)
}

func (h Header) FragmentOffset() uint16 {
	return (uint16(h[6])&0b00011111)<<8 | uint16(h[7])
}

func (h Header) TTL() uint8 {
	return h[8]
}

func (h Header) Protocol() ip.Protocol {
	return ip.Protocol(h[9])
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

func (h Header) SetVersion(version uint8) {
	h[0] = h[0]&0b00001111 | version<<4
}

func (h Header) SetIHL(ihl uint8) {
	h[0] = h[0]&0b11110000 | ihl&0b1111
}

func (h Header) SetDSCP(dscp uint8) {
	h[1] = h[1]&0b00000011 | dscp<<2
}

func (h Header) SetECN(ecn uint8) {
	h[1] = h[1]&0b11111100 | ecn&0b11
}

func (h Header) SetTotalLen(totalLen uint16) {
	h[2] = uint8(totalLen >> 8)
	h[3] = uint8(totalLen)
}

func (h Header) SetID(id uint16) {
	h[4] = uint8(id >> 8)
	h[5] = uint8(id)
}

func (h Header) SetFlags(flags ip.Flags) {
	h[6] = h[6]&0b00011111 | uint8(flags)<<5
}

func (h Header) SetFragmentOffset(fragmentOffset uint16) {
	h[6] = h[6]&0b11100000 | uint8(fragmentOffset>>8)&0b00011111
	h[7] = uint8(fragmentOffset)
}

func (h Header) SetTTL(ttl uint8) {
	h[8] = ttl
}

func (h Header) SetProtocol(protocol ip.Protocol) {
	h[9] = uint8(protocol)
}

func (h Header) SetChecksum(checksum uint16) {
	h[10] = uint8(checksum >> 8)
	h[11] = uint8(checksum)
}

func (h Header) SetSrcIP(ip [4]byte) {
	*(*[4]byte)(h[12:16]) = ip
}

func (h Header) SetDstIP(ip [4]byte) {
	*(*[4]byte)(h[16:20]) = ip
}
