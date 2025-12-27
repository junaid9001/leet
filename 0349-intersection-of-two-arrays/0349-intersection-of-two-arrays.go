func intersection(nums1 []int, nums2 []int) []int {
    var ans []int
    set:=make(map[int]bool)
    for _,n:=range nums1{
        set[n]=true
    }

    for _,val:=range nums2{
        if set[val]{
            ans=append(ans,val)
            delete(set,val)
        }
    }
    return ans
}