package optimization

import ma "github.com/entropyx/math"

// Gradient decent implementation methods

// GradientDecent find local minimum or maximum
func GradientDecent(fn func([]float64) []float64, p []float64) (optim []float64) {
	iter := 100000
	alpha := 1.0

	for i := 0; i < iter; i++ {
		delta := fn(p)
		for j := 0; j < len(p); j++ {
			delta[j] = alpha * delta[j]
		}
		p = ma.VectorDiff(p, delta)
	}
	optim = p
	return
}
