package optimization

import (
	"math"

	ma "github.com/entropyx/math"
	//. "github.com/smartystreets/goconvey/convey"
)

// G sigmoid
func G(z float64) float64 {
	return 1 / (1 + math.Exp(-z))
}

var y []float64
var X [][]float64

//H hypothesis
func H(theta []float64) (out []float64) {
	l1 := len(X)
	c := make(chan []float64)
	go ma.MatrixVectorProduct(X, theta, c)
	prod := <-c
	for i := 0; i < l1; i++ {
		out = append(out, G(prod[i]))
	}
	return
}

const lambda = 0.00

// J Cost function
func J(theta []float64) float64 {
	penal := 0.00
	m := len(y)
	h := H(theta)
	out := 0.00
	for i := 0; i < len(theta); i++ {
		penal = penal + lambda*math.Pow(theta[i], 2)
	}
	for i := 0; i < m; i++ {
		out = out + y[i]*math.Log(h[i]) + (1-y[i])*math.Log(1-h[i])
	}
	out = -out/float64(m) + penal/float64(2*m)
	return out
}

func Grad(theta []float64) []float64 {
	var penal, gradient2 []float64
	m := len(y)
	if lambda != 0.00 {
		for i := 0; i < len(theta); i++ {
			penal = append(penal, (lambda/float64(m))*theta[i])
		}
		penal[0] = 1.00
	} else {
		penal = make([]float64, len(theta))
	}
	error := ma.VectorDiff(H(X, theta), y)
	c := make(chan [][]float64)
	go ma.MatrixProduct(ma.Traspose(X), [][]float64{error}, c)
	gradient := <-c
	for i := 0; i < len(gradient); i++ {
		gradient2 = append(gradient2, (1/float64(m))*gradient[i][0])
	}
	gradient2 = ma.VectorAdd(gradient2, penal)
	return gradient2
}

// func TestGradientFunction(t *testing.T) {
//
// 	Data obtained from the Coursera course https://www.coursera.org/course/ml
// 	of Andrew Ng.
//
// 	Convey("Given X ...", t, func() {
// 		var data [][]float64
// 		filePath := "/home/gibran/Work/Go/src/github.com/entropyx/optimization/datasets/dataset2.txt"
// 		strInfo, err := ioutil.ReadFile(filePath)
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		trainingData := strings.Split(string(strInfo), "\n")
// 		for _, line := range trainingData {
// 			if line == "" {
// 				break
// 			}
//
// 			var values []float64
// 			for _, value := range strings.Split(line, " ") {
// 				floatVal, err := strconv.ParseFloat(value, 64)
// 				if err != nil {
// 					panic(err)
// 				}
// 				values = append(values, floatVal)
// 			}
// 			data = append(data, values)
// 		}
//
// 		for i := 0; i < len(data); i++ {
// 			X = append(X, data[i][1:3])
// 			y = append(y, data[i][3])
// 		}
// 		X2 := make([][]float64, len(X))
// 		copy(X2, X)
// 		mu := tools.Apply(X2, 2, ma.Mean)
// 		sigma := tools.Apply(X2, 2, ma.Sd)
//
// 		X = ma.Normalize(X)
//
// 		for i := 0; i < len(X); i++ {
// 			X[i] = append([]float64{1}, X[i]...)
// 		}
//
// 		theta := []float64{0, 0, 0}
//
// 		Convey("The global minimun is [-25.16133355416168 0.20623171363284806 0.20147159995083574]", func() {
// 			minimum := GradientDecent(Grad, theta)
// 			minimum = ma.RescaleCoef(minimum, mu, sigma)
// 			fmt.Printf("Minimum: %v \n", minimum)
// 			// So(minimun, ShouldResemble, []float64{-25.160920349329682, 0.20622792851376404, 0.201468441137889})
// 		})
// 	})
// }
