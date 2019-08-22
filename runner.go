package sosfilter

import (
	"errors"
	"reflect"
)

//NewFilter create a new filter based using slice of string and filter rules
func NewFilter(SliceOfString []string, FilterRules []FilterRule) Filter {
	return Filter{
		SliceOfString: SliceOfString,
		FilterRules:   FilterRules,
	}
}

// Run the SoS filtering
func (f *Filter) Run() ([]string, error) {
	filtered := make([]string, 0)
	err := validate(*f)
	if err != nil {
		return filtered, err
	}
	filterFunctionMap := make(map[FilterFuncID]filterFunc)
	for _, fr := range f.FilterRules {
		switch fr.FilterFunc {
		case StartsWith:
			filterFunctionMap[StartsWith] = startsWith
		case StartsWithNoCase:
			filterFunctionMap[StartsWithNoCase] = startsWithNoCase
		case EndsWith:
			filterFunctionMap[EndsWith] = endsWith
		case EndsWithNoCase:
			filterFunctionMap[EndsWithNoCase] = endsWithNoCase
		case Contains:
			filterFunctionMap[Contains] = contains
		case ContainsNoCase:
			filterFunctionMap[ContainsNoCase] = containsNoCase
		case HasLengthOf:
			filterFunctionMap[HasLengthOf] = hasLengthOf
		case EqualCase:
			filterFunctionMap[EqualCase] = equalCase
		case EqualNoCase:
			filterFunctionMap[EqualNoCase] = equalNoCase
		default:
			continue
		}
	}
	for _, str := range f.SliceOfString {
		previousValid := true
		for _, rules := range f.FilterRules {
			singleFilterFunc, ok := filterFunctionMap[rules.FilterFunc]
			if !ok {
				continue
			}
			previousValid = singleFilterFunc(str, rules.Param, previousValid)
		}
		if previousValid {
			filtered = append(filtered, str)
		}
	}
	return filtered, nil
}

func validate(filter Filter) error {
	if filter.SliceOfString == nil {
		return errors.New("Slice of String must be defined")
	}
	if len(filter.SliceOfString) == 0 {
		return errors.New("Empty slice of string, nothing to be filtered")
	}
	for _, fr := range filter.FilterRules {
		err := validateFilterRule(fr)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateFilterRule(fr FilterRule) error {
	if fr.Param == nil {
		return errors.New("Filter Param must be defined")
	}
	paramType := reflect.TypeOf(fr.Param)
	switch fr.FilterFunc {
	case StartsWith, StartsWithNoCase, EndsWith, EndsWithNoCase, Contains, ContainsNoCase, EqualCase, EqualNoCase:
		if paramType.Kind() != reflect.String {
			return errors.New("Invalid param type")
		}
	case HasLengthOf:
		if paramType.Kind() != reflect.Int {
			return errors.New("Invalid param type")
		}
	default:
		return errors.New("Invalid Filter Func")
	}
	return nil
}
