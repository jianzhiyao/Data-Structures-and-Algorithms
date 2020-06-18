package SortedBinaryTree

type Node struct {
	data  int
	left  *Node
	right *Node
}

type SortedBinaryTree struct {
	root *Node
}

func (tree *SortedBinaryTree) InsertItem(i int) {
	if tree.root == nil {
		tree.root = &Node{data: i}
		return
	}
	currentNode := tree.root
	for {
		if i > currentNode.data {
			if currentNode.right == nil {
				currentNode.right = &Node{data: i}
				return
			}
			currentNode = currentNode.right
		} else {
			if currentNode.left == nil {
				currentNode.left = &Node{data: i}
				return
			}
			currentNode = currentNode.left
		}
	}
}

func (tree *SortedBinaryTree) SearchItem(i int) (*Node, bool) {
	if tree.root == nil {
		return nil, false
	}
	currentNode := tree.root
	for currentNode != nil {
		if i == currentNode.data {
			return currentNode, true
		} else if i > currentNode.data {
			currentNode = currentNode.right
		} else if i < currentNode.data {
			currentNode = currentNode.left
		}
	}
	return nil, false
}

func (tree *SortedBinaryTree) InorderTraversal(subtree *Node, callback func(int)) {
	if subtree == nil {
		return
	}

	tree.InorderTraversal(subtree.left, callback)
	callback(subtree.data)
	tree.InorderTraversal(subtree.right, callback)
}

func (tree *SortedBinaryTree) PreorderTraversal(subtree *Node, callback func(int)) {
	if subtree == nil {
		return
	}

	callback(subtree.data)
	tree.PreorderTraversal(subtree.left, callback)
	tree.PreorderTraversal(subtree.right, callback)
}

func (tree *SortedBinaryTree) PostorderTraversal(subtree *Node, callback func(int)) {
	if subtree == nil {
		return
	}

	tree.PostorderTraversal(subtree.left, callback)
	tree.PostorderTraversal(subtree.right, callback)
	callback(subtree.data)
}

func (tree *SortedBinaryTree) GetRoot() *Node {
	return tree.root
}
