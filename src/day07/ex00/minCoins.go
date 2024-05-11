package ex00

import (
	"sort"
)

func MinCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}

// MinCoins2 calculates the minimum number of coins needed to cover a given value using certain coin denominations.
//
// This implementation checks the boundary cases and sorts the coins in ascending order before processing them.
//
// Parameters:
//   - val: The value that needs to be covered.
//   - coins: The available denominations of coins.
//
// Returns:
//   - []int: The minimum number of coins needed to make change, represented as a slice of coin values.
func MinCoins2(val int, coins []int) []int {
	if len(coins) == 0 || val == 0 {
		return []int{}
	}

	sort.Sort(sort.IntSlice(coins))

	res := make([]int, 0)
	i := len(coins) - 1

	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}

	return res
}

// Generate documentation:
// To generate HTML documentation, execute the following command in your terminal:
// godoc -http=:6060
// Then open a web browser and go to http://localhost:6060/pkg/<package-name>
// Replace <package-name> with the name of the package containing your code.
