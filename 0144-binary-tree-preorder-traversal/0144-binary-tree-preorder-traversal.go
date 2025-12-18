/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    var res []int
    PreOrder(root,&res)
    return res
}

func PreOrder(root *TreeNode,slice *[]int){
    if root==nil{
        return
    }
    *slice=append(*slice,root.Val)
    PreOrder(root.Left,slice)
    PreOrder(root.Right,slice)

}