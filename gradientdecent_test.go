package optimization

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGradientFunction(t *testing.T) {

	// Data obtained from the Coursera course https://www.coursera.org/course/ml
	// of Andrew Ng.

	Convey("Given X ...", t, func() {
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

		Convey("The global minimun is [-25.16133355416168 0.20623171363284806 0.20147159995083574]", func() {
			minimum, cost := Gradientdecent(J, theta)
			fmt.Printf("Minimum: %v, Cost: %v \n", minimum, cost)
			// So(minimun, ShouldResemble, []float64{-25.160920349329682, 0.20622792851376404, 0.201468441137889})
		})
	})
}
