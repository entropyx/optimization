package optimization

import (
	"fmt"
	"math"
	"math/rand"
)

type parameter struct {
	theta, y []float64
	X        [][]float64
}

type fn func([]float64) float64

type function func(parameter) float64

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

func mean(x []float64) float64 {
	out := 0.00
	n := len(x)
	for i := 0; i < n; i++ {
		out = out + x[i]
	}
	out = out / float64(n)
	return out
}

func apply(X [][]float64, n int, f fn) (out []float64) {
	switch {
	// apply by row
	case n == 1:
		for i := 0; i < len(X); i++ {
			out = append(out, f(X[i]))
		}
		// apply by column
	case n == 2:
		var t [][]float64
		for i := 0; i < len(X[0]); i++ {
			var column []float64
			for j := 0; j < len(X); j++ {
				column = append(column, X[j][i])
			}
			t = append(t, column)
		}
		fmt.Println(t)
		for i := 0; i < len(t); i++ {
			out = append(out, f(t[i]))
		}
	case n > 2 || n < 1:
		panic("n must be 1 or 2!.")
	}
	return out
}

func neldermead(variables parameter, fn function, iter int) {
	n := len(variables.theta)
	p := append(around(variables.theta, n+1), variable)
	fmt.Println(p)
	for i := 0; i < iter; i++ {
		//var f []float64
		fmt.Println(fn(variables))

	}

}
