package knapsack

type Present struct {
	Value int
	Size  int
}

func grabPresents(presents []Present, capacity int) []Present {
	n := len(presents)

	table := make([][]int, n+1)

	for i := range table {
		table[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			if presents[i-1].Size <= w {
				table[i][w] = max(presents[i-1].Value+table[i-1][w-presents[i-1].Size], table[i-1][w])
			} else {
				table[i][w] = table[i-1][w]
			}
		}
	}

	var chosenPresents []Present
	w := capacity
	for i := n; i > 0 && w > 0; i-- {
		if table[i][w] != table[i-1][w] {
			chosenPresents = append(chosenPresents, presents[i-1])
			w -= presents[i-1].Size
		}
	}

	if len(chosenPresents) == 0 {
		return []Present{}
	}

	return chosenPresents
}
