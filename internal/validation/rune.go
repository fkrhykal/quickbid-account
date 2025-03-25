package validation

var lowercase map[rune]any = map[rune]any{
	'a': nil, 'b': nil, 'c': nil, 'd': nil, 'e': nil, 'f': nil,
	'g': nil, 'h': nil, 'i': nil, 'j': nil, 'k': nil, 'l': nil,
	'm': nil, 'n': nil, 'o': nil, 'p': nil, 'q': nil, 'r': nil,
	's': nil, 't': nil, 'u': nil, 'v': nil, 'w': nil, 'x': nil,
	'y': nil, 'z': nil,
}

var uppercase map[rune]any = map[rune]any{
	'A': nil, 'B': nil, 'C': nil, 'D': nil, 'E': nil, 'F': nil,
	'G': nil, 'H': nil, 'I': nil, 'J': nil, 'K': nil, 'L': nil,
	'M': nil, 'N': nil, 'O': nil, 'P': nil, 'Q': nil, 'R': nil,
	'S': nil, 'T': nil, 'U': nil, 'V': nil, 'W': nil, 'X': nil,
	'Y': nil, 'Z': nil,
}

var digit map[rune]any = map[rune]any{
	'0': nil, '1': nil, '2': nil, '3': nil, '4': nil, '5': nil,
	'6': nil, '7': nil, '8': nil, '9': nil,
}

var specialChar map[rune]any = map[rune]any{
	'!': nil, '@': nil, '#': nil, '$': nil, '%': nil, '^': nil,
	'&': nil, '*': nil, '(': nil, ')': nil, '-': nil, '_': nil,
	'=': nil, '+': nil, '[': nil, ']': nil, '{': nil, '}': nil,
	'|': nil, '\\': nil, ':': nil, ';': nil, '"': nil, '\'': nil,
	'<': nil, '>': nil, ',': nil, '.': nil, '?': nil, '/': nil,
	'~': nil, '`': nil,
}

func IsLowercase(r rune) bool {
	_, ok := lowercase[r]
	return ok
}

func IsUppercase(r rune) bool {
	_, ok := uppercase[r]
	return ok
}

func IsDigit(r rune) bool {
	_, ok := digit[r]
	return ok
}

func IsSpecialChar(r rune) bool {
	_, ok := specialChar[r]
	return ok
}

func IsAlphanumeric(r rune) bool {
	return IsUppercase(r) || IsLowercase(r) || IsDigit(r)
}
