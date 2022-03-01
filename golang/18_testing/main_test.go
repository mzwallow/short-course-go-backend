package main

import "testing"

func TestAdd(t *testing.T) {
	data := []struct {
		a, b, res int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{3414, 32, 3446},
		{-1, 1, 0},
	}

	for _, v := range data {
		if got := Add(v.a, v.b); got != v.res {
			t.Errorf("Add(1, 2) == %d, expected %d", got, v.res)
		}
	}
}

//t.Errorf("Add(1, 2) == %d, expected %d", got, expected)
