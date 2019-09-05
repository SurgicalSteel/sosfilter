package sosfilter

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFilter(t *testing.T) {
	sos := []string{"Avenged Death", "Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Napalm Corpse", "Dream Corpse", "Dream Sevenfold", "Dream Death", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: ContainsIgnoreCase,
			Param:      "th",
		},
		FilterRule{
			FilterFunc: HasMinLengthOf,
			Param:      "11",
		},
	}

	expected := Filter{
		SliceOfString: sos,
		FilterRules:   filterRules,
	}

	actual := NewFilter(sos, filterRules)
	assert.Equal(t, expected, actual)
}

func TestFilterStartsWithIgnoreCase(t *testing.T) {
	sos := []string{"Avenged Death", "Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Napalm Corpse", "Dream Corpse", "Dream Sevenfold", "Dream Death", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: StartsWithIgnoreCase,
			Param:      "n",
		},
	}

	expected := []string{"Napalm Sevenfold", "Napalm Theater", "Napalm Corpse"}
	filter := NewFilter(sos, filterRules)
	actual, _ := filter.Run()
	assert.Equal(t, expected, actual)
}

func TestFilterStartsWith(t *testing.T) {
	sos := []string{"Avenged Death", "Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Napalm Corpse", "Dream Corpse", "Dream Sevenfold", "Dream Death", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: StartsWith,
			Param:      "Av",
		},
	}

	expected := []string{"Avenged Death", "Avenged Corpse", "Avenged Theater"}
	filter := NewFilter(sos, filterRules)
	actual, _ := filter.Run()
	assert.Equal(t, expected, actual)
}

func TestFilterContainsIgnoreCase(t *testing.T) {
	sos := []string{"Avenged Death", "Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Napalm Corpse", "Dream Corpse", "Dream Sevenfold", "Dream Death", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: ContainsIgnoreCase,
			Param:      "th",
		},
	}

	expected := []string{"Avenged Death", "Avenged Theater", "Napalm Theater", "Dream Death", "Cannibal Theater", "Cannibal Death"}
	filter := NewFilter(sos, filterRules)
	actual, _ := filter.Run()
	assert.Equal(t, expected, actual)
}

func TestFilterContains(t *testing.T) {
	sos := []string{"Avenged Death", "Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Napalm Corpse", "Dream Corpse", "Dream Sevenfold", "Dream Death", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: Contains,
			Param:      "Th",
		},
	}

	expected := []string{"Avenged Theater", "Napalm Theater", "Cannibal Theater"}
	filter := NewFilter(sos, filterRules)
	actual, _ := filter.Run()
	assert.Equal(t, expected, actual)
}

func TestFilterEndsWithIgnoreCase(t *testing.T) {
	sos := []string{"Avenged Death", "Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Napalm Corpse", "Dream Corpse", "Dream Sevenfold", "Dream Death", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: EndsWithIgnoreCase,
			Param:      "death",
		},
	}

	expected := []string{"Avenged Death", "Dream Death", "Cannibal Death"}
	filter := NewFilter(sos, filterRules)
	actual, _ := filter.Run()
	assert.Equal(t, expected, actual)
}

func TestFilterEndsWith(t *testing.T) {
	sos := []string{"Avenged Death", "Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Napalm Corpse", "Dream Corpse", "Dream Sevenfold", "Dream Death", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: EndsWith,
			Param:      "Death",
		},
	}

	expected := []string{"Avenged Death", "Dream Death", "Cannibal Death"}
	filter := NewFilter(sos, filterRules)
	actual, _ := filter.Run()
	assert.Equal(t, expected, actual)
}

func TestFilterEqualsIgnoreCase(t *testing.T) {
	sos := []string{"Avenged Death", "Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Napalm Corpse", "Dream Corpse", "Dream Sevenfold", "Dream Death", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: EqualsIgnoreCase,
			Param:      "aVeNgEd DeAtH",
		},
	}

	expected := []string{"Avenged Death"}
	filter := NewFilter(sos, filterRules)
	actual, _ := filter.Run()
	assert.Equal(t, expected, actual)
}

func TestFilterEquals(t *testing.T) {
	sos := []string{"Avenged Death", "Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Napalm Corpse", "Dream Corpse", "Dream Sevenfold", "Dream Death", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: Equals,
			Param:      "Napalm Corpse",
		},
	}

	expected := []string{"Napalm Corpse"}
	filter := NewFilter(sos, filterRules)
	actual, _ := filter.Run()
	assert.Equal(t, expected, actual)
}

func TestFilterHasLengthOf(t *testing.T) {
	sos := []string{"Avenged Death", "Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Napalm Corpse", "Dream Corpse", "Dream Sevenfold", "Dream Death", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: HasLengthOf,
			Param:      13,
		},
	}

	expected := []string{"Avenged Death", "Napalm Corpse"}
	filter := NewFilter(sos, filterRules)
	actual, _ := filter.Run()
	assert.Equal(t, expected, actual)
}

func TestFilterHasMinLengthOf(t *testing.T) {
	sos := []string{"Avenged Death", "Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Napalm Corpse", "Dream Corpse", "Dream Sevenfold", "Dream Death", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: HasMinLengthOf,
			Param:      14,
		},
	}

	expected := []string{"Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Dream Sevenfold", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
	filter := NewFilter(sos, filterRules)
	actual, _ := filter.Run()
	assert.Equal(t, expected, actual)
}

func TestFilterHasMaxLengthOf(t *testing.T) {
	sos := []string{"Avenged Death", "Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Napalm Corpse", "Dream Corpse", "Dream Sevenfold", "Dream Death", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: HasMaxLengthOf,
			Param:      12,
		},
	}

	expected := []string{"Dream Corpse", "Dream Death"}
	filter := NewFilter(sos, filterRules)
	actual, _ := filter.Run()
	assert.Equal(t, expected, actual)
}

func TestInvalidNoParamFilterRule(t *testing.T) {
	sos := []string{"Avenged Death", "Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Napalm Corpse", "Dream Corpse", "Dream Sevenfold", "Dream Death", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: HasMaxLengthOf,
			Param:      nil,
		},
	}

	expectedError := errors.New("Filter Param must be defined")
	filter := NewFilter(sos, filterRules)
	_, actualError := filter.Run()
	assert.Equal(t, expectedError, actualError)
}

func TestInvalidParamTypeFilterRuleExpectInteger(t *testing.T) {
	sos := []string{"Avenged Death", "Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Napalm Corpse", "Dream Corpse", "Dream Sevenfold", "Dream Death", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: HasMaxLengthOf,
			Param:      "avenged death",
		},
	}

	expectedError := errors.New("Invalid param type")
	filter := NewFilter(sos, filterRules)
	_, actualError := filter.Run()
	assert.Equal(t, expectedError, actualError)
}

func TestInvalidParamTypeFilterRuleExpectString(t *testing.T) {
	sos := []string{"Avenged Death", "Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Napalm Corpse", "Dream Corpse", "Dream Sevenfold", "Dream Death", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: EqualsIgnoreCase,
			Param:      666,
		},
	}

	expectedError := errors.New("Invalid param type")
	filter := NewFilter(sos, filterRules)
	_, actualError := filter.Run()
	assert.Equal(t, expectedError, actualError)
}

func TestInvalidNilSliceOfString(t *testing.T) {
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: EqualsIgnoreCase,
			Param:      666,
		},
	}

	expectedError := errors.New("Slice of String must be defined")
	filter := NewFilter(nil, filterRules)
	_, actualError := filter.Run()
	assert.Equal(t, expectedError, actualError)
}

func TestInvalidEmptySliceOfString(t *testing.T) {
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: EqualsIgnoreCase,
			Param:      666,
		},
	}

	expectedError := errors.New("Empty slice of string, nothing to be filtered")
	filter := NewFilter([]string{}, filterRules)
	_, actualError := filter.Run()
	assert.Equal(t, expectedError, actualError)
}

func TestInvalidNilParam(t *testing.T) {
	sos := []string{"Avenged Death", "Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Napalm Corpse", "Dream Corpse", "Dream Sevenfold", "Dream Death", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
	filterRules := []FilterRule{
		FilterRule{
			FilterFunc: EqualsIgnoreCase,
			Param:      nil,
		},
	}

	expectedError := errors.New("Filter Param must be defined")
	filter := NewFilter(sos, filterRules)
	_, actualError := filter.Run()
	assert.Equal(t, expectedError, actualError)
}
