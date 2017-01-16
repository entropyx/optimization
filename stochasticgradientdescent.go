package optimization

import ma "github.com/entropyx/math"

// Stochastic Gradient decent implementation methods

// StochasticGradientDecent find local minimum or maximum
func StochasticGradientDecent(fn func([]float64) []float64, p []float64, alpha float64, iter int) (optim []float64) {
	for i := 0; i < iter; i++ {
		alpha = alpha - alpha/float64(iter)
		delta := fn(p)
		for j := 0; j < len(p); j++ {
			delta[j] = alpha * delta[j]
		}
		p = ma.VectorDiff(p, delta)
	}
	optim = p
	return
}
