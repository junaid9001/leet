/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root==nil{
        return nil
    }
    Invert(root)
    return root
    
}

func Invert(node *TreeNode){
    if node==nil{
        return
    }
    node.Left,node.Right=node.Right,node.Left
    Invert(node.Left)
    Invert(node.Right)
    
}