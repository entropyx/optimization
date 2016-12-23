package optimization

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// sigmoid
func g(z float64) float64 {
	return 1 / (1 + math.Exp(-z))
}

// hypothesis
func h(x [][]float64, theta []float64) (out []float64) {
	l1 := len(x)
	l2 := len(x[0])
	for i := 0; i < l1; i++ {
		prod := 0.0
		for j := 0; j < l2; j++ {
			prod = prod + x[i][j]*theta[j]
		}
		out = append(out, g(prod))
	}
	return
}

// Cost
func J(theta []float64, y []float64, X [][]float64) float64 {
	m := len(y)
	h := h(X, theta)
	out := 0.00
	for i := 0; i < m; i++ {
		out = out + y[i]*math.Log(h[i]) + (1-y[i])*math.Log(1-h[i])
	}
	out = -out / float64(m)
	return out
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

		Convey("The mean of [1 2 3] is 2", func() {
			out := mean(x)
			So(out, ShouldEqual, 2)
		})

		Convey("The mean of each raw of [[1 2 3] [4 5 6]] is [2 5]", func() {
			x := [][]float64{
				[]float64{1, 2, 3},
				[]float64{4, 5, 6},
			}
			out := apply(x, 1, mean)
			So(out, ShouldResemble, []float64{2, 5})
		})

		Convey("The mean of each column of [[1 2 3] [4 5 6]] is [2.5 3.5 4.5]", func() {
			x := [][]float64{
				[]float64{1, 2, 3},
				[]float64{4, 5, 6},
			}
			out := apply(x, 2, mean)
			So(out, ShouldResemble, []float64{2.5, 3.5, 4.5})
		})

		Convey("Given the following dataset ...", func() {
			var X [][]float64
			var y []float64
			var data [][]float64
			filePath := "/home/gibran/Work/Go/src/github.com/entropyx/optimization/datasets/dataset2.txt"
			strInfo, err := ioutil.ReadFile(filePath)
			if err != nil {
				panic(err)
			}

			trainingData := strings.Split(string(strInfo), "\n")
			for _, line := range trainingData {
				if line == "" {
					break
				}

				var values []float64
				for _, value := range strings.Split(line, " ") {
					floatVal, err := strconv.ParseFloat(value, 64)
					if err != nil {
						panic(err)
					}
					values = append(values, floatVal)
				}
				data = append(data, values)
			}

			for i := 0; i < len(data); i++ {
				X = append(X, data[i][:2])
				y = append(y, data[i][3])
			}

			Convey("The cost of data for theta [0 0 0] is 0.6931471805599458.", func() {
				theta := []float64{0, 0, 0}
				cost := J(theta, y, X)
				So(cost, ShouldEqual, 0.6931471805599458)
			})

			Convey("The minimun ... ", func() {
				//variable := []float64{0, 0, 0}
				//out := neldermead(variable, J)
				//So(out, ShouldResemble, []float64{0.5, 1, 1.5})
			})
		})
	})
}
