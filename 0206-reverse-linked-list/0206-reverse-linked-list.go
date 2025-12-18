/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var vals []int
    current:=head
    for current!=nil{
        vals=append(vals,current.Val)
        current=current.Next
    }

    for i:=0;i<len(vals)/2;i++{
        vals[i],vals[len(vals)-1-i]=vals[len(vals)-1-i],vals[i]
    }
    if len(vals)<1{
        return nil
    }
    newList:=&ListNode{Val:vals[0]}
    cr2:=newList
    for i:=1;i<len(vals);i++{
        cr2.Next=&ListNode{Val:vals[i]}
        cr2=cr2.Next
    }
    return newList
}