package sosfilter

//FilterFuncID is a defined filtering func type
type FilterFuncID int

const (
	//StartsWith is a filter func ID for checking a string starts with a substring with case-sensitive comparison
	StartsWith FilterFuncID = iota
	//StartsWithIgnoreCase is a filter func ID for checking a string starts with a substring with ignoring case-sensitive comparison
	StartsWithIgnoreCase
	//EndsWith is a filter func ID for checking a string ends with a substring with case-sensitive comparison
	EndsWith
	//EndsWithIgnoreCase is a filter func ID for checking a string ends with a substring with ignoring case-sensitive comparison
	EndsWithIgnoreCase
	//Contains is a filter func ID for checking a string contains a substring with case-sensitive comparison
	Contains
	//ContainsIgnoreCase is a filter func ID for checking a string contains a substring with ignoring case-sensitive comparison
	ContainsIgnoreCase
	//HasLengthOf is a filter func ID for checking a string has length of n characters
	HasLengthOf
	// HasMinLengthOf is a filter func ID for checking a string has minimum length of n characters
	HasMinLengthOf
	// HasMaxLengthOf is a filter func ID for checking a string has maximum length of n characters
	HasMaxLengthOf
	//Equals is a filter func ID for checking a string is equal to another string with case-sensitive comparison
	Equals
	//EqualsIgnoreCase is a filter func ID for checking a string is equal to another string with ignoring case-sensitive comparison
	EqualsIgnoreCase
)
