package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	equations []string
	expected  bool
}

var data []testData = []testData{
	testData{
		equations: []string{"a==b", "b!=a"},
		expected:  false,
	},
	testData{
		equations: []string{"b==a", "a==b"},
		expected:  true,
	},
	testData{
		equations: []string{"a==b", "b==c", "a==c"},
		expected:  true,
	},
	testData{
		equations: []string{"a==b", "b!=c", "c==a"},
		expected:  false,
	},
	testData{
		equations: []string{"c==c", "b==d", "x!=z"},
		expected:  true,
	},
	testData{
		equations: []string{"c==c"},
		expected:  true,
	},
	testData{
		equations: []string{"c!=c"},
		expected:  false,
	},
	testData{
		equations: []string{"c!=z"},
		expected:  true,
	},
	testData{
		equations: []string{"a==b", "b==c", "c==d"},
		expected:  true,
	},
	testData{
		equations: []string{"a!=b", "b!=c", "c!=d"},
		expected:  true,
	},
	testData{
		equations: []string{"a==b", "b==c", "c==d", "d!=b"},
		expected:  false,
	},
}

func TestSol0(t *testing.T) {
	for _, item := range data {
		assert.Equal(t, item.expected, equationsPossible(item.equations))
	}
}
