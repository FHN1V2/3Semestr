package main

import (
    "fmt"
)

type BSTNode struct {
    Key   int
    Left  *BSTNode
    Right *BSTNode
}

func NewNode(key int) *BSTNode {
    return &BSTNode{Key: key, Left: nil, Right: nil}
}

func BSTadd(root *BSTNode, key int) *BSTNode {
    if root == nil {
        return NewNode(key)
    }

    if key < root.Key {
        root.Left = BSTadd(root.Left, key)
    } else if key > root.Key {
        root.Right = BSTadd(root.Right, key)
    }

    return root
}

/*func PrintTree(root *BSTNode, prefix string, isRoot bool) {
    if root != nil {
        fmt.Println(prefix + func(isRoot bool) string {
            if isRoot {
                return " "
            }
            return "├── "
        }(isRoot) + fmt.Sprintf("%d", root.Key))

        if root.Left != nil || root.Right != nil {
            newPrefix := prefix + func(isRoot bool) string {
                if isRoot {
                    return " "
                }
                return "│   "
            }(isRoot)

            if root.Left != nil {
                PrintTree(root.Left, newPrefix, false)
            }
            if root.Right != nil {
                PrintTree(root.Right, newPrefix, false)
            }
        }
    }
}*/
func PrintTree(BSTNode *BSTNode, prefix string, isLeft bool) {
    if BSTNode == nil {
        fmt.Println("Empty tree")
        return
    }

    if BSTNode.Right != nil {
        PrintTree(BSTNode.Right, prefix + "   ", false)
    }

    fmt.Print(prefix)
    if isLeft {
        fmt.Print("↳⟶ ")
    } else {
        fmt.Print("↱⟶ ")
    }
    fmt.Println(BSTNode.Key)

    if BSTNode.Left != nil {
        PrintTree(BSTNode.Left, prefix + "   ", true)
    }
}
func BStdel(root *BSTNode, value int) *BSTNode {
	if root == nil {
		return root
	}

	if value < root.Key {
		root.Left = BStdel(root.Left, value)
	} else if value > root.Key {
		root.Right = BStdel(root.Right, value)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		root.Key = minValue(root.Right)
		root.Right = BStdel(root.Right, root.Key)
	}

	return root
}
func minValue(root *BSTNode) int {
	minValue := root.Key
	for root.Left != nil {
		minValue = root.Left.Key
		root = root.Left
	}
	return minValue
}

func BstFind(root *BSTNode, value int) string {
	if root == nil {
		return "false"
	}

	if value == root.Key {
		return "true"
	}

	if value < root.Key {
		return BstFind(root.Left, value)
	} else {
		return BstFind(root.Right, value)
	}
}