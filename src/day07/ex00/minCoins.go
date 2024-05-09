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
