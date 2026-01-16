package smscredit

var gsmExtended = map[rune]struct{}{
	'\\': {}, '€': {}, '{': {}, '}': {}, '[': {}, '~': {}, ']': {}, '^': {}, '|': {},
}

var turkishExtended = map[rune]struct{}{
	'\\': {}, '€': {}, '{': {}, '}': {}, '[': {}, '~': {}, ']': {}, '^': {}, '|': {},
	'ş': {}, 'ğ': {}, 'ı': {}, 'Ş': {}, 'İ': {}, 'Ğ': {}, 'ç': {}, 'Ç': {}, 'ü': {}, 'Ü': {}, 'ö': {}, 'Ö': {},
}

func countWithExtended(s string, ext map[rune]struct{}) int {
	n := 0
	for _, r := range s {
		if _, ok := ext[r]; ok {
			n += 2
		} else {
			n++
		}
	}
	return n
}
