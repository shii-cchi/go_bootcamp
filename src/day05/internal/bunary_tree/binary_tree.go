package bunary_tree

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func areToysBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftToys := countToys(root.Left)
	rightToys := countToys(root.Right)

	return leftToys == rightToys
}

func countToys(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftToys := countToys(node.Left)
	rightToys := countToys(node.Right)

	if node.HasToy {
		return leftToys + rightToys + 1
	}

	return leftToys + rightToys
}

func unrollGarland(root *TreeNode) []bool {
	if root == nil {
		return nil
	}

	var result []bool
	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		size := len(queue)
		currentLevel := make([]bool, size)

		for i := 0; i < size; i++ {
			node := queue[i]

			if level%2 != 0 {
				currentLevel[i] = node.HasToy
			} else {
				currentLevel[size-i-1] = node.HasToy
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLevel...)

		queue = queue[size:]
		level++
	}

	return result
}
