package main

import (
	"fmt"
	"math/rand"
)

type terminationFunc func(testResult) bool

func runSim(startmoney, winchance float64, f terminationFunc) testResult {
	result := testResult{Winnings: startmoney, Numruns: 0}
	for !f(result) {
		if rand.Float64() < winchance {
			// Win!
			result.Winnings++
		} else {
			// Lose :(
			result.Winnings--
		}
		result.Numruns++
	}
	return result
}

func lowerLimit(limit float64) terminationFunc {
	return func(t testResult) bool {
		return t.Winnings <= limit
	}
}

func upperLimit(limit float64) terminationFunc {
	return func(t testResult) bool {
		return t.Winnings >= limit
	}
}

func dualLimit(lower, upper float64) terminationFunc {
	return func(t testResult) bool {
		return t.Winnings <= lower || t.Winnings >= upper
	}
}

type testResult struct {
	Winnings float64
	Numruns  int
}

func runManySims(startmoney, winchance float64, f terminationFunc, n int) (mean testResult) {
	results := make([]testResult, 0, n)
	for i := 0; i < n; i++ {
		result := runSim(startmoney, winchance, f)
		results = append(results, result)
		mean.Winnings += result.Winnings
		mean.Numruns += result.Numruns
	}
	mean.Winnings /= float64(n)
	mean.Numruns /= n

	return
}

func main() {
	// for i := 0; i < 10; i++ {
	fmt.Printf("%+v\n", runManySims(10, 0.5, dualLimit(9, 20), 10000))
	// }
}
