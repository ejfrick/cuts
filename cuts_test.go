package cuts

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
			input:         []int{0, 1, 2, 3, 4},
			expectedIndex: 2,
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

func TestDedupeFunc(t *testing.T) {
	elems := []struct {
		Name string
	}{
		{
			Name: "first",
		},
		{
			Name: "second",
		},
		{
			Name: "third",
		},
		{
			Name: "third",
		},
	}

	dedupeFunc := func(in struct{ Name string }) string {
		return in.Name
	}

	expectedLen := 3
	out := DedupeFunc(elems, dedupeFunc)
	assert.Equal(t, expectedLen, len(out))
}
