package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 3, sum(1, 2), "1 + 2 should be 3")
}
