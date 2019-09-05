package sosfilter

import (
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
