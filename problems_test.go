package euler

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestProblem006(t *testing.T) {

	Convey("Sum square difference", t, func() {
		f_x := Problem006(10)
		So(f_x, ShouldNotEqual, 0)
		So(f_x, ShouldEqual, 2640)
	})

	Convey("Sum square difference", t, func() {
		f_x := Problem006(100)
		So(f_x, ShouldNotEqual, 0)
		So(f_x, ShouldEqual, 25164150)
	})
}

func TestProblem005(t *testing.T) {

	Convey("Smallest multiple", t, func() {
		f_x := Problem005(10)
		So(f_x, ShouldNotEqual, 0)
		So(f_x, ShouldEqual, 2520)
	})

	Convey("Smallest multiple", t, func() {
		f_x := Problem005(20)
		So(f_x, ShouldNotEqual, 0)
		So(f_x, ShouldEqual, 232792560)
	})
}

func TestProblem004(t *testing.T) {

	Convey("The largest palindrome made from the product of two 2-digit numbers", t, func() {
		f_x := Problem004(99)
		So(f_x, ShouldNotEqual, 0)
		So(f_x, ShouldEqual, 9009)
	})

	Convey("The largest palindrome made from the product of two 3-digit numbers", t, func() {
		f_x := Problem004(999)
		So(f_x, ShouldNotEqual, 0)
		So(f_x, ShouldEqual, 906609)
	})
}

func TestProblem003(t *testing.T) {

	Convey("The largest prime factor of 13195", t, func() {
		f_x := Problem003(13195)
		So(f_x, ShouldNotEqual, 0)
		So(f_x, ShouldEqual, 29)
	})

	Convey("The largest prime factor of 600851475143", t, func() {
		f_x := Problem003(600851475143)
		So(f_x, ShouldNotEqual, 0)
		So(f_x, ShouldEqual, 6857)
	})
}

func TestProblem002(t *testing.T) {

	Convey("The Sum of Even Fibonaci Numbers below 4 Million", t, func() {
		f_x := Problem002()
		So(f_x, ShouldNotEqual, 0)
		So(f_x, ShouldEqual, 4613732)
	})
}

func TestProblem001(t *testing.T) {

	Convey("Given an upperbound of 10", t, func() {
		x := 10
		f_x := Problem001(x)

		So(f_x, ShouldEqual, 23)

	})

	Convey("Given an upperbound of 1000", t, func() {
		x := 1000
		f_x := Problem001(x)

		So(f_x, ShouldEqual, 233168)

	})
}

func BenchmarkProblem002(b *testing.B) {

	for n := 0; n < b.N; n++ {
		Problem002()
	}
}

func BenchmarkProblem001(b *testing.B) {

	for n := 0; n < b.N; n++ {
		Problem001(1000)
	}
}
