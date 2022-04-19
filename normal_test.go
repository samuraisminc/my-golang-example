package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssertFailure(t *testing.T) {
	tests := []struct {
		exp int
		val int
	}{
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 8},
	}
	for _, ts := range tests {
		assert.Equal(t, ts.exp, ts.val)
	}
}

func TestNormalFailure(t *testing.T) {
	tests := []struct {
		exp int
		val int
	}{
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 8},
	}
	for _, ts := range tests {
		if ts.exp != ts.val {
			t.Errorf("exp = %v, value = %v", ts.exp, ts.val)
		}
	}
}
