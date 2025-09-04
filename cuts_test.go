package cuts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFirstWhere(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name          string
		whereFunc     func(val int) bool
		input         []int
		expectedIndex int
		expectedFound bool
	}{
		{
			name: "find_first_elem",
			whereFunc: func(val int) bool {
				return val == 0
			},
			input:         []int{0, 1, 2, 3, 4},
			expectedIndex: 0,
			expectedFound: true,
		},
		{
			name: "find_last_elem",
			whereFunc: func(val int) bool {
				return val == 4
			},
			input:         []int{0, 1, 2, 3, 4},
			expectedIndex: 4,
			expectedFound: true,
		},
		{
			name: "no_matching",
			whereFunc: func(val int) bool {
				return val > 5
			},
			input:         []int{0, 1, 2, 3, 4},
			expectedIndex: -1,
			expectedFound: false,
		},
		{
			name: "multi_matching",
			whereFunc: func(val int) bool {
				return val%2 == 0
			},
			input:         []int{1, 2, 3, 4},
			expectedIndex: 1,
			expectedFound: true,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actualIndex, actualFound := FirstWhere(tc.input, tc.whereFunc)
			assert.Equal(t, tc.expectedIndex, actualIndex)
			if assert.Equal(t, tc.expectedFound, actualFound) && tc.expectedFound {
				assert.Equal(t, tc.input[tc.expectedIndex], tc.input[actualIndex])
			}
		})
	}
}

func TestLastWhere(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name          string
		whereFunc     func(val int) bool
		input         []int
		expectedIndex int
		expectedFound bool
	}{
		{
			name: "find_first_elem",
			whereFunc: func(val int) bool {
				return val == 0
			},
			input:         []int{0, 1, 2, 3, 4},
			expectedIndex: 0,
			expectedFound: true,
		},
		{
			name: "find_last_elem",
			whereFunc: func(val int) bool {
				return val == 4
			},
			input:         []int{0, 1, 2, 3, 4},
			expectedIndex: 4,
			expectedFound: true,
		},
		{
			name: "no_matching",
			whereFunc: func(val int) bool {
				return val > 5
			},
			input:         []int{0, 1, 2, 3, 4},
			expectedIndex: -1,
			expectedFound: false,
		},
		{
			name: "multi_matching",
			whereFunc: func(val int) bool {
				return val%2 == 0
			},
			input:         []int{0, 1, 2, 3, 4},
			expectedIndex: 4,
			expectedFound: true,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actualIndex, actualFound := LastWhere(tc.input, tc.whereFunc)
			assert.Equal(t, tc.expectedIndex, actualIndex)
			if assert.Equal(t, tc.expectedFound, actualFound) && tc.expectedFound {
				assert.Equal(t, tc.input[tc.expectedIndex], tc.input[actualIndex])
			}
		})
	}
}

func TestAnyWhere(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name      string
		whereFunc func(val int) bool
		input     []int
		expected  bool
	}{
		{
			name: "find_first_elem",
			whereFunc: func(val int) bool {
				return val == 0
			},
			input:    []int{0, 1, 2, 3, 4},
			expected: true,
		},
		{
			name: "find_last_elem",
			whereFunc: func(val int) bool {
				return val == 3
			},
			input:    []int{0, 1, 2, 3, 4},
			expected: true,
		},
		{
			name: "no_matching",
			whereFunc: func(val int) bool {
				return val == 5
			},
			input:    []int{0, 1, 2, 3, 4},
			expected: false,
		},
		{
			name: "multi_matching",
			whereFunc: func(val int) bool {
				return val%2 == 0
			},
			input:    []int{0, 1, 2, 3, 4},
			expected: true,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := AnyWhere(tc.input, tc.whereFunc)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestAllWhere(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name      string
		whereFunc func(val int) bool
		input     []int
		expected  bool
	}{
		{
			name: "no_matching",
			whereFunc: func(val int) bool {
				return val == 5
			},
			input:    []int{0, 1, 2, 3, 4},
			expected: false,
		},
		{
			name: "at_least_one_matching",
			whereFunc: func(val int) bool {
				return val%2 == 0
			},
			input:    []int{0, 1, 2, 3, 4},
			expected: false,
		},
		{
			name: "all_matching",
			whereFunc: func(val int) bool {
				return val < 5
			},
			input:    []int{0, 1, 2, 3, 4},
			expected: true,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := AllWhere(tc.input, tc.whereFunc)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestDedupe(t *testing.T) {
	t.Parallel()
	tt := []struct {
		Name     string
		Input    []string
		Expected []string
	}{
		{
			Name:     "empty",
			Input:    []string{},
			Expected: []string{},
		},
		{
			Name:     "nil",
			Input:    nil,
			Expected: nil,
		},
		{
			Name:     "single element",
			Input:    []string{"a"},
			Expected: []string{"a"},
		},
		{
			Name:     "no duplicates",
			Input:    []string{"a", "b", "c"},
			Expected: []string{"a", "b", "c"},
		},
		{
			Name:     "duplicates",
			Input:    []string{"a", "b", "c", "b"},
			Expected: []string{"a", "b", "c"},
		},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			var actual []string
			require.NotPanics(t, func() {
				actual = Dedupe(tc.Input)
			})
			assert.Equal(t, tc.Expected, actual)
		})
	}
}

func TestDedupeFunc(t *testing.T) {
	t.Parallel()
	tt := []struct {
		Name     string
		Input    []struct{ Key string }
		Expected []struct{ Key string }
	}{
		{
			Name:     "empty",
			Input:    []struct{ Key string }{},
			Expected: []struct{ Key string }{},
		},
		{
			Name:     "nil",
			Input:    nil,
			Expected: nil,
		},
		{
			Name: "single element",
			Input: []struct{ Key string }{
				{Key: "a"},
			},
			Expected: []struct{ Key string }{
				{Key: "a"},
			},
		},
		{
			Name: "no duplicates",
			Input: []struct{ Key string }{
				{Key: "a"},
				{Key: "b"},
				{Key: "c"},
			},
			Expected: []struct{ Key string }{
				{Key: "a"},
				{Key: "b"},
				{Key: "c"},
			},
		},
		{
			Name: "duplicates",
			Input: []struct{ Key string }{
				{Key: "a"},
				{Key: "b"},
				{Key: "c"},
				{Key: "a"},
			},
			Expected: []struct{ Key string }{
				{Key: "a"},
				{Key: "b"},
				{Key: "c"},
			},
		},
	}

	dedupeFunc := func(in struct{ Key string }) string {
		return in.Key
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			var actual []struct{ Key string }
			require.NotPanics(t, func() {
				actual = DedupeFunc(tc.Input, dedupeFunc)
			})
			assert.ElementsMatch(t, tc.Expected, actual)
		})
	}
}
