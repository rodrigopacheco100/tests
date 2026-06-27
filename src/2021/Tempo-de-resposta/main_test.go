package main_test

import (
	tempoderesposta "obi-resolutions/src/2021/Tempo-de-resposta"
	"reflect"
	"testing"
)

type tempoTestCase struct {
	events   []tempoderesposta.Event
	expected map[int]int
}

func TestCalculateResponseTimes(t *testing.T) {
	testsSuite := []tempoTestCase{
		{
			events: []tempoderesposta.Event{
				{Type: "R", Value: 2},
				{Type: "T", Value: 3},
				{Type: "E", Value: 2},
			},
			expected: map[int]int{2: 3},
		},
		{
			events: []tempoderesposta.Event{
				{Type: "R", Value: 2},
				{Type: "T", Value: 3},
				{Type: "E", Value: 2},
				{Type: "R", Value: 1},
				{Type: "T", Value: 2},
				{Type: "E", Value: 1},
				{Type: "R", Value: 3},
			},
			expected: map[int]int{1: 2, 2: 3, 3: -1},
		},
		{
			events: []tempoderesposta.Event{
				{Type: "R", Value: 1},
				{Type: "E", Value: 1},
			},
			expected: map[int]int{1: 1},
		},
	}

	for _, testCase := range testsSuite {
		if result := tempoderesposta.CalculateResponseTimes(testCase.events); !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("CalculateResponseTimes(%v) = %v, expected %v", testCase.events, result, testCase.expected)
		}
	}
}
