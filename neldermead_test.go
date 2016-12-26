package optimization

import (
	"fmt"
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
func J(p parameter) float64 {
	m := len(p.y)
	h := h(p.X, p.theta)
	out := 0.00
	for i := 0; i < m; i++ {
		out = out + p.y[i]*math.Log(h[i]) + (1-p.y[i])*math.Log(1-h[i])
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

		x = []float64{3, 2.5, 3.5, 1, 4.5}
		order := order(x, true)
		Convey("The decreasing order of [3 2.5 3.5 1 4.5] is [5 3 1 2 4]", func() {
			So(order, ShouldResemble, []int{5, 3, 1, 2, 4})
		})

		Convey("The decreasing sort of [3 2.5 3.5 1 4.5] is [4.5 3.5 3 2.5 1]", func() {
			y := [][]float64{
				[]float64{1, 1, 1},
				[]float64{2, 2, 2},
				[]float64{3, 3, 3},
				[]float64{4, 4, 4},
				[]float64{5, 5, 5},
			}
			sort := sort(y, order)
			So(sort[0], ShouldResemble, []float64{5, 5, 5})
		})

		Convey("The euclidean diastance between [1 1 1] [-1 -1 -1]] is 3.4641016151377544", func() {
			x := []float64{1, 1, 1}
			y := []float64{-1, -1, -1}
			dist := distance(x, y)
			So(dist, ShouldEqual, 3.4641016151377544)
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
				X = append(X, data[i][:3])
				y = append(y, data[i][3])
			}

			var par parameter
			par.theta = []float64{0, 0, 0}
			par.y = y
			par.X = X

			Convey("The cost of data for theta [0 0 0] is 0.6931471805599458.", func() {
				cost := J(par)
				So(cost, ShouldEqual, 0.6931471805599458)
			})

			Convey("The global minimun of cost function is [-25.16133355416168 0.20623171363284806 0.20147159995083574]", func() {

				minimun, cost, iter := neldermead(par, J)
				fmt.Printf("Cost: %v, Iter: %v \n", cost, iter)
				So(minimun, ShouldResemble, []float64{-25.16133355416168, 0.20623171363284806, 0.20147159995083574})
			})
		})
	})
}
