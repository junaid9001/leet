/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 import "sort"
func sortList(head *ListNode) *ListNode {
    current:=head
    if current==nil{
        return nil
    }
    var vals []int
    for current!=nil{
        vals=append(vals,current.Val)
        current=current.Next
    }
    sort.Ints(vals)
    newHead:=&ListNode{Val:vals[0]}
    current=newHead

    for i:=1;i<len(vals);i++{
        current.Next=&ListNode{Val:vals[i]}
        current=current.Next
    }
    return newHead

}