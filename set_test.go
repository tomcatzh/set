package set

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSuperSet(t *testing.T) {
	Convey("Create two new HashSet", t, func() {
		super := NewSimpleSet()
		sub := NewSimpleSet()

		Convey("If the sets is nil should return false", func() {
			So(IsSuperset(super, nil), ShouldBeFalse)
			So(IsSuperset(nil, super), ShouldBeFalse)
			So(IsSuperset(nil, nil), ShouldBeFalse)
		})

		Convey("Super set is empty should always return false", func() {
			So(IsSuperset(super, sub), ShouldBeFalse)
		})

		Convey("Add elements \"One\" \"Two\" to super set should be OK", func() {
			So(super.Add("One"), ShouldBeTrue)
			So(super.Add("Two"), ShouldBeTrue)

			Convey("Sub set is empty should return true", func() {
				So(IsSuperset(super, sub), ShouldBeTrue)
			})

			Convey("Add elements \"One\" \"Two\" \"Three\" to sub set should be OK", func() {
				So(sub.Add("One"), ShouldBeTrue)
				So(sub.Add("Two"), ShouldBeTrue)
				So(sub.Add("Three"), ShouldBeTrue)

				Convey("Sub set bigger than super set should be false", func() {
					So(IsSuperset(super, sub), ShouldBeFalse)
				})
			})

			Convey("Add elements \"Three\" to sub set should be OK", func() {
				So(sub.Add("Three"), ShouldBeTrue)

				Convey("Sub set have the element does not contain in super set return false", func() {
					So(IsSuperset(super, sub), ShouldBeFalse)
				})
			})

			Convey("Add elements \"One\" \"Two\" to sub set should be OK", func() {
				So(sub.Add("One"), ShouldBeTrue)
				So(sub.Add("Two"), ShouldBeTrue)

				Convey("Sub set is same of super set return false", func() {
					So(IsSuperset(super, sub), ShouldBeFalse)
				})
			})

			Convey("Add elements \"One\" to sub set should be OK", func() {
				So(sub.Add("One"), ShouldBeTrue)

				Convey("Super bigger than sub set and contain all then sub set's elements should return true", func() {
					So(IsSuperset(super, sub), ShouldBeTrue)
				})
			})

		})
	})
}

func TestUnion(t *testing.T) {
	Convey("Create two new HashSet", t, func() {
		a := NewSimpleSet()
		b := NewSimpleSet()

		Convey("Union a nil set should cause a panic", func() {
			So(func() { Union(a, nil) }, ShouldPanicWith, "The set is nil")
			So(func() { Union(nil, a) }, ShouldPanicWith, "The set is nil")
			So(func() { Union(nil, nil) }, ShouldPanicWith, "The set is nil")
		})

		Convey("Add \"One\" \"Two\" to set a and add \"Three\" to set b", func() {
			So(a.Add("One"), ShouldBeTrue)
			So(a.Add("Two"), ShouldBeTrue)
			So(b.Add("Three"), ShouldBeTrue)

			checkOrig := func() {
				So(a.Len(), ShouldEqual, 2)
				So(a.Contains("One"), ShouldBeTrue)
				So(a.Contains("Two"), ShouldBeTrue)
				So(b.Len(), ShouldEqual, 1)
				So(b.Contains("Three"), ShouldBeTrue)
			}

			checkUnion := func(u Set) {
				So(u.Len(), ShouldEqual, 3)
				So(u.Contains("One"), ShouldBeTrue)
				So(u.Contains("Two"), ShouldBeTrue)
				So(u.Contains("Three"), ShouldBeTrue)
			}

			Convey("A set union b set should be OK, both a and b should not be changed", func() {
				u := Union(a, b)

				checkUnion(u)

				checkOrig()
			})

			Convey("B set union a set should be OK, both a and b should not be changed", func() {
				u := Union(b, a)

				checkUnion(u)

				checkOrig()
			})
		})
	})
}

func TestIntersect(t *testing.T) {
	Convey("Create two new HashSet", t, func() {
		a := NewSimpleSet()
		b := NewSimpleSet()

		Convey("Intersect a nil set should cause a panic", func() {
			So(func() { Intersect(a, nil) }, ShouldPanicWith, "The set is nil")
			So(func() { Intersect(nil, a) }, ShouldPanicWith, "The set is nil")
			So(func() { Intersect(nil, nil) }, ShouldPanicWith, "The set is nil")
		})

		Convey("Two empty set have an empty intersect set", func() {
			i := Intersect(a, b)

			So(i.Len(), ShouldEqual, 0)
		})

		Convey("Add \"One\" \"Two\" to set a and add \"Three\" to set b", func() {
			So(a.Add("One"), ShouldBeTrue)
			So(a.Add("Two"), ShouldBeTrue)
			So(b.Add("Three"), ShouldBeTrue)

			checkOrig := func() {
				So(a.Len(), ShouldEqual, 2)
				So(a.Contains("One"), ShouldBeTrue)
				So(a.Contains("Two"), ShouldBeTrue)
				So(b.Len(), ShouldEqual, 1)
				So(b.Contains("Three"), ShouldBeTrue)
			}

			Convey("A set intersect b set should be OK, both a and b should not be changed", func() {
				i := Intersect(a, b)

				So(i.Len(), ShouldEqual, 0)

				checkOrig()
			})

			Convey("B set intersect a set should be OK, both a and b should not be changed", func() {
				i := Intersect(b, a)

				So(i.Len(), ShouldEqual, 0)

				checkOrig()
			})
		})

		Convey("Add \"One\" \"Two\" to set a and add \"Two\" \"Three\" to set b", func() {
			So(a.Add("One"), ShouldBeTrue)
			So(a.Add("Two"), ShouldBeTrue)
			So(b.Add("Two"), ShouldBeTrue)
			So(b.Add("Three"), ShouldBeTrue)
			So(b.Add("Four"), ShouldBeTrue)

			checkOrig := func() {
				So(a.Len(), ShouldEqual, 2)
				So(a.Contains("One"), ShouldBeTrue)
				So(a.Contains("Two"), ShouldBeTrue)
				So(b.Len(), ShouldEqual, 3)
				So(b.Contains("Two"), ShouldBeTrue)
				So(b.Contains("Three"), ShouldBeTrue)
				So(b.Contains("Four"), ShouldBeTrue)
			}

			checkUnion := func(i Set) {
				So(i.Len(), ShouldEqual, 1)
				So(i.Contains("Two"), ShouldBeTrue)
			}

			Convey("A set intersect b set should be OK, both a and b should not be changed", func() {
				i := Intersect(a, b)

				checkUnion(i)

				checkOrig()
			})

			Convey("B set intersect a set should be OK, both a and b should not be changed", func() {
				i := Intersect(b, a)

				checkUnion(i)

				checkOrig()
			})
		})
	})
}

func TestDifference(t *testing.T) {
	Convey("Create two new HashSet", t, func() {
		a := NewSimpleSet()
		b := NewSimpleSet()

		Convey("Difference a nil set should cause a panic", func() {
			So(func() { Difference(a, nil) }, ShouldPanicWith, "The set is nil")
			So(func() { Difference(nil, a) }, ShouldPanicWith, "The set is nil")
			So(func() { Difference(nil, nil) }, ShouldPanicWith, "The set is nil")
		})

		Convey("Add \"One\" \"Two\" to set a", func() {
			So(a.Add("One"), ShouldBeTrue)
			So(a.Add("Two"), ShouldBeTrue)

			check := func(a Set) {
				So(a.Len(), ShouldEqual, 2)
				So(a.Contains("One"), ShouldBeTrue)
				So(a.Contains("Two"), ShouldBeTrue)
			}

			Convey("A set difference empty should be same of a, both a and b should not be changed", func() {
				d := Difference(a, b)

				check(d)
				check(a)
			})

			Convey("Add \"Three\" to set b", func() {
				So(b.Add("Three"), ShouldBeTrue)

				checkOrig := func() {
					So(a.Len(), ShouldEqual, 2)
					So(a.Contains("One"), ShouldBeTrue)
					So(a.Contains("Two"), ShouldBeTrue)
					So(b.Len(), ShouldEqual, 1)
					So(b.Contains("Three"), ShouldBeTrue)
				}

				Convey("A set difference b set should be OK, both a and b should not be changed", func() {
					d := Difference(a, b)

					So(d.Len(), ShouldEqual, 2)
					So(d.Contains("One"), ShouldBeTrue)
					So(d.Contains("Two"), ShouldBeTrue)

					checkOrig()
				})

				Convey("B set difference a set should be OK, both a and b should not be changed", func() {
					d := Difference(b, a)

					So(d.Len(), ShouldEqual, 1)
					So(d.Contains("Three"), ShouldBeTrue)

					checkOrig()
				})
			})
		})

	})
}

func TestSymmetricDifference(t *testing.T) {
	Convey("Create two new HashSet", t, func() {
		a := NewSimpleSet()
		b := NewSimpleSet()

		Convey("SymmetricDifference a nil set should cause a panic", func() {
			So(func() { SymmetricDifference(a, nil) }, ShouldPanicWith, "The set is nil")
			So(func() { SymmetricDifference(nil, a) }, ShouldPanicWith, "The set is nil")
			So(func() { SymmetricDifference(nil, nil) }, ShouldPanicWith, "The set is nil")
		})

		Convey("Add \"One\" \"Two\" \"Three\" to set a and add \"Three\" \"Four\" to set b", func() {
			So(a.Add("One"), ShouldBeTrue)
			So(a.Add("Two"), ShouldBeTrue)
			So(a.Add("Three"), ShouldBeTrue)
			So(b.Add("Three"), ShouldBeTrue)
			So(b.Add("Four"), ShouldBeTrue)

			checkOrig := func() {
				So(a.Len(), ShouldEqual, 3)
				So(a.Contains("One"), ShouldBeTrue)
				So(a.Contains("Two"), ShouldBeTrue)
				So(b.Contains("Three"), ShouldBeTrue)
				So(b.Len(), ShouldEqual, 2)
				So(b.Contains("Three"), ShouldBeTrue)
				So(b.Contains("Four"), ShouldBeTrue)
			}

			checkUnion := func(u Set) {
				So(u.Len(), ShouldEqual, 3)
				So(u.Contains("One"), ShouldBeTrue)
				So(u.Contains("Two"), ShouldBeTrue)
				So(u.Contains("Four"), ShouldBeTrue)
			}

			Convey("A set SymmetricDifference b set should be OK, both a and b should not be changed", func() {
				d := SymmetricDifference(a, b)

				checkUnion(d)
				checkOrig()
			})

			Convey("B set SymmetricDifference a set should be OK, both a and b should not be changed", func() {
				d := SymmetricDifference(b, a)

				checkUnion(d)
				checkOrig()
			})
		})
	})
}
