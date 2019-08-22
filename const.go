package sosfilter

//FilterFuncType is a defined filtering func type
type FilterFuncID int

const (
	StartsWith FilterFuncID = iota
	EndsWith
	Contains
	HasLengthOf
	EqualCase
	EqualNoCase
)
