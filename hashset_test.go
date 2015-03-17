package set

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	Convey("Create a new HashSet", t, func() {
		s := NewHashSet()

		Convey("Add first element \"One\" to set should be OK", func() {
			So(s.Add("One"), ShouldBeTrue)

			Convey("The set length should be one", func() {
				So(s.Len(), ShouldEqual, 1)
			})

			Convey("The set should contains an \"One\" element", func() {
				So(s.Contains("One"), ShouldBeTrue)
			})

			Convey("Add the same element \"One\" to set should be fail", func() {
				So(s.Add("One"), ShouldBeFalse)
			})
		})
	})
}

func TestRemove(t *testing.T) {
	Convey("Create a new HashSet", t, func() {
		s := NewHashSet()

		Convey("Add elements \"One\" \"Two\" to set should be OK", func() {
			So(s.Add("One"), ShouldBeTrue)
			So(s.Add("Two"), ShouldBeTrue)

			Convey("The set length should be two", func() {
				So(s.Len(), ShouldEqual, 2)
			})

			Convey("The set should contains an \"One\" element", func() {
				So(s.Contains("One"), ShouldBeTrue)
			})

			Convey("Remove element \"One\" from set", func() {
				s.Remove("One")

				Convey("The set should not cotaines an \"One\" element", func() {
					So(s.Contains("One"), ShouldBeFalse)
				})

				Convey("The set length should be one", func() {
					So(s.Len(), ShouldEqual, 1)
				})

				Convey("Add the same element \"One\" to set should be true", func() {
					So(s.Add("One"), ShouldBeTrue)

					Convey("The set length should be two", func() {
						So(s.Len(), ShouldEqual, 2)
					})
				})
			})
		})
	})
}

func TestClear(t *testing.T) {
	Convey("Create a new HashSet", t, func() {
		s := NewHashSet()

		Convey("Add elements \"One\" \"Two\" to set should be OK", func() {
			So(s.Add("One"), ShouldBeTrue)
			So(s.Add("Two"), ShouldBeTrue)

			Convey("The set length should be two", func() {
				So(s.Len(), ShouldEqual, 2)
			})

			Convey("The set should contains an \"One\" element", func() {
				So(s.Contains("One"), ShouldBeTrue)
			})

			Convey("Clear the set", func() {
				s.Clear()

				Convey("The set length should be zero", func() {
					So(s.Len(), ShouldEqual, 0)
				})

				Convey("The set should not cotaines an \"One\" element", func() {
					So(s.Contains("One"), ShouldBeFalse)
				})

				Convey("Add the same element \"One\" to set should be true", func() {
					So(s.Add("One"), ShouldBeTrue)

					Convey("The set length should be one", func() {
						So(s.Len(), ShouldEqual, 1)
					})
				})
			})
		})
	})
}

func TestSame(t *testing.T) {
	Convey("Create a new HashSet", t, func() {
		a := NewHashSet()

		Convey("Add elements \"One\" \"Two\" to set should be OK", func() {
			So(a.Add("One"), ShouldBeTrue)
			So(a.Add("Two"), ShouldBeTrue)

			Convey("Other set is nil should return false", func() {
				So(a.Same(nil), ShouldBeFalse)
			})

			Convey("Create another HashSet", func() {
				b := NewHashSet()

				Convey("Add elements \"One\" to new set should be OK", func() {
					So(b.Add("One"), ShouldBeTrue)

					Convey("Other is less then the set should return false", func() {
						So(a.Same(b), ShouldBeFalse)
					})
				})

				Convey("Add elements \"One\" \"Two\" \"Three\" to new set should be OK", func() {
					So(b.Add("One"), ShouldBeTrue)
					So(b.Add("Two"), ShouldBeTrue)
					So(b.Add("Three"), ShouldBeTrue)

					Convey("Other is more then the set should return false", func() {
						So(a.Same(b), ShouldBeFalse)
					})
				})

				Convey("Add elements \"One\" \"Three\" to new set should be OK", func() {
					So(b.Add("One"), ShouldBeTrue)
					So(b.Add("Three"), ShouldBeTrue)

					Convey("Other is different from the set should return false", func() {
						So(a.Same(b), ShouldBeFalse)
					})
				})

				Convey("Add elements \"One\" \"Two\" to new set should be OK", func() {
					So(b.Add("One"), ShouldBeTrue)
					So(b.Add("Two"), ShouldBeTrue)

					Convey("Other is same of the set should return true", func() {
						So(a.Same(b), ShouldBeTrue)
					})
				})
			})
		})
	})
}

func TestElements(t *testing.T) {
	Convey("Create a new HashSet", t, func() {
		s := NewHashSet()

		Convey("Add elements \"One\" \"Two\" to set should be OK", func() {
			So(s.Add("One"), ShouldBeTrue)
			So(s.Add("Two"), ShouldBeTrue)

			Convey("Get elements from set", func() {
				e := s.Elements()

				Convey("Elements length should be two", func() {
					So(len(e), ShouldEqual, 2)
				})

				Convey("Elements should have \"One\"", func() {
					So(e, ShouldContain, "One")
				})

				Convey("Elements should have \"Two\"", func() {
					So(e, ShouldContain, "Two")
				})

				Convey("Elements should not have \"Three\"", func() {
					So(e, ShouldNotContain, "Three")
				})
			})
		})
	})
}

func TestString(t *testing.T) {
	Convey("Create a new HashSet", t, func() {
		s := NewHashSet()

		Convey("Add elements \"One\" \"Two\" to set should be OK", func() {
			So(s.Add("One"), ShouldBeTrue)
			So(s.Add("Two"), ShouldBeTrue)

			Convey("Testing the set to the string", func() {
				ss := fmt.Sprintf("%v", s)
				So(len(ss), ShouldEqual, len("Set{One Two}"))
				So(ss, ShouldStartWith, "Set{")
				So(ss, ShouldEndWith, "}")
				So(ss, ShouldContainSubstring, "One")
				So(ss, ShouldContainSubstring, "Two")
			})
		})
	})
}

func TestSuperAndSub(t *testing.T) {
	Convey("Create two new HashSet", t, func() {
		super := NewHashSet()
		sub := NewHashSet()

		Convey("If other set is nil should return false", func() {
			So(super.IsSuperset(nil), ShouldBeFalse)
			So(sub.IsSubset(nil), ShouldBeFalse)
		})

		Convey("Super set is empty should always return false", func() {
			So(super.IsSuperset(sub), ShouldBeFalse)

			Convey("Also check subset is false", func() {
				So(sub.IsSubset(super), ShouldBeFalse)
			})
		})

		Convey("Add elements \"One\" \"Two\" to super set should be OK", func() {
			So(super.Add("One"), ShouldBeTrue)
			So(super.Add("Two"), ShouldBeTrue)

			Convey("Sub set is empty should return true", func() {
				So(super.IsSuperset(sub), ShouldBeTrue)

				Convey("Also check subset is true", func() {
					So(sub.IsSubset(super), ShouldBeTrue)
				})
			})

			Convey("Add elements \"One\" \"Two\" \"Three\" to sub set should be OK", func() {
				So(sub.Add("One"), ShouldBeTrue)
				So(sub.Add("Two"), ShouldBeTrue)
				So(sub.Add("Three"), ShouldBeTrue)

				Convey("Sub set bigger than super set should be false", func() {
					So(super.IsSuperset(sub), ShouldBeFalse)

					Convey("Also check subset is false", func() {
						So(sub.IsSubset(super), ShouldBeFalse)
					})
				})
			})

			Convey("Add elements \"Three\" to sub set should be OK", func() {
				So(sub.Add("Three"), ShouldBeTrue)

				Convey("Sub set have the element does not contain in super set return false", func() {
					So(super.IsSuperset(sub), ShouldBeFalse)

					Convey("Also check subset is false", func() {
						So(sub.IsSubset(super), ShouldBeFalse)
					})
				})
			})

			Convey("Add elements \"One\" \"Two\" to sub set should be OK", func() {
				So(sub.Add("One"), ShouldBeTrue)
				So(sub.Add("Two"), ShouldBeTrue)

				Convey("Sub set is same of super set return false", func() {
					So(super.IsSuperset(sub), ShouldBeFalse)

					Convey("Also check subset is false", func() {
						So(sub.IsSubset(super), ShouldBeFalse)
					})
				})
			})

			Convey("Add elements \"One\" to sub set should be OK", func() {
				So(sub.Add("One"), ShouldBeTrue)

				Convey("Super bigger than sub set and contain all then sub set's elements should return true", func() {
					So(super.IsSuperset(sub), ShouldBeTrue)
				})

				Convey("Also check subset is true", func() {
					So(sub.IsSubset(super), ShouldBeTrue)
				})
			})

		})
	})
}

func TestUnion(t *testing.T) {
	Convey("Create two new HashSet", t, func() {
		a := NewHashSet()
		b := NewHashSet()

		Convey("Union a nil set should cause a panic", func() {
			So(func() { a.Union(nil) }, ShouldPanicWith, "Other set is nil")
		})

		Convey("Add \"One\" \"Two\" to set a and add \"Three\" to set b", func() {
			So(a.Add("One"), ShouldBeTrue)
			So(a.Add("Two"), ShouldBeTrue)
			So(b.Add("Three"), ShouldBeTrue)

			Convey("A set union b set should be OK, both a and b should not be changed", func() {
				u := a.Union(b)

				So(u.Len(), ShouldEqual, 3)
				So(u.Contains("One"), ShouldBeTrue)
				So(u.Contains("Two"), ShouldBeTrue)
				So(u.Contains("Three"), ShouldBeTrue)

				So(a.Len(), ShouldEqual, 2)
				So(a.Contains("One"), ShouldBeTrue)
				So(a.Contains("Two"), ShouldBeTrue)
				So(b.Len(), ShouldEqual, 1)
				So(b.Contains("Three"), ShouldBeTrue)
			})

			Convey("B set union a set should be OK, both a and b should not be changed", func() {
				u := b.Union(a)

				So(u.Len(), ShouldEqual, 3)
				So(u.Contains("One"), ShouldBeTrue)
				So(u.Contains("Two"), ShouldBeTrue)
				So(u.Contains("Three"), ShouldBeTrue)

				So(a.Len(), ShouldEqual, 2)
				So(a.Contains("One"), ShouldBeTrue)
				So(a.Contains("Two"), ShouldBeTrue)
				So(b.Len(), ShouldEqual, 1)
				So(b.Contains("Three"), ShouldBeTrue)
			})
		})
	})
}
