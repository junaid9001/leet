func sortArray(nums []int) []int {
    if len(nums)<=1{
        return nums
    }
    mid:=len(nums)/2
    left:=sortArray(nums[:mid])
    right:=sortArray(nums[mid:])
    return merge(left,right)
}

func merge(left,right []int)[]int{
    r:=make([]int,0,len(left)+len(right))
    i,j:=0,0
    
    for i<len(left)&&j<len(right){
        if left[i]<=right[j]{
            r=append(r,left[i])
            i++
        }else{
            r=append(r,right[j])
            j++
        }
    }
    r=append(r,left[i:]...)
    r=append(r,right[j:]...)
    return r
}