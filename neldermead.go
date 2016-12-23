package optimization

import (
	"fmt"
	"math"
	"math/rand"
)

type function interface {
	f() []float64
}

func reflection(x []float64, c []float64, alpha float64) (out []float64) {
	l := len(x)
	for i := 0; i < l; i++ {
		out = append(out, c[i]+alpha*(c[i]-x[i]))
	}
	return
}

func expansion(x []float64, c []float64, gamma float64) (out []float64) {
	l := len(x)
	for i := 0; i < l; i++ {
		out = append(out, c[i]+gamma*(x[i]-c[i]))
	}
	return
}

func contraction(x []float64, c []float64, beta float64) (out []float64) {
	l := len(x)
	for i := 0; i < l; i++ {
		out = append(out, c[i]+beta*(x[i]-c[i]))
	}
	return
}

func shrink(x []float64, y []float64, delta float64) (out []float64) {
	l := len(x)
	for i := 0; i < l; i++ {
		out = append(out, x[i]+delta*(y[i]-x[i]))
	}
	return
}

func around(c []float64, n int) (out [][]float64) {
	var p []float64
	radius := 0.1
	for i := 0; i < n; i++ {
		degrees := float64(rand.Intn(360))
		for i := 0; i < len(c); i++ {
			p = append(p, c[i]+radius*math.Cos(degrees*math.Pi/180.00))
		}
		out = append(out, p)
	}
	return
}

func neldermead(variable []float64, fn function, iter int) {
	n := len(variable)
	p := append(variable, around(variable, n+1))
	fmt.Println(p)
	for i := 0; i < iter; i++ {
		var f []float64

	}

}
