// Copyright 2015 Tomcat Zhang. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This an execises for for 'Go Programming & Concurrency in Practice'
// It provided such as HashSet data structs.
package set

// HashSet just like java.util.HashSet
type HashSet struct {
	m map[interface{}]bool
}

// NewHashSet create a HashSet instance.
func NewHashSet() *HashSet {
	return &HashSet{m: make(map[interface{}]bool)}
}

// Add an element to the set.
// If success return true. If the element already in the set return false.
func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	}

	return false
}

// Remove an element from the set.
func (set *HashSet) Remove(e interface{}) {
	delete(set.m, e)
}

// Clear all elements in the set.
func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool)
}

// Check if an element is in the set.
func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

// Return how many elements is in the set
func (set *HashSet) Len() int {
	return len(set.m)
}

func (set *HashSet) Same(other *HashSet) bool {
	if other == nil {
		return false
	}

	if set.Len() != other.Len() {
		return false
	}

	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
	}

	return true
}
