package optimization

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// sigmoid
func g(z float64) float64 {
	return 1 / (1 + math.Exp(-z))
}

// hypothesis
func h(x []float64, theta []float64) float64 {
	l := len(x)
	prod := 0.0
	for i := 0; i < l; i++ {
		prod = prod + x[i]*theta[i]
	}
	return g(prod)
}

// Cost
func J(theta []float64, y []float64, X [][]float64) float64 {
	m := len(y)

}

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
			out := contraction(x, c, beta)
			So(out, ShouldResemble, []float64{0.5, 1, 1.5})
		})

		Convey("The shrink of [1 2 3] respect to [0,0,0] is [0.5,1,1.5]", func() {
			delta := 0.5
			out := shrink(x, c, delta)
			So(out, ShouldResemble, []float64{0.5, 1, 1.5})
		})

		Convey("The minimun ... ", func() {
			variable := []float64{0, 0, 0}
			out := neldermead(variable, J)
			So(out, ShouldResemble, []float64{0.5, 1, 1.5})
		})

	})
}
