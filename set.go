// Copyright 2015 Tomcat Zhang. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// This an execises for for 'Go Programming & Concurrency in Practice'
// It provided such as HashSet data structs.
package set

// Set interface
type Set interface {
	Add(e interface{}) bool
	Remove(e interface{})
	Clear()
	Contains(e interface{}) bool
	Len() int
	Same(other Set) bool
	Elements() []interface{}
	String() string
}

// Check is this set is a super set of other.
func IsSuperset(set, other Set) bool {
	if other == nil || set == nil {
		return false
	}

	oneLen := set.Len()
	otherLen := other.Len()

	if oneLen == 0 || oneLen <= otherLen {
		return false
	}

	if oneLen > 0 && otherLen == 0 {
		return true
	}

	for _, v := range other.Elements() {
		if !set.Contains(v) {
			return false
		}
	}

	return true
}

// Return the union set of the set and the other
func Union(set, other Set) Set {
	if other == nil || set == nil {
		panic("The set is nil")
	}

	result := NewSimpleSet()

	if set.Len() > 0 {
		for _, v := range set.Elements() {
			result.Add(v)
		}
	}

	if other.Len() > 0 {
		for _, v := range other.Elements() {
			result.Add(v)
		}
	}

	return result
}

// Return the intersect set of the set and the other
func Intersect(set, other Set) Set {
	if other == nil || set == nil {
		panic("The set is nil")
	}

	result := NewSimpleSet()

	if other.Len() == 0 || set.Len() == 0 {
		return result
	}

	if other.Len() > set.Len() {
		for _, v := range set.Elements() {
			if other.Contains(v) {
				result.Add(v)
			}
		}
	} else {
		for _, v := range other.Elements() {
			if set.Contains(v) {
				result.Add(v)
			}
		}
	}

	return result
}

// Return the difference of the set and the other
func Difference(set, other Set) Set {
	if other == nil || set == nil {
		panic("The set is nil")
	}

	result := NewSimpleSet()

	if set.Len() > 0 {
		if other.Len() == 0 {
			for _, v := range set.Elements() {
				result.Add(v)
			}
		} else {
			for _, v := range set.Elements() {
				if !other.Contains(v) {
					result.Add(v)
				}
			}
		}
	}

	return result
}

func SymmetricDifference(set, other Set) Set {
	if other == nil || set == nil {
		panic("The set is nil")
	}

	return Union(Difference(set, other), Difference(other, set))
}

func NewSimpleSet() Set {
	return NewHashSet()
}
