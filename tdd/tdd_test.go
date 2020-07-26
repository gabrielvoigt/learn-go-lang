package tdd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiples(t *testing.T)  {
	assert.Equal(t, []int{3}, Multiples(1,5))
	assert.Equal(t, []int{5,6,9,10}, Multiples(4,11))
	assert.Equal(t, []int{}, Multiples(1,3))
}