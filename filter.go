package sosfilter

import (
	"fmt"
	"strings"
)

var startsWith = func(s string, subintf interface{}, previousValid bool) bool {
	if !previousValid {
		return previousValid
	}
	sub := fmt.Sprintf("%v", subintf)
	if len(s) < len(sub) {
		return false
	}
	return strings.HasPrefix(s, sub)
}

var startsWithIgnoreCase = func(s string, subintf interface{}, previousValid bool) bool {
	if !previousValid {
		return previousValid
	}
	sub := fmt.Sprintf("%v", subintf)
	if len(s) < len(sub) {
		return false
	}
	return strings.HasPrefix(strings.ToLower(s), strings.ToLower(sub))
}

var endsWith = func(s string, subintf interface{}, previousValid bool) bool {
	if !previousValid {
		return previousValid
	}
	sub := fmt.Sprintf("%v", subintf)
	if len(s) < len(sub) {
		return false
	}
	return strings.HasSuffix(s, sub)
}

var endsWithIgnoreCase = func(s string, subintf interface{}, previousValid bool) bool {
	if !previousValid {
		return previousValid
	}
	sub := fmt.Sprintf("%v", subintf)
	if len(s) < len(sub) {
		return false
	}
	return strings.HasSuffix(strings.ToLower(s), strings.ToLower(sub))
}

var contains = func(s string, subintf interface{}, previousValid bool) bool {
	if !previousValid {
		return previousValid
	}
	sub := fmt.Sprintf("%v", subintf)
	if len(s) < len(sub) {
		return false
	}
	return strings.Contains(s, sub)
}

var containsIgnoreCase = func(s string, subintf interface{}, previousValid bool) bool {
	if !previousValid {
		return previousValid
	}
	sub := fmt.Sprintf("%v", subintf)
	if len(s) < len(sub) {
		return false
	}
	return strings.Contains(strings.ToLower(s), strings.ToLower(sub))
}

var hasLengthOf = func(s string, lengthintf interface{}, previousValid bool) bool {
	if !previousValid {
		return previousValid
	}
	length, ok := lengthintf.(int)
	if !ok {
		return false
	}
	return (len(s) == length)
}

var hasMinLengthOf = func(s string, minLengthintf interface{}, previousValid bool) bool {
	if !previousValid {
		return previousValid
	}
	minLength, ok := minLengthintf.(int)
	if !ok {
		return false
	}
	return (len(s) >= minLength)
}

var hasMaxLengthOf = func(s string, maxLengthintf interface{}, previousValid bool) bool {
	if !previousValid {
		return previousValid
	}
	maxLength, ok := maxLengthintf.(int)
	if !ok {
		return false
	}
	return (len(s) <= maxLength)
}

var equals = func(s string, strintf interface{}, previousValid bool) bool {
	if !previousValid {
		return previousValid
	}
	str := fmt.Sprintf("%v", strintf)
	return (s == str)
}

var equalsIgnoreCase = func(s string, strintf interface{}, previousValid bool) bool {
	if !previousValid {
		return previousValid
	}
	str := fmt.Sprintf("%v", strintf)
	return (strings.ToLower(s) == strings.ToLower(str))
}
