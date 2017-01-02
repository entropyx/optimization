package optimization

import (
	"fmt"
	"reflect"

	"github.com/entropyx/tools"
)

type coord struct {
	xr, xs, xl, xe, xci, xco, xh float64
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

func minimize(fn interface{}, p [][]float64) (center []float64, height float64, iter int) {
	var xh, xl, c, xr, xci, xco, xe []float64
	var z coord
	n := len(p[0])
	beta := 0.75 - 1/(2*float64(n))
	gamma := 2.0 + 2.0/float64(n)
	delta := 1.0 - 1/float64(n)
	wse := 1.00
	iter = 0

	f := reflect.ValueOf(fn)
	fnType := f.Type()
	if fnType.Kind() != reflect.Func || fnType.NumIn() != 1 || fnType.NumOut() != 1 {
		panic("Expected a unary function returning a single value")
	}

	for wse > 1e-3 {
		iter++
		fmt.Printf("Iter %v \n", iter)
		var f2 []float64
		for j := 0; j < len(p); j++ {
			f2 = append(f2, f.Call([]reflect.Value{reflect.ValueOf(p[j])})[0].Float())
		}
		order := tools.Order(f2, true)
		p = tools.Sort(p, order)
		xh = p[0]
		xl = p[len(p)-1]
		c = tools.Apply(p[1:], 2, mean)
		xr = reflection(xh, c, 1)
		xci, xco = contraction(xh, c, beta)
		z.xr = f.Call([]reflect.Value{reflect.ValueOf(xr)})[0].Float()
		z.xci = f.Call([]reflect.Value{reflect.ValueOf(xci)})[0].Float()
		z.xco = f.Call([]reflect.Value{reflect.ValueOf(xco)})[0].Float()
		z.xl = f2[len(f2)-1]
		z.xs = f2[1]
		z.xh = f2[0]
		if z.xr >= z.xl && z.xr < z.xs {
			fmt.Println("Reflect")
			p[0] = xr
		} else if z.xr < z.xl {
			p[0] = xr
			xe = expansion(xr, c, gamma)
			z.xe = f.Call([]reflect.Value{reflect.ValueOf(xe)})[0].Float()
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
		height = f.Call([]reflect.Value{reflect.ValueOf(xl)})[0].Float()
		fmt.Println(center, height)
		wse = 0
		c = tools.Apply(p, 2, mean)
		for j := 0; j < len(p); j++ {
			wse = wse + tools.Dist(c, p[j])
		}
	}
	return
}

func maximize(fn interface{}, p [][]float64) (center []float64, height float64, iter int) {
	var xh, xl, c, xr, xci, xco, xe []float64
	var z coord
	n := len(p[0])
	beta := 0.75 - 1/(2*float64(n))
	gamma := 2.0 + 2.0/float64(n)
	delta := 1.0 - 1/float64(n)
	wse := 1.00
	iter = 0

	f := reflect.ValueOf(fn)
	fnType := f.Type()
	if fnType.Kind() != reflect.Func || fnType.NumIn() != 1 || fnType.NumOut() != 1 {
		panic("Expected a unary function returning a single value")
	}

	for wse > 1e-3 {
		iter++
		fmt.Printf("Iter %v \n", iter)
		var f2 []float64
		for j := 0; j < len(p); j++ {
			f2 = append(f2, f.Call([]reflect.Value{reflect.ValueOf(p[j])})[0].Float())
		}
		order := tools.Order(f2, false)
		p = tools.Sort(p, order)
		xl = p[0]
		xh = p[len(p)-1]
		c = tools.Apply(p[1:], 2, mean)
		xr = reflection(xl, c, 1)
		xci, xco = contraction(xl, c, beta)
		z.xr = f.Call([]reflect.Value{reflect.ValueOf(xr)})[0].Float()
		z.xci = f.Call([]reflect.Value{reflect.ValueOf(xci)})[0].Float()
		z.xco = f.Call([]reflect.Value{reflect.ValueOf(xco)})[0].Float()
		z.xh = f2[len(f2)-1]
		z.xs = f2[1]
		z.xl = f2[0]
		if z.xr <= z.xh && z.xr > z.xs {
			fmt.Println("Reflect")
			p[0] = xr
		} else if z.xr > z.xh {
			p[0] = xr
			xe = expansion(xr, c, gamma)
			z.xe = f.Call([]reflect.Value{reflect.ValueOf(xe)})[0].Float()
			fmt.Printf("f(xe) = %v \n", z.xe)
			if z.xe > z.xr {
				fmt.Println("Expand")
				p[0] = xe
			}
		} else if z.xci > z.xl || z.xco > z.xl {
			if z.xci > z.xco {
				fmt.Println("Contract inside")
				p[0] = xci
			} else {
				fmt.Println("Contract outside")
				p[0] = xco
			}
		} else {
			fmt.Println("Shrink")
			for i := 0; i < len(p)-1; i++ {
				p[i] = shrink(xh, p[i], delta)
			}
		}
		center = xh
		height = f.Call([]reflect.Value{reflect.ValueOf(xh)})[0].Float()
		fmt.Println(center, height)
		wse = 0
		c = tools.Apply(p, 2, mean)
		for j := 0; j < len(p); j++ {
			wse = wse + tools.Dist(c, p[j])
		}
	}
	return
}

// Neldermead maximize o minimize
func Neldermead(par []float64, fn interface{}, minimum bool) (center []float64, height float64, iter int) {
	n := len(par)
	p := around(par, n)
	p = append(p, par)
	switch minimum {
	// Minimize
	case true:
		center, height, iter = minimize(fn, p)
	// Maximize
	case false:
		center, height, iter = maximize(fn, p)
	}
	return
}
