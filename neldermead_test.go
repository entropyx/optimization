package optimization

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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

			theta := []float64{0, 0, 0}

			Convey("The cost of data for theta [0 0 0] is 0.6931471805599458.", func() {
				cost := J(theta)
				So(cost, ShouldEqual, 0.6931471805599465)
			})

			Convey("The global minimun of cost function is [-25.16133355416168 0.20623171363284806 0.20147159995083574]", func() {
				minimun, cost, iter := Neldermead(theta, J, true)
				fmt.Printf("Cost: %v, Iter: %v \n", cost, iter)
				So(minimun, ShouldResemble, []float64{-22.888963783387748, 0.18806165951719844, 0.18308160110641086})
			})
		})
	})
}
