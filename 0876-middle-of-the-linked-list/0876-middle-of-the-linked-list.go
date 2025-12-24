/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    var count int

    current:=head
    c2:=head

    for current!=nil{
        count++
        current=current.Next
    }
    tg:=count/2
    for i:=1;i<=tg;i++{
        c2=c2.Next
    }
    return c2
}