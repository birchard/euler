package euler

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSieveOfEratosthenes(t *testing.T) {

}

func TestRelativelyPrime(t *testing.T) {

	Convey("Test relativelyPrime", t, func() {
		Convey("3 and 5 are Co-Prime", func() {
			So(relativelyPrime(3, 5), ShouldBeTrue)
		})

		Convey("21 and 8 are Co-Prime", func() {
			So(relativelyPrime(21, 8), ShouldBeTrue)
		})

		Convey("8 and 21 are Co-Prime", func() {
			So(relativelyPrime(8, 21), ShouldBeTrue)
		})

		Convey("2 and 4 are NOT Co-Prime", func() {
			So(relativelyPrime(2, 4), ShouldBeFalse)
		})

		Convey("2 and 8 are NOT Co-Prime", func() {
			So(relativelyPrime(2, 8), ShouldBeFalse)
		})

		Convey("3 and 12 are NOT Co-Prime", func() {
			So(relativelyPrime(3, 12), ShouldBeFalse)
		})
	})
}
