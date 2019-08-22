package sosfilter

// Filter is the main struct to be defined before running filtering rules
type Filter struct {
	SliceOfString []string
	FilterRules   []FilterRule
}

//FilterRule is a single rule which can be applied into slice of string filtering
type FilterRule struct {
	FilterFunc FilterFuncID
	Param      interface{}
}

type filterFunc func(string, interface{}, bool) bool
