package optimization

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
