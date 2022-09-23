package gsa

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
		//add v.Left to stack, use seen-map to avoid inf loop
		if v.Left != nil && !seen[v.Left] {
			stack = append(stack, v.Left)
		} else {
			//pop v from stack
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
	queue := make([]*T, 0)
	result := []int{}

	if v == nil {
		return result
	}

	queue = append(queue, v)
	for len(queue) > 0 {
		//dequeue
		v = queue[0]
		queue = queue[1:]
		result = append(result, v.Val)

		//add left to queue
		if v.Left != nil {
			queue = append(queue, v.Left)
		}
		//add right to queue
		if v.Right != nil {
			queue = append(queue, v.Right)
		}
	}
	return result
}
