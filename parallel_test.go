package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestParallel(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())
	type test struct {
		value      int
		collection []int
	}
	gen := func() test {
		value := rand.Intn(10)
		count := rand.Intn(10) + 1
		ints := make([]int, count)
		for i := 0; i < count; i++ {
			ints[i] = rand.Intn(10)
		}
		sort.Ints(ints)
		return test{
			value:      value,
			collection: ints,
		}
	}
	contains := func(x int, xs []int) bool {
		for _, v := range xs {
			if v == x {
				return true
			}
		}
		return false
	}

	tests := []test{
		gen(),
		gen(),
		gen(),
		gen(),
		gen(),
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("%d %d in %v", i, ts.value, ts.collection), func(t *testing.T) {
			if !contains(ts.value, ts.collection) {
				t.Errorf("%d not int %v", ts.value, ts.collection)
			}
		})
	}
}
