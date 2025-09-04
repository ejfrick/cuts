package cuts

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type snapToCase string

const (
	snapToPrev     snapToCase = "snap_to_prev"
	snapToNext     snapToCase = "snap_to_next"
	equidistant    snapToCase = "equidistant"
	containsTarget snapToCase = "contains_target"
	snapToLast     snapToCase = "snap_to_last_element"
	snapToFirst    snapToCase = "snap_to_first_element"
)

type SuiteSnapTo struct {
	suite.Suite
}

func TestSuiteSnapTo(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(SuiteSnapTo))
}

func (s *SuiteSnapTo) TestTimeDuration() {
	s.T().Parallel()

	tt := []struct {
		name     snapToCase
		input    []time.Duration
		target   time.Duration
		expected time.Duration
	}{
		{
			name:     snapToPrev,
			input:    []time.Duration{time.Minute, time.Minute * 10, time.Minute * 20},
			target:   time.Minute * 4,
			expected: time.Minute,
		},
		{
			name:     snapToNext,
			input:    []time.Duration{time.Minute, time.Minute * 10, time.Minute * 20},
			target:   time.Minute * 6,
			expected: time.Minute * 10,
		},
		{
			name:     equidistant,
			input:    []time.Duration{time.Minute, time.Minute * 3},
			target:   time.Minute * 2,
			expected: time.Minute * 3,
		},
		{
			name:     containsTarget,
			input:    []time.Duration{time.Minute, time.Minute * 10, time.Minute * 20},
			target:   time.Minute * 10,
			expected: time.Minute * 10,
		},
		{
			name:     snapToLast,
			input:    []time.Duration{time.Minute, time.Minute * 10, time.Minute * 20},
			target:   time.Minute * 30,
			expected: time.Minute * 20,
		},
		{
			name:     snapToFirst,
			input:    []time.Duration{time.Minute, time.Minute * 10, time.Minute * 20},
			target:   time.Second,
			expected: time.Minute,
		},
	}

	for _, tc := range tt {
		tc := tc
		s.Run(string(tc.name), func() {
			actual := SnapTo(tc.input, tc.target)
			s.Equal(tc.expected, actual)
			s.IsType(tc.expected, actual)
		})
	}
}

func (s *SuiteSnapTo) TestInt() {
	s.T().Parallel()

	tt := []struct {
		name     snapToCase
		input    []int
		target   int
		expected int
	}{
		{
			name:     snapToPrev,
			input:    []int{1, 10, 20},
			target:   4,
			expected: 1,
		},
		{
			name:     snapToNext,
			input:    []int{1, 10, 20},
			target:   6,
			expected: 10,
		},
		{
			name:     equidistant,
			input:    []int{1, 3},
			target:   2,
			expected: 3,
		},
		{
			name:     containsTarget,
			input:    []int{1, 10, 20},
			target:   10,
			expected: 10,
		},
		{
			name:     snapToLast,
			input:    []int{1, 10, 20},
			target:   30,
			expected: 20,
		},
		{
			name:     snapToFirst,
			input:    []int{1, 10, 20},
			target:   0,
			expected: 1,
		},
	}

	for _, tc := range tt {
		tc := tc
		s.Run(string(tc.name), func() {
			actual := SnapTo(tc.input, tc.target)
			s.Equal(tc.expected, actual)
			s.IsType(tc.expected, actual)
		})
	}
}

func (s *SuiteSnapTo) TestInt8() {
	s.T().Parallel()
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	tt := []struct {
		name     snapToCase
		input    []int8
		target   int8
		expected int8
	}{
		{
			name:     snapToPrev,
			input:    []int8{1, 10, 20},
			target:   4,
			expected: 1,
		},
		{
			name:     snapToNext,
			input:    []int8{1, 10, 20},
			target:   6,
			expected: 10,
		},
		{
			name:     equidistant,
			input:    []int8{1, 3},
			target:   2,
			expected: 3,
		},
		{
			name:     containsTarget,
			input:    []int8{1, 10, 20},
			target:   10,
			expected: 10,
		},
		{
			name:     snapToLast,
			input:    []int8{1, 10, 20},
			target:   30,
			expected: 20,
		},
		{
			name:     snapToFirst,
			input:    []int8{1, 10, 20},
			target:   0,
			expected: 1,
		},
	}

	for _, tc := range tt {
		tc := tc
		s.Run(string(tc.name), func() {
			actual := SnapTo(tc.input, tc.target)
			s.Equal(tc.expected, actual)
			s.IsType(tc.expected, actual)
		})
	}
}

func (s *SuiteSnapTo) TestInt16() {
	s.T().Parallel()
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	tt := []struct {
		name     snapToCase
		input    []int16
		target   int16
		expected int16
	}{
		{
			name:     snapToPrev,
			input:    []int16{1, 10, 20},
			target:   4,
			expected: 1,
		},
		{
			name:     snapToNext,
			input:    []int16{1, 10, 20},
			target:   6,
			expected: 10,
		},
		{
			name:     equidistant,
			input:    []int16{1, 3},
			target:   2,
			expected: 3,
		},
		{
			name:     containsTarget,
			input:    []int16{1, 10, 20},
			target:   10,
			expected: 10,
		},
		{
			name:     snapToLast,
			input:    []int16{1, 10, 20},
			target:   30,
			expected: 20,
		},
		{
			name:     snapToFirst,
			input:    []int16{1, 10, 20},
			target:   0,
			expected: 1,
		},
	}

	for _, tc := range tt {
		tc := tc
		s.Run(string(tc.name), func() {
			actual := SnapTo(tc.input, tc.target)
			s.Equal(tc.expected, actual)
			s.IsType(tc.expected, actual)
		})
	}
}

func (s *SuiteSnapTo) TestInt32() {
	s.T().Parallel()
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	tt := []struct {
		name     snapToCase
		input    []int32
		target   int32
		expected int32
	}{
		{
			name:     snapToPrev,
			input:    []int32{1, 10, 20},
			target:   4,
			expected: 1,
		},
		{
			name:     snapToNext,
			input:    []int32{1, 10, 20},
			target:   6,
			expected: 10,
		},
		{
			name:     equidistant,
			input:    []int32{1, 3},
			target:   2,
			expected: 3,
		},
		{
			name:     containsTarget,
			input:    []int32{1, 10, 20},
			target:   10,
			expected: 10,
		},
		{
			name:     snapToLast,
			input:    []int32{1, 10, 20},
			target:   30,
			expected: 20,
		},
		{
			name:     snapToFirst,
			input:    []int32{1, 10, 20},
			target:   0,
			expected: 1,
		},
	}

	for _, tc := range tt {
		tc := tc
		s.Run(string(tc.name), func() {
			actual := SnapTo(tc.input, tc.target)
			s.Equal(tc.expected, actual)
			s.IsType(tc.expected, actual)
		})
	}
}

func (s *SuiteSnapTo) TestInt64() {
	s.T().Parallel()
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	tt := []struct {
		name     snapToCase
		input    []int64
		target   int64
		expected int64
	}{
		{
			name:     snapToPrev,
			input:    []int64{1, 10, 20},
			target:   4,
			expected: 1,
		},
		{
			name:     snapToNext,
			input:    []int64{1, 10, 20},
			target:   6,
			expected: 10,
		},
		{
			name:     equidistant,
			input:    []int64{1, 3},
			target:   2,
			expected: 3,
		},
		{
			name:     containsTarget,
			input:    []int64{1, 10, 20},
			target:   10,
			expected: 10,
		},
		{
			name:     snapToLast,
			input:    []int64{1, 10, 20},
			target:   30,
			expected: 20,
		},
		{
			name:     snapToLast,
			input:    []int64{1, 10, 20},
			target:   0,
			expected: 1,
		},
	}

	for _, tc := range tt {
		tc := tc
		s.Run(string(tc.name), func() {
			actual := SnapTo(tc.input, tc.target)
			s.Equal(tc.expected, actual)
			s.IsType(tc.expected, actual)
		})
	}
}

func (s *SuiteSnapTo) TestUint() {
	s.T().Parallel()

	tt := []struct {
		name     snapToCase
		input    []uint
		target   uint
		expected uint
	}{
		{
			name:     snapToPrev,
			input:    []uint{1, 10, 20},
			target:   4,
			expected: 1,
		},
		{
			name:     snapToNext,
			input:    []uint{1, 10, 20},
			target:   6,
			expected: 10,
		},
		{
			name:     equidistant,
			input:    []uint{1, 3},
			target:   2,
			expected: 3,
		},
		{
			name:     containsTarget,
			input:    []uint{1, 10, 20},
			target:   10,
			expected: 10,
		},
		{
			name:     snapToLast,
			input:    []uint{1, 10, 20},
			target:   30,
			expected: 20,
		},
		{
			name:     snapToFirst,
			input:    []uint{1, 10, 20},
			target:   0,
			expected: 1,
		},
	}

	for _, tc := range tt {
		tc := tc
		s.Run(string(tc.name), func() {
			actual := SnapTo(tc.input, tc.target)
			s.Equal(tc.expected, actual)
			s.IsType(tc.expected, actual)
		})
	}
}

func (s *SuiteSnapTo) TestUint8() {
	s.T().Parallel()
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	tt := []struct {
		name     snapToCase
		input    []uint8
		target   uint8
		expected uint8
	}{
		{
			name:     snapToPrev,
			input:    []uint8{1, 10, 20},
			target:   4,
			expected: 1,
		},
		{
			name:     snapToNext,
			input:    []uint8{1, 10, 20},
			target:   6,
			expected: 10,
		},
		{
			name:     equidistant,
			input:    []uint8{1, 3},
			target:   2,
			expected: 3,
		},
		{
			name:     containsTarget,
			input:    []uint8{1, 10, 20},
			target:   10,
			expected: 10,
		},
		{
			name:     snapToLast,
			input:    []uint8{1, 10, 20},
			target:   30,
			expected: 20,
		},
		{
			name:     snapToFirst,
			input:    []uint8{1, 10, 20},
			target:   0,
			expected: 1,
		},
	}

	for _, tc := range tt {
		tc := tc
		s.Run(string(tc.name), func() {
			actual := SnapTo(tc.input, tc.target)
			s.Equal(tc.expected, actual)
			s.IsType(tc.expected, actual)
		})
	}
}

func (s *SuiteSnapTo) TestUint16() {
	s.T().Parallel()
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	tt := []struct {
		name     snapToCase
		input    []uint16
		target   uint16
		expected uint16
	}{
		{
			name:     snapToPrev,
			input:    []uint16{1, 10, 20},
			target:   4,
			expected: 1,
		},
		{
			name:     snapToNext,
			input:    []uint16{1, 10, 20},
			target:   6,
			expected: 10,
		},
		{
			name:     equidistant,
			input:    []uint16{1, 3},
			target:   2,
			expected: 3,
		},
		{
			name:     containsTarget,
			input:    []uint16{1, 10, 20},
			target:   10,
			expected: 10,
		},
		{
			name:     snapToLast,
			input:    []uint16{1, 10, 20},
			target:   30,
			expected: 20,
		},
		{
			name:     snapToFirst,
			input:    []uint16{1, 10, 20},
			target:   0,
			expected: 1,
		},
	}

	for _, tc := range tt {
		tc := tc
		s.Run(string(tc.name), func() {
			actual := SnapTo(tc.input, tc.target)
			s.Equal(tc.expected, actual)
			s.IsType(tc.expected, actual)
		})
	}
}

func (s *SuiteSnapTo) TestUint32() {
	s.T().Parallel()
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	tt := []struct {
		name     snapToCase
		input    []uint32
		target   uint32
		expected uint32
	}{
		{
			name:     snapToPrev,
			input:    []uint32{1, 10, 20},
			target:   4,
			expected: 1,
		},
		{
			name:     snapToNext,
			input:    []uint32{1, 10, 20},
			target:   6,
			expected: 10,
		},
		{
			name:     equidistant,
			input:    []uint32{1, 3},
			target:   2,
			expected: 3,
		},
		{
			name:     containsTarget,
			input:    []uint32{1, 10, 20},
			target:   10,
			expected: 10,
		},
		{
			name:     snapToLast,
			input:    []uint32{1, 10, 20},
			target:   30,
			expected: 20,
		},
		{
			name:     snapToFirst,
			input:    []uint32{1, 10, 20},
			target:   0,
			expected: 1,
		},
	}

	for _, tc := range tt {
		tc := tc
		s.Run(string(tc.name), func() {
			actual := SnapTo(tc.input, tc.target)
			s.Equal(tc.expected, actual)
			s.IsType(tc.expected, actual)
		})
	}
}

func (s *SuiteSnapTo) TestUint64() {
	s.T().Parallel()
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	tt := []struct {
		name     snapToCase
		input    []uint64
		target   uint64
		expected uint64
	}{
		{
			name:     snapToPrev,
			input:    []uint64{1, 10, 20},
			target:   4,
			expected: 1,
		},
		{
			name:     snapToNext,
			input:    []uint64{1, 10, 20},
			target:   6,
			expected: 10,
		},
		{
			name:     equidistant,
			input:    []uint64{1, 3},
			target:   2,
			expected: 3,
		},
		{
			name:     containsTarget,
			input:    []uint64{1, 10, 20},
			target:   10,
			expected: 10,
		},
		{
			name:     snapToLast,
			input:    []uint64{1, 10, 20},
			target:   30,
			expected: 20,
		},
		{
			name:     snapToFirst,
			input:    []uint64{1, 10, 20},
			target:   0,
			expected: 1,
		},
	}

	for _, tc := range tt {
		tc := tc
		s.Run(string(tc.name), func() {
			actual := SnapTo(tc.input, tc.target)
			s.Equal(tc.expected, actual)
			s.IsType(tc.expected, actual)
		})
	}
}

func (s *SuiteSnapTo) TestUintptr() {
	s.T().Parallel()

	tt := []struct {
		name     snapToCase
		input    []uintptr
		target   uintptr
		expected uintptr
	}{
		{
			name:     snapToPrev,
			input:    []uintptr{1, 10, 20},
			target:   4,
			expected: 1,
		},
		{
			name:     snapToNext,
			input:    []uintptr{1, 10, 20},
			target:   6,
			expected: 10,
		},
		{
			name:     equidistant,
			input:    []uintptr{1, 3},
			target:   2,
			expected: 3,
		},
		{
			name:     containsTarget,
			input:    []uintptr{1, 10, 20},
			target:   10,
			expected: 10,
		},
		{
			name:     snapToLast,
			input:    []uintptr{1, 10, 20},
			target:   30,
			expected: 20,
		},
		{
			name:     snapToFirst,
			input:    []uintptr{1, 10, 20},
			target:   0,
			expected: 1,
		},
	}

	for _, tc := range tt {
		tc := tc
		s.Run(string(tc.name), func() {
			actual := SnapTo(tc.input, tc.target)
			s.Equal(tc.expected, actual)
			s.IsType(tc.expected, actual)
		})
	}
}

func (s *SuiteSnapTo) TestFloat32() {
	s.T().Parallel()
	if testing.Short() {
		s.T().Skip("skipping test in short mode.")
	}

	tt := []struct {
		name     snapToCase
		input    []float32
		target   float32
		expected float32
	}{
		{
			name:     snapToPrev,
			input:    []float32{1, 10, 20},
			target:   4,
			expected: 1,
		},
		{
			name:     snapToNext,
			input:    []float32{1, 10, 20},
			target:   6,
			expected: 10,
		},
		{
			name:     equidistant,
			input:    []float32{1, 3},
			target:   2,
			expected: 3,
		},
		{
			name:     containsTarget,
			input:    []float32{1, 10, 20},
			target:   10,
			expected: 10,
		},
		{
			name:     snapToLast,
			input:    []float32{1, 10, 20},
			target:   30,
			expected: 20,
		},
		{
			name:     snapToFirst,
			input:    []float32{1, 10, 20},
			target:   0,
			expected: 1,
		},
	}

	for _, tc := range tt {
		tc := tc
		s.Run(string(tc.name), func() {
			actual := SnapTo(tc.input, tc.target)
			s.Equal(tc.expected, actual)
			s.IsType(tc.expected, actual)
		})
	}
}

func (s *SuiteSnapTo) TestFloat64() {
	s.T().Parallel()

	tt := []struct {
		name     snapToCase
		input    []float64
		target   float64
		expected float64
	}{
		{
			name:     snapToPrev,
			input:    []float64{1, 10, 20},
			target:   4,
			expected: 1,
		},
		{
			name:     snapToNext,
			input:    []float64{1, 10, 20},
			target:   6,
			expected: 10,
		},
		{
			name:     equidistant,
			input:    []float64{1, 3},
			target:   2,
			expected: 3,
		},
		{
			name:     containsTarget,
			input:    []float64{1, 10, 20},
			target:   10,
			expected: 10,
		},
		{
			name:     snapToLast,
			input:    []float64{1, 10, 20},
			target:   30,
			expected: 20,
		},
		{
			name:     snapToFirst,
			input:    []float64{1, 10, 20},
			target:   0,
			expected: 1,
		},
	}

	for _, tc := range tt {
		tc := tc
		s.Run(string(tc.name), func() {
			actual := SnapTo(tc.input, tc.target)
			s.Equal(tc.expected, actual)
			s.IsType(tc.expected, actual)
		})
	}
}
