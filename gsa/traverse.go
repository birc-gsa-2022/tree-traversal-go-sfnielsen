package gsa

import "fmt"

type T struct {
	Val         int
	Left, Right *T
}

// Do an in-order traversal of v and output the
// values in the tree.
func InOrder(v *T) []int {
	stack := make([]*T, 0)
	seen := make(map[*T]bool)
	result := []int{}

	if v == nil {
		return result
	}

	stack = append(stack, v)
	for len(stack) > 0 {
		v = stack[len(stack)-1]
		seen[v] = true
		fmt.Println(v, v.Left, v.Right, v.Val)
		//add v.Left to stack, use seen-map to avoid inf loop
		if v.Left != nil && !seen[v.Left] {
			stack = append(stack, v.Left)
		} else {
			//pop v from stack
			fmt.Println(v)
			result = append(result, v.Val)
			stack = stack[:len(stack)-1]
			//add v.Right to stack
			if v.Right != nil {
				stack = append(stack, v.Right)
			}
		}

	}

	return result
}

// Do a breadth-order traversal of v and output the
// values in the tree.
func BfOrder(v *T) []int {
	return []int{}
}
