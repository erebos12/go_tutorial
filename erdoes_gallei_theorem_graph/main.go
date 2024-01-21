package main

import (
	"fmt"
	"sort"
)

// Erdős-Gallai-Theorem:
// Überprüft, ob die gegebene Sequenz eine gültige Gradsequenz ist
func isValidDegreeSequence(degrees []int) bool {
	n := len(degrees)
	sum := 0

	// Überprüfen Sie, ob die Summe der Grade gerade ist
	for _, d := range degrees {
		sum += d
	}
	if sum%2 != 0 {
		return false
	}

	// Sortieren der Grade in absteigender Reihenfolge
	sort.Sort(sort.Reverse(sort.IntSlice(degrees)))

	for k := 1; k <= n; k++ {
		leftSum := 0
		for i := 0; i < k; i++ {
			leftSum += degrees[i]
		}
		rightSum := k * (k - 1)
		for i := k; i < n; i++ {
			min(k, degrees[i])
		}

		if leftSum > rightSum {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Beispiel für eine Gradsequenz
	degrees := []int{4, 3, 3, 3, 2}

	if isValidDegreeSequence(degrees) {
		fmt.Println("Die Sequenz ist eine gültige Gradsequenz.")
	} else {
		fmt.Println("Die Sequenz ist keine gültige Gradsequenz.")
	}
}
