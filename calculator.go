package smscredit

import (
	"fmt"
	"unicode/utf8"
)

type CharacterType int

const (
	CharacterType_Normal CharacterType = iota
	CharacterType_Turkish
	CharacterType_Unicode
)

func (ct CharacterType) String() string {
	switch ct {
	case CharacterType_Normal:
		return "normal"
	case CharacterType_Turkish:
		return "turkish"
	case CharacterType_Unicode:
		return "unicode"
	default:
		return "normal"
	}
}

type Options struct {
	CharType   CharacterType
	IsWhatsApp bool

	AddCancelLink bool
	AddReadLink   bool
}

type Calculator struct {
	cfg Config
}

func New(cfg Config) *Calculator {
	if cfg.BaseOverhead == 0 {
		cfg.BaseOverhead = 5
	}
	return &Calculator{cfg: cfg}
}

func (c *Calculator) Credits(message string, opts Options) (int, error) {
	if opts.IsWhatsApp {
		return 1, nil
	}

	length := c.cfg.BaseOverhead + c.messageLength(message, opts.CharType)

	if opts.AddCancelLink {
		length += c.cfg.CancelLinkLen
	}
	if opts.AddReadLink {
		length += c.cfg.ReadLinkLen
	}

	return creditsByLength(length, opts.CharType)
}

func (c *Calculator) messageLength(message string, ct CharacterType) int {
	switch ct {
	case CharacterType_Unicode:
		return utf8.RuneCountInString(message)

	case CharacterType_Turkish:
		return countWithExtended(message, turkishExtended)

	case CharacterType_Normal:
		return countWithExtended(message, gsmExtended)

	default:
		return countWithExtended(message, gsmExtended)
	}
}

func creditsByLength(length int, ct CharacterType) (int, error) {
	switch ct {
	case CharacterType_Turkish:
		return segmentCredits(length, turkishSegments, 882, "mesaj uzunluğu 882 karakterden fazla olmamalı")
	case CharacterType_Unicode:
		return segmentCredits(length, unicodeSegments, 402, "mesaj uzunluğu 402 karakterden fazla olmamalı")
	case CharacterType_Normal:
		return segmentCredits(length, normalSegments, 918, "mesaj uzunluğu 918 karakterden fazla olmamalı")
	default:
		return 0, fmt.Errorf("geçersiz sms tipi: %d", int(ct))
	}
}

type segmentRule struct {
	MaxLen  int
	Credits int
}

func segmentCredits(length int, segments []segmentRule, _ int, maxMsg string) (int, error) {
	for _, s := range segments {
		if length <= s.MaxLen {
			return s.Credits, nil
		}
	}
	return 0, fmt.Errorf("%s (şu an: %d)", maxMsg, length)
}
