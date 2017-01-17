package optimization

import (
	"fmt"
	"math/rand"
	"testing"

	ma "github.com/entropyx/math"
	. "github.com/smartystreets/goconvey/convey"
)

var lambda2 = 0.0001

//H hypothesis
func H2(Xi, theta []float64) float64 {
	prod := ma.VectorProduct(Xi, theta)
	return G(prod)
}

func Grad2(theta []float64) (grad []float64) {
	m := len(y)
	sample := rand.Intn(m)
	Xi := X[sample]
	yi := y[sample]
	error := H2(Xi, theta) - yi
	for j := 0; j < len(Xi); j++ {
		grad = append(grad, Xi[j]*error-lambda2*theta[j])
	}
	return
}

func TestStochasticGradientDecentFunction(t *testing.T) {
	Convey("The stochastic gradient decent ...", t, func() {
		theta := []float64{0, 0, 0}

		Convey("... must converge to [-25.16133355416168 0.20623171363284806 0.20147159995083574]", func() {
			minimum := StochasticGradientDecent(Grad2, theta, 0.1, 100000)
			minimum = ma.RescaleCoef(minimum, mu, sigma)
			fmt.Printf("Stochastic Gradient Decent Minimum: %v \n", minimum)
			So(minimum, ShouldResemble, []float64{-23.663158419526443, 0.19485272382138286, 0.18867193554561834})
		})
	})
}
