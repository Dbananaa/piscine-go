package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Data == elem {
		return root
	}

	if r := BTreeSearchItem(root.Left, elem); r != nil {
		return r
	}

	if r := BTreeSearchItem(root.Right, elem); r != nil {
		return r
	}

	return nil
}
