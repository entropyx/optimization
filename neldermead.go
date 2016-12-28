package optimization

import (
	"math"
	"math/rand"
)

type Parameter struct {
	Y     []int
	Theta []float64
	X     [][]float64
}

type coord struct {
	xr, xs, xl, xe, xc, xh float64
}

type fn func([]float64) float64

type function func(Parameter) float64

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
	radius := 5.00
	for i := 0; i < n; i++ {
		var p []float64
		for i := 0; i < len(c); i++ {
			degrees := float64(rand.Intn(360))
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

func distance(x1, x2 []float64) (dist float64) {
	l1 := len(x1)
	l2 := len(x2)
	s := 0.00
	if l1 == l2 {
		for i := 0; i < l1; i++ {
			s = s + math.Pow(x1[i]-x2[i], 2)
		}
		dist = math.Sqrt(s)
	} else {
		panic("The vector hasn't the same length!.")
	}
	return
}

// Neldermead maximize o minimize
func Neldermead(variables Parameter, fn function, minimize bool) (center []float64, cost float64, iter int) {
	var z coord
	var xh, xs, xl, c, xr, xc, xe, xb []float64
	wse := 1.00
	iter = 0
	n := len(variables.Theta)
	p := append(around(variables.Theta, n), variables.Theta)
	switch minimize {
	// Minimize
	case true:
		for wse > 1e-10 {
			iter++
			var f []float64
			for j := 0; j < len(p); j++ {
				variables.Theta = p[j]
				f = append(f, fn(variables))
			}
			order := order(f, true)
			p = sort(p, order)
			xh = p[0]
			xs = p[1]
			xl = p[len(p)-1]
			c = apply(p[1:], 2, mean)
			xr = reflection(xh, c, 1)
			xc = contraction(xh, c, 0.5)
			variables.Theta = xr
			z.xr = fn(variables)
			variables.Theta = xl
			z.xl = fn(variables)
			variables.Theta = xs
			z.xs = fn(variables)
			variables.Theta = xc
			z.xc = fn(variables)
			variables.Theta = xh
			z.xh = fn(variables)
			if z.xr >= z.xl && z.xr < z.xs {
				p[0] = xr
			} else if z.xr < z.xl {
				xb = xr
				gamma := 1.00
				bool := true
				for bool == true {
					gamma++
					xe = expansion(xr, c, gamma)
					variables.Theta = xe
					z.xe = fn(variables)
					if z.xe >= z.xr {
						p[0] = xb
						bool = false
					} else {
						xb = xe
					}
				}
			} else if z.xr > z.xh && z.xc < z.xh {
				p[0] = xc
			} else {
				for i := 0; i < len(p); i++ {
					p[i] = shrink(xl, p[i], 0.5)
				}
			}
			center = apply(p, 2, mean)
			variables.Theta = center
			cost = fn(variables)
			wse = 0
			for j := 0; j < len(p); j++ {
				wse = wse + distance(c, p[j])
			}
		}
		// Maximize
	case false:
		for wse > 1e-10 {
			iter++
			var f []float64
			for j := 0; j < len(p); j++ {
				variables.Theta = p[j]
				f = append(f, fn(variables))
			}
			order := order(f, false)
			p = sort(p, order)
			xl = p[0]
			xs = p[1]
			xh = p[len(p)-1]
			c = apply(p[1:], 2, mean)
			xr = reflection(xl, c, 1)
			xc = contraction(xl, c, 0.5)
			variables.Theta = xr
			z.xr = fn(variables)
			variables.Theta = xh
			z.xh = fn(variables)
			variables.Theta = xs
			z.xs = fn(variables)
			variables.Theta = xc
			z.xc = fn(variables)
			variables.Theta = xl
			z.xl = fn(variables)
			if z.xr <= z.xh && z.xr > z.xs {
				p[0] = xr
			} else if z.xr > z.xh {
				xb = xr
				gamma := 1.00
				bool := true
				for bool == true {
					gamma++
					xe = expansion(xr, c, gamma)
					variables.Theta = xe
					z.xe = fn(variables)
					if z.xe <= z.xr {
						p[0] = xb
						bool = false
					} else {
						xb = xe
					}
				}
			} else if z.xr < z.xl && z.xc > z.xl {
				p[0] = xc
			} else {
				for i := 0; i < len(p); i++ {
					p[i] = shrink(xh, p[i], 0.5)
				}
			}
			center = apply(p, 2, mean)
			variables.Theta = center
			cost = fn(variables)
			wse = 0
			for j := 0; j < len(p); j++ {
				wse = wse + distance(c, p[j])
			}
		}
	}
	return
}
