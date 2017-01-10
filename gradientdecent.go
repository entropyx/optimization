package optimization

// import (
// 	"reflect"
//
// 	ma "github.com/entropyx/math"
// )
//
// // Gradient decent implementation methods
//
// // GradientDecent find local minimum or maximum
// func GradientDecent(fn interface{}, p []float64) (optim []float64) {
// 	iter := 10000
// 	alpha := 1.0
// 	f := reflect.ValueOf(fn)
// 	fnType := f.Type()
// 	if fnType.Kind() != reflect.Func || fnType.NumIn() != 1 || fnType.NumOut() != 1 {
// 		panic("Expected a unary function returning a single value")
// 	}
//
// 	for i := 0; i < iter; i++ {
// 		delta := f.CallSlice([]reflect.Value{reflect.ValueOf(p)})[1]
// 		for j := 0; j < len(p); j++ {
// 			delta[j] = alpha * delta[j]
// 		}
// 		p = ma.VectorDiff(p, delta)
// 	}
// 	optim = p
// 	return
// }
