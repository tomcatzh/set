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
				So(ss, ShouldEqual, "Set{One Two}")
			})
		})
	})
}
