package validation

func HasLength(s string, min int, max int) bool {
	return len(s) >= min && len(s) <= max
}
func HasPrefixChar(s string, checkFuncs ...func(char rune) bool) bool {
	firstChar := rune(s[0])
	HasValidPrefix := false
	for _, checkFunc := range checkFuncs {
		HasValidPrefix = HasValidPrefix || checkFunc(firstChar)
	}
	return HasValidPrefix
}

func HasSuffixChar(s string, checkFuncs ...func(char rune) bool) bool {
	lastChar := rune(s[len(s)-1])
	HasValidSuffix := false
	for _, checkFunc := range checkFuncs {
		HasValidSuffix = HasValidSuffix || checkFunc(lastChar)
	}
	return HasValidSuffix
}

func HasAffixChar(s string, checkFuncs ...func(char rune) bool) bool {
	firstChar := rune(s[0])
	lastChar := rune(s[len(s)-1])
	HasValidPrefix := false
	HasValidSuffix := false
	for _, checkFunc := range checkFuncs {
		HasValidPrefix = HasValidPrefix || checkFunc(firstChar)
		HasValidSuffix = HasValidSuffix || checkFunc(lastChar)
	}
	return HasValidPrefix && HasValidSuffix
}
