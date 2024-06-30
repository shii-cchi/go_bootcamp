package binary_tree

import (
	"reflect"
	"testing"
)

func TestAreToysBalanced_SingleNode(t *testing.T) {
	root := &TreeNode{HasToy: true}

	expected := true

	result := areToysBalanced(root)

	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestAreToysBalanced_NilNode(t *testing.T) {
	root := &TreeNode{}

	expected := true

	result := areToysBalanced(root)

	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestAreToysBalanced_UnbalancedTree1(t *testing.T) {
	root := &TreeNode{HasToy: true}
	root.Left = &TreeNode{HasToy: true}
	root.Right = &TreeNode{HasToy: false}
	root.Left.Left = &TreeNode{HasToy: true}
	root.Left.Right = &TreeNode{HasToy: false}
	root.Right.Left = &TreeNode{HasToy: true}

	expected := false

	result := areToysBalanced(root)

	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestAreToysBalanced_UnbalancedTree2(t *testing.T) {
	root := &TreeNode{HasToy: false}
	root.Left = &TreeNode{HasToy: true}
	root.Right = &TreeNode{HasToy: false}
	root.Left.Left = &TreeNode{HasToy: true}
	root.Left.Right = &TreeNode{HasToy: true}
	root.Right.Left = &TreeNode{HasToy: true}
	root.Right.Right = &TreeNode{HasToy: false}

	expected := false

	result := areToysBalanced(root)

	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestAreToysBalanced_BalancedTree1(t *testing.T) {
	root := &TreeNode{HasToy: false}
	root.Left = &TreeNode{HasToy: false}
	root.Right = &TreeNode{HasToy: true}
	root.Left.Left = &TreeNode{HasToy: false}
	root.Left.Right = &TreeNode{HasToy: true}

	expected := true

	result := areToysBalanced(root)

	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestAreToysBalanced_BalancedTree2(t *testing.T) {
	root := &TreeNode{HasToy: true}
	root.Left = &TreeNode{HasToy: true}
	root.Right = &TreeNode{HasToy: false}
	root.Left.Left = &TreeNode{HasToy: true}
	root.Left.Right = &TreeNode{HasToy: false}
	root.Right.Left = &TreeNode{HasToy: true}
	root.Right.Right = &TreeNode{HasToy: true}

	expected := true

	result := areToysBalanced(root)

	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestUnrollGarland_SingleNode(t *testing.T) {
	root := &TreeNode{HasToy: true}

	expected := []bool{true}

	result := unrollGarland(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestUnrollGarland_NilNode(t *testing.T) {
	var root *TreeNode

	expected := []bool(nil)
	result := unrollGarland(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestUnrollGarland_Var1(t *testing.T) {
	root := &TreeNode{HasToy: true}
	root.Left = &TreeNode{HasToy: true}
	root.Right = &TreeNode{HasToy: true}
	root.Left.Left = &TreeNode{HasToy: true}
	root.Left.Right = &TreeNode{HasToy: false}
	root.Right.Left = &TreeNode{HasToy: true}
	root.Right.Right = &TreeNode{HasToy: true}

	expected := []bool{true, true, true, true, true, false, true}
	result := unrollGarland(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestUnrollGarland_Var2(t *testing.T) {
	root := &TreeNode{HasToy: true}
	root.Left = &TreeNode{HasToy: true}
	root.Right = &TreeNode{HasToy: false}
	root.Left.Left = &TreeNode{HasToy: true}
	root.Left.Right = &TreeNode{HasToy: false}
	root.Right.Left = &TreeNode{HasToy: true}
	root.Right.Right = &TreeNode{HasToy: true}

	expected := []bool{true, true, false, true, true, false, true}
	result := unrollGarland(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
