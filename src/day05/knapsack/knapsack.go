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
		for j := 1; j <= capacity; j++ {
			if presents[i-1].Size <= j {
				table[i][j] = max(presents[i-1].Value+table[i-1][j-presents[i-1].Size], table[i-1][j])
			} else {
				table[i][j] = table[i-1][j]
			}
		}
	}

	var chosenPresents []Present
	c := capacity
	for i := n; i > 0 && c > 0; i-- {
		if table[i][c] != table[i-1][c] {
			chosenPresents = append(chosenPresents, presents[i-1])
			c -= presents[i-1].Size
		}
	}

	return chosenPresents
}
