/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head==nil||head.Next==nil{
        return false
    }
    
    count:=make(map[*ListNode]bool)
    current:=head
    for current!=nil{
        if count[current.Next]{
            return true
        }else{
            count[current]=true
        }
        current=current.Next

    }
    return false
}