import "sort"
func merge(nums1 []int, m int, nums2 []int, n int)  {
   ans := append([]int{}, nums1[:m]...)
    ans = append(ans, nums2...)

    sort.Ints(ans)

    for i := 0; i < len(ans); i++ {
        nums1[i] = ans[i]
    }
    
}