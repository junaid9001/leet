func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
    var ans []int

	for k, val := range nums2 {
		for i := k + 1; i < len(nums2); i++ {
			if nums2[i] > val {
				m[val] = nums2[i]
				break
			}
		}
	}

    for _,v:=range nums1{
        if m[v]>0{
            ans=append(ans,m[v])
        }else{
            ans=append(ans,-1)
        }
    }

    return ans
}