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

var equalCase = func(s string, strintf interface{}, previousValid bool) bool {
	if !previousValid {
		return previousValid
	}
	str := fmt.Sprintf("%v", strintf)
	return (s == str)
}

var equalNoCase = func(s string, strintf interface{}, previousValid bool) bool {
	if !previousValid {
		return previousValid
	}
	str := fmt.Sprintf("%v", strintf)
	return (strings.ToLower(s) == strings.ToLower(str))
}
