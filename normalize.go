package smscredit

import "strings"

func NormalizeMessage(message string, ct CharacterType) string {
	if ct != CharacterType_Normal {
		return message
	}

	replacements := map[string]string{
		"ç": "c", "Ç": "C",
		"ğ": "g", "Ğ": "G",
		"ı": "i", "İ": "I",
		"ö": "o", "Ö": "O",
		"ş": "s", "Ş": "S",
		"ü": "u", "Ü": "U",
	}

	for tr, en := range replacements {
		message = strings.ReplaceAll(message, tr, en)
	}

	return message
}
