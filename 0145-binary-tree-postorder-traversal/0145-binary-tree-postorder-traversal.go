/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
  var res []int
  PostOrder(root,&res)
  return res
}

func PostOrder(root *TreeNode,slice *[]int){
    if root==nil{
        return 
    }
    PostOrder(root.Left,slice)
    PostOrder(root.Right,slice)
    *slice=append(*slice,root.Val)
}