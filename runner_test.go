package sosfilter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFilter(t *testing.T) {
	sos := []string{"Avenged Death", "Avenged Theater", "Avenged Corpse", "Avenged Theater", "Napalm Sevenfold", "Napalm Theater", "Napalm Corpse", "Dream Corpse", "Dream Sevenfold", "Dream Death", "Cannibal Theater", "Cannibal Death", "Cannibal Sevenfold"}
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
