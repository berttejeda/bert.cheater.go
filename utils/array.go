package utils

import (
	"fmt"
)

// Function to generate all permutations of a slice of strings
func ArrayPermute(arr []string, l, r int, results *[][]string) {
	if l == r {
		// If the left index is equal to the right index, it means we have a valid permutation
		// Append the current permutation to the results
		perm := make([]string, len(arr))
		copy(perm, arr)
		*results = append(*results, perm)
	} else {
		for i := l; i <= r; i++ {
			// Swap the elements at indices l and i
			arr[l], arr[i] = arr[i], arr[l]
			// Recursively generate permutations for the remaining elements
			ArrayPermute(arr, l+1, r, results)
			// Backtrack to restore the original configuration
			arr[l], arr[i] = arr[i], arr[l]
		}
	}
}

func ArrayFormatEachItem(slice []string) []string {
    for i := range slice {
        (slice)[i] = fmt.Sprintf(".%s$", (slice)[i])
    }
    return slice
}

func ArrayContains(slice []string, item string) bool {
    for _, str := range slice {
        if str == item {
            return true
        }
    }
    return false
}

