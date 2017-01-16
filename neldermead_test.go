package optimization

import (
	"fmt"
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

		Convey("The expansion of [1 2 3] is [2,4,6]", func() {
			gamma := 2.00
			out := expansion(x, c, gamma)
			So(out, ShouldResemble, []float64{2, 4, 6})
		})

		Convey("The contraction of [1 2 3] is [0.5,1,1.5]", func() {
			beta := 0.5
			ci, _ := contraction(x, c, beta)
			So(ci, ShouldResemble, []float64{0.5, 1, 1.5})
		})

		Convey("The shrink of [1 2 3] respect to [0,0,0] is [0.5,1,1.5]", func() {
			delta := 0.5
			out := shrink(x, c, delta)
			So(out, ShouldResemble, []float64{0.5, 1, 1.5})
		})

		Convey("The mean of [1 2 3] is 2", func() {
			out := mean(x)
			So(out, ShouldEqual, 2)
		})

		Convey("Given the following dataset ...", func() {

			theta := []float64{0, 0, 0}

			Convey("The cost of data for theta [0 0 0] is 0.6931471805599458.", func() {
				cost := J(theta)
				So(cost, ShouldEqual, 0.6931471805599458)
			})

			Convey("The global minimun of cost function is [-25.19463710955625 0.20646950187695884 0.2017454240918573]", func() {
				minimum, cost, iter := Neldermead(theta, J, true)
				fmt.Printf("Minimum: %v, Cost: %v, Iter: %v \n", minimum, cost, iter)
				So(minimum, ShouldResemble, []float64{1.717559043890748, 39.93914541231272, 37.26620779876064})
			})

		})
	})
}
