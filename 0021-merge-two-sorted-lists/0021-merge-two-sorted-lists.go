/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var arr []int
	for list1 != nil {
		arr = append(arr, list1.Val)
		list1 = list1.Next
	}
	for list2 != nil {
		arr = append(arr, list2.Val)
		list2 = list2.Next
	}
    if len(arr)<1{
        return nil
    }
	slices.Sort(arr)
	newList := &ListNode{Val:arr[0]}
    dummy:=newList
    for i:=1;i<len(arr);i++{
        dummy.Next=&ListNode{Val:arr[i]}
        dummy=dummy.Next
    }
    return newList

}