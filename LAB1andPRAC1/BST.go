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

    if key <= root.Key {
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

