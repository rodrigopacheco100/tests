package main

import (
	"reflect"
	"testing"
)

type maniaTestCase struct {
	name     string
	grid     [][]int
	cost     int
	result   [][]int
}

func TestManiaDeImpar(t *testing.T) {
	tests := []maniaTestCase{
		{
			name: "example 1 (3x3)",
			grid: [][]int{
				{1, 2, 1},
				{2, 2, 2},
				{1, 2, 1},
			},
			cost: 1,
			result: [][]int{
				{1, 2, 1},
				{2, 3, 2},
				{1, 2, 1},
			},
		},
		{
			name: "example 2 (5x5)",
			grid: [][]int{
				{8, 7, 2, 5, 7},
				{9, 9, 9, 8, 7},
				{2, 7, 4, 5, 6},
				{6, 2, 8, 2, 1},
				{2, 3, 4, 7, 8},
			},
			cost: 4,
			result: [][]int{
				{8, 7, 2, 5, 8},
				{9, 10, 9, 8, 7},
				{2, 7, 4, 5, 6},
				{7, 2, 9, 2, 1},
				{2, 3, 4, 7, 8},
			},
		},
		{
			name: "example 3 (1x5 already organized)",
			grid: [][]int{
				{1, 2, 3, 4, 5},
			},
			cost: 0,
			result: [][]int{
				{1, 2, 3, 4, 5},
			},
		},
		{
			name: "single cell (1x1)",
			grid: [][]int{
				{7},
			},
			cost: 0,
			result: [][]int{
				{7},
			},
		},
		{
			name: "flip middle cell (1x3)",
			grid: [][]int{
				{2, 2, 4},
			},
			cost: 1,
			result: [][]int{
				{2, 3, 4},
			},
		},
		{
			name: "single odd column (3x1)",
			grid: [][]int{
				{1},
				{3},
				{5},
			},
			cost: 1,
			result: [][]int{
				{1},
				{4},
				{5},
			},
		},
		{
			name: "all evens (3x3)",
			grid: [][]int{
				{2, 4, 6},
				{8, 10, 12},
				{14, 16, 18},
			},
			cost: 4,
			result: [][]int{
				{2, 5, 6},
				{9, 10, 13},
				{14, 17, 18},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cost, result := Solve(tc.grid)

			if cost != tc.cost {
				t.Errorf("cost = %d, expected %d", cost, tc.cost)
			}

			if !reflect.DeepEqual(result, tc.result) {
				t.Errorf("result = %v, expected %v", result, tc.result)
			}
		})
	}
}
