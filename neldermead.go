package optimization

import (
	"fmt"
	"math"
)

type Parameter struct {
	Variable []float64
	Y        []int
	X        [][]float64
}

type coord struct {
	xr, xs, xl, xe, xci, xco, xh, xb float64
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

func contraction(x []float64, c []float64, beta float64) (ic []float64, oc []float64) {
	l := len(x)
	for i := 0; i < l; i++ {
		ic = append(ic, c[i]+beta*(x[i]-c[i]))
		oc = append(oc, c[i]-beta*(x[i]-c[i]))
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
	delta := []float64{0.1, -0.1}
	for i := 0; i < n; i++ {
		p := make([]float64, len(c))
		copy(p, c)
		if (i % 2) == 0 {
			p[i] = p[i] + delta[0]
		} else {
			p[i] = p[i] + delta[1]
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
func Neldermead(Variable string, parameter Parameter, fn function, minimize bool) (center []float64, cost float64, iter int) {
	var z coord
	var xh, xs, xl, c, xr, xci, xco, xe, xb []float64
	wse := 1.00
	iter = 0
	n := len(parameter.Variable)
	beta := 0.75 - 1/(2*float64(n))
	gamma := 2.0 + 2.0/float64(n)
	delta := 1.0 - 1/float64(n)
	p := around(parameter.Variable, n)
	p = append(p, parameter.Variable)
	switch minimize {
	// Minimize
	case true:
		for wse > 1e-3 {
			iter++
			fmt.Printf("Iter %v \n", iter)
			var f []float64
			for j := 0; j < len(p); j++ {
				parameter.Variable = p[j]
				f = append(f, fn(parameter))
			}
			order := order(f, true)
			p = sort(p, order)
			xh = p[0]
			xs = p[1]
			xl = p[len(p)-1]
			c = apply(p[1:], 2, mean)
			xr = reflection(xh, c, 1)
			xci, xco = contraction(xh, c, beta)
			parameter.Variable = xr
			z.xr = fn(parameter)
			parameter.Variable = xci
			z.xci = fn(parameter)
			parameter.Variable = xco
			z.xco = fn(parameter)
			z.xl = f[len(f)-1]
			z.xs = f[1]
			z.xh = f[0]
			if z.xr >= z.xl && z.xr < z.xs {
				fmt.Println("Reflect")
				p[0] = xr
			} else if z.xr < z.xl {
				p[0] = xr
				xe = expansion(xr, c, gamma)
				parameter.Variable = xe
				z.xe = fn(parameter)
				fmt.Printf("f(xe) = %v \n", z.xe)
				if z.xe < z.xr {
					fmt.Println("Expand")
					p[0] = xe
				}
			} else if z.xci < z.xh || z.xco < z.xh {
				if z.xci < z.xco {
					fmt.Println("Contract inside")
					p[0] = xci
				} else {
					fmt.Println("Contract outside")
					p[0] = xco
				}
			} else {
				fmt.Println("Shrink")
				for i := 0; i < len(p)-1; i++ {
					p[i] = shrink(xl, p[i], delta)
				}
			}
			center = xl
			parameter.Variable = center
			cost = fn(parameter)
			fmt.Println(center, cost)
			wse = 0
			c = apply(p, 2, mean)
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
				parameter.Variable = p[j]
				f = append(f, fn(parameter))
			}
			order := order(f, false)
			p = sort(p, order)
			xl = p[0]
			xs = p[1]
			xh = p[len(p)-1]
			c = apply(p[1:], 2, mean)
			xr = reflection(xl, c, 1)
			xci, xco = contraction(xl, c, 0.5)
			parameter.Variable = xr
			z.xr = fn(parameter)
			parameter.Variable = xh
			z.xh = fn(parameter)
			parameter.Variable = xs
			z.xs = fn(parameter)
			parameter.Variable = xci
			z.xci = fn(parameter)
			parameter.Variable = xco
			z.xco = fn(parameter)
			parameter.Variable = xl
			z.xl = fn(parameter)
			if z.xr <= z.xh && z.xr > z.xs {
				p[0] = xr
			} else if z.xr > z.xh {
				xb = xr
				gamma := 1.00
				bool := true
				for bool == true {
					gamma++
					xe = expansion(xr, c, gamma)
					parameter.Variable = xe
					z.xe = fn(parameter)
					if z.xe <= z.xr {
						p[0] = xb
						bool = false
					} else {
						xb = xe
					}
				}
			} else if z.xr < z.xl && z.xci > z.xl {
				p[0] = xci
			} else {
				for i := 0; i < len(p)-1; i++ {
					p[i] = shrink(xh, p[i], 0.5)
				}
			}
			center = apply(p, 2, mean)
			parameter.Variable = center
			cost = fn(parameter)
			wse = 0
			for j := 0; j < len(p); j++ {
				wse = wse + distance(c, p[j])
			}
		}
	}
	return
}
