package sosfilter

//FilterFuncID is a defined filtering func type
type FilterFuncID int

const (
	//StartsWith is a filter func ID for checking a string starts with a substring with case-sensitive comparison
	StartsWith FilterFuncID = iota
	//StartsWithNoCase is a filter func ID for checking a string starts with a substring with ignoring case-sensitive comparison
	StartsWithNoCase
	//EndsWith is a filter func ID for checking a string ends with a substring with case-sensitive comparison
	EndsWith
	//EndsWithNoCase is a filter func ID for checking a string ends with a substring with ignoring case-sensitive comparison
	EndsWithNoCase
	//Contains is a filter func ID for checking a string contains a substring with case-sensitive comparison
	Contains
	//ContainsNoCase is a filter func ID for checking a string contains a substring with ignoring case-sensitive comparison
	ContainsNoCase
	//HasLengthOf is a filter func ID for checking a string has length of n characters
	HasLengthOf
	//EqualCase is a filter func ID for checking a string is equal to another string with case-sensitive comparison
	EqualCase
	//EqualNoCase is a filter func ID for checking a string is equal to another string with ignoring case-sensitive comparison
	EqualNoCase
)
