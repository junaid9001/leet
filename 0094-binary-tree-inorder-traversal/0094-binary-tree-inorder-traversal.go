/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
   var newSlice []int
    Inorder(root,&newSlice)
    return newSlice
}

func Inorder(root *TreeNode,slice *[]int){
    if root==nil{
        return
    }
    Inorder(root.Left,slice)
    *slice=append(*slice,root.Val)
    Inorder(root.Right,slice)
}