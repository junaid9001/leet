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
    
    for head!=nil{
        if count[head.Next]{
            return true
        }else{
            count[head]=true
        }
        head=head.Next

    }
    return false
}