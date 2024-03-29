package cuts

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// SnapTo returns the element in an array that is closest to the target.
func SnapTo[S ~[]E, E constraints.Integer | constraints.Float](vals S, target E) E {
	var res E
	index, found := slices.BinarySearch(vals, target)
	switch {
	case found:
		res = target
	case len(vals) < index+1:
		res = vals[len(vals)-1]
	case index == 0:
		res = vals[0]
	default:
		prevDist := target - vals[index-1]
		nextDist := vals[index] - target

		if prevDist < nextDist {
			res = vals[index-1]
		} else {
			res = vals[index]
		}
	}

	return res
}

// SnapToFunc returns the element in an array that is closest to the target,
// given a custom comparison and closest function.
func SnapToFunc[S ~[]E, E any](vals S, target E, cmp func(E, E) int, closest func(tgt, nxt, prv E) E) E {
	var res E
	index, found := slices.BinarySearchFunc(vals, target, cmp)
	switch {
	case found:
		res = target
	case len(vals) < index+1:
		res = vals[len(vals)-1]
	default:
		res = closest(target, vals[index], vals[index-1])
	}

	return res
}
