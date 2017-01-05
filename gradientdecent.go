package optimization

import (
	"math"
	"reflect"

	ma "github.com/entropyx/math"
)

// Gradient decent implementation methods

// G sigmoid
func G(z float64) float64 {
	return 1 / (1 + math.Exp(-z))
}

//H hypothesis
func H(x [][]float64, theta []float64) (out []float64) {
	l1 := len(x)
	Theta := [][]float64{theta}
	c := make(chan [][]float64)
	go ma.MatrixProduct(x, Theta, c)
	prod := <-c
	for i := 0; i < l1; i++ {
		out = append(out, G(prod[i][0]))
	}
	return
}

var y []float64
var X [][]float64

// J Cost function
func J(theta []float64) float64 {
	penal := 0.00
	lambda := 0.00
	m := len(y)
	h := H(X, theta)
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

func grad(X [][]float64, y []float64, theta []float64, lambda float64) []float64 {
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

// Gradientdecent find local minimum or maximum
func Gradientdecent(fn interface{}, p []float64) (optim []float64, height float64) {
	iter := 100
	alpha := 0.001
	lambda := 0.00
	f := reflect.ValueOf(fn)
	fnType := f.Type()
	if fnType.Kind() != reflect.Func || fnType.NumIn() != 1 || fnType.NumOut() != 1 {
		panic("Expected a unary function returning a single value")
	}

	for i := 0; i < iter; i++ {
		delta := grad(X, y, p, lambda)
		for j := 0; j < len(p); j++ {
			delta[j] = alpha * delta[j]
		}
		p = ma.VectorDiff(p, delta)
	}
	optim = p
	height = f.Call([]reflect.Value{reflect.ValueOf(optim)})[0].Float()
	return
}
