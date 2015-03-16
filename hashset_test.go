package set

import (
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
