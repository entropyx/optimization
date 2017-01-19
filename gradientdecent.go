package optimization

import ma "github.com/entropyx/math"

// Gradient decent implementation methods

// GradientChecking to stop the iteration.
func GradientChecking(fn func([]float64) float64, theta []float64, epsilon float64) float64 {
	var x, y []float64
	for i := 0; i < len(theta); i++ {
		x = append(x, theta[i]+epsilon)
		y = append(y, theta[i]-epsilon)
	}
	checking := (fn(x) - fn(y)) / 2 * epsilon
	return checking
}

// GradientDecent find local minimum
func GradientDecent(gradient func([]float64) []float64, p []float64, alpha float64, iter int) (optim []float64) {
	for i := 0; i < iter; i++ {
		delta := gradient(p)
		for j := 0; j < len(p); j++ {
			delta[j] = alpha * delta[j]
		}
		p = ma.VectorDiff(p, delta)
	}
	optim = p
	return
}
