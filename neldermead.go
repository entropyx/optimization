package optimization

import (
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
	radius := 0.1
	for i := 0; i < n; i++ {
		var p []float64
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
		for i := 0; i < len(t); i++ {
			out = append(out, f(t[i]))
		}
	case n > 2 || n < 1:
		panic("n must be 1 or 2!.")
	}
	return out
}

func order(x []float64, decreasing bool) (out []int) {
	l := len(x)
	k := 0
	for i := 0; i < l; i++ {
		out = append(out, i+1)
	}
	switch decreasing {
	case true:
		for k < l-1 {
			if x[k] < x[k+1] {
				x[k], x[k+1] = x[k+1], x[k]
				out[k], out[k+1] = out[k+1], out[k]
				k = 0
			} else {
				k++
			}
		}
	case false:
		for k < l-1 {
			if x[k] > x[k+1] {
				x[k], x[k+1] = x[k+1], x[k]
				out[k], out[k+1] = out[k+1], out[k]
				k = 0
			} else {
				k++
			}
		}
	}
	return
}

func sort(x [][]float64, by []int) (out [][]float64) {
	for _, i := range by {
		out = append(out, x[i-1])
	}
	return
}

func neldermead(variables parameter, fn function, iter int) (center []float64, cost float64) {
	coord := make()
	n := len(variables.theta)
	p := append(around(variables.theta, n+1), variables.theta)
	for i := 0; i < iter; i++ {
		var f []float64
		for i := 0; i < len(p); i++ {
			variables.theta = p[i]
			f = append(f, fn(variables))
		}
		order := order(f, true)
		p = sort(p, order)
		xh := p[1]
		xs := p[2]
		xl := p[len(p)-1]
		c := apply(p[1:], 2, mean)
		xr := reflection(xh, c, 1)
		xc := contraction(xh, c, 0.5)
		if fn(xr) >= fn(xl) && fn(xr) < fn(xs) {
			p[1] = xr
		} else if fn(xr) < fn(xl) {
			xb := xr
			gamma := 1
			bool := true
			for bool == true {
				gamma++
				xe := expansion(xr, c, gamma)
				if fn(xe) > fn(xr) {
					p[1] = xb
					bool = false
				} else {
					xb = xe
				}
			}
		} else if fn(xr) > fn(xh) && fn(xc) < fn(xh) {
			p[1] = xc
		} else {
			for i := 0; i < len(p)-1; i++ {
				p[i] = shrink(xl, p[i], 0.5)
			}
		}
	}
	center = apply(p, 2, mean)
	cost = fn(c)
	return
}
