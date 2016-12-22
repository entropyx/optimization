package optimization

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNelderMead(t *testing.T) {

	Convey("Given the vector [1 2 3] and the centroid [0 0 0]", t, func() {
		x := []float64{1, 2, 3}
		c := []float64{0, 0, 0}
		Convey("The reflection of x is [-1 -2 -3]", func() {
			alpha := 1.00
			out := reflection(x, c, alpha)
			So(out, ShouldResemble, []float64{-1, -2, -3})
		})
	})
}
