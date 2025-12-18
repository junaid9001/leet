/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    var res []int

    for i:=0;i<len(lists);i++{
        current:=lists[i]
        for current!=nil{
            res=append(res,current.Val)
            current=current.Next
        }
    }
    slices.Sort(res)
    if len(res)<1{
        return nil
    }
    head:=&ListNode{Val:res[0]}
    current:=head
    for i:=1;i<len(res);i++{
        current.Next=&ListNode{Val:res[i]}
        current=current.Next
    }
    return head
}