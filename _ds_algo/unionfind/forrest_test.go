package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionFindForrest(t *testing.T) {
	f := newForest(8)

	assert.Equal(t, 4, f.find(4))

	f.union(3, 4)
	assert.Equal(t, f.find(3), f.find(4))
	assert.NotEqual(t, f.find(2), f.find(4))

	f.union(1, 2)
	assert.Equal(t, f.find(1), f.find(2))
	assert.NotEqual(t, f.find(2), f.find(4))

	f.union(1, 3)
	assert.Equal(t, f.find(1), f.find(3))
	assert.Equal(t, f.find(2), f.find(4))
	assert.NotEqual(t, f.find(4), f.find(7))

}
