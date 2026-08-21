// Copyright (c) 2026 David Corvaglia
// SPDX-License-Identifier: MIT

package ipv4

import "errors"

type ParseError error

var (
	ErrShort          ParseError = errors.New("ipv4: header len too short")
	ErrInvalidLen     ParseError = errors.New("ipv4: invalid header len")
	ErrInvalidVersion ParseError = errors.New("ipv4: header version incorrect")
	ErrInvalidIHL     ParseError = errors.New("ipv4: header contains invalid IHL")
)

const (
	MinHeaderLen = 20
	MaxHeaderLen = 60
	Version      = 4
)

func Parse(b []byte) (Header, error) {
	if len(b) < MinHeaderLen {
		return nil, ErrShort
	}

	h := Header(b)

	if h.Version() != Version {
		return nil, ErrInvalidVersion
	}

	hLen := int(h.IHL() * 4)
	if hLen < MinHeaderLen {
		return nil, ErrInvalidIHL
	}
	if hLen > len(h) {
		return nil, ErrInvalidIHL
	}

	if int(h.TotalLen()) < hLen {
		return nil, ErrInvalidLen
	}

	return h, nil
}
