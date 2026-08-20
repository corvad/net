// Copyright (c) 2026 David Corvaglia
// SPDX-License-Identifier: MIT

package ip

type Flags uint8

const (
	MoreFragments Flags = 1 << iota
	DontFragment
	Reserved
)
