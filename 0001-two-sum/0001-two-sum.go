func twoSum(nums []int, target int) []int {
    m:=make(map[int]int)

    for i,num:=range nums{
        
        val,ok:=m[num]
        if !ok{ 
            m[target-num]=i
            continue
        }
        return []int{val,i}
        
    }
    return nil
}
//loop over nums
//check if current val exist as key in map
//adding- key==next num to find to match and val==keys index 
//