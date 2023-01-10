package array_test

import (
	"peanut/pkg/array"
	"testing"
)

func TestIsInWithString(t *testing.T) {
	isInTests := []struct {
		description string
		element     string
		array       []string
		expect      bool
	}{
		{
			description: "String value with true",
			element:     "dog",
			array:       []string{"dog", "cat", "monkey", "pig", "chicken"},
			expect:      true,
		},
		{
			description: "String value with false",
			element:     "apple",
			array:       []string{"dog", "cat", "monkey", "pig", "chicken"},
			expect:      false,
		},
		{
			description: "String value with empty array",
			element:     "apple",
			array:       []string{},
			expect:      false,
		},
		{
			description: "String value with empty element",
			element:     "",
			array:       []string{"dog", "cat", "monkey", "pig", "chicken"},
			expect:      false,
		},
	}

	for _, tc := range isInTests {
		if result := array.IsIn(tc.element, tc.array); result != tc.expect {
			t.Errorf("Failed test case: %s", tc.description)
		}
	}
}

func TestIsInWithInt(t *testing.T) {
	isInTests := []struct {
		description string
		element     int
		array       []int
		expect      bool
	}{
		{
			description: "Int value with true",
			element:     1,
			array:       []int{1, 2, 3, 4},
			expect:      true,
		},
		{
			description: "Int value with false",
			element:     5,
			array:       []int{1, 2, 3, 4},
			expect:      false,
		},
		{
			description: "Int value with empty array",
			element:     1,
			array:       []int{},
			expect:      false,
		},
	}

	for _, tc := range isInTests {
		if result := array.IsIn(tc.element, tc.array); result != tc.expect {
			t.Errorf("Failed test case: %s", tc.description)
		}
	}
}

func TestIsInWithFloat(t *testing.T) {
	isInTests := []struct {
		description string
		element     float64
		array       []float64
		expect      bool
	}{
		{
			description: "Float value with true",
			element:     0.1,
			array:       []float64{0.1, 2.1, 0.3, 1.41},
			expect:      true,
		},
		{
			description: "Float value with false",
			element:     0.11,
			array:       []float64{0.1, 2.1, 0.3, 1.41},
			expect:      false,
		},
		{
			description: "Float value with empty array",
			element:     0.99,
			array:       []float64{},
			expect:      false,
		},
	}

	for _, tc := range isInTests {
		if result := array.IsIn(tc.element, tc.array); result != tc.expect {
			t.Errorf("Failed test case: %s", tc.description)
		}
	}
}
