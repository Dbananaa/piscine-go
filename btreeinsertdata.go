package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// func BTreeInsertData(root *TreeNode, data string) *TreeNode {
// 	return bTreeInsertData(root, root, data)
// }

// func bTreeInsertData(root *TreeNode, current *TreeNode, data string) *TreeNode {
// 	if current == nil {
// 		return &TreeNode{
// 			Data: data,
// 		}
// 	}

// 	if current.Data > data && current.Left == nil {
// 		current.Left = &TreeNode{
// 			Data: data,
// 		}

// 		return root
// 	}

// 	if current.Data < data && current.Right == nil {
// 		current.Right = &TreeNode{
// 			Data: data,
// 		}

// 		return root
// 	}

// 	if current.Data < data {
// 		return bTreeInsertData(root, current.Right, data)
// 	}

// 	return bTreeInsertData(root, current.Left, data)
// }

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	return bTreeInsertData(root, nil, root, data)
}

func bTreeInsertData(root, parent, current *TreeNode, data string) *TreeNode {
	if current == nil {
		return &TreeNode{
			Data:   data,
			Parent: parent,
		}
	}

	if current.Data > data && current.Left == nil {
		current.Left = &TreeNode{
			Data:   data,
			Parent: current,
		}

		return root
	}

	if current.Data < data && current.Right == nil {
		current.Right = &TreeNode{
			Data:   data,
			Parent: current,
		}

		return root
	}

	if current.Data < data {
		return bTreeInsertData(root, current, current.Right, data)
	}

	return bTreeInsertData(root, current, current.Left, data)
}
