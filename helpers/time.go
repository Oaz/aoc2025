package helpers

import "time"

func TimeFunction(fn func(), iterations int) time.Duration {
	if iterations <= 0 {
		iterations = 100
	}

	start := time.Now()
	for i := 0; i < iterations; i++ {
		fn()
	}
	elapsed := time.Since(start)

	return elapsed / time.Duration(iterations)
}
