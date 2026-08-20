// Copyright (c) 2026 David Corvaglia
// SPDX-License-Identifier: MIT

package ip

type Protocol uint8

const (
	ICMP   Protocol = 1
	IGMP   Protocol = 2
	IPv4   Protocol = 4
	TCP    Protocol = 6
	UDP    Protocol = 17
	IPv6   Protocol = 41
	ICMPv6 Protocol = 58
)
