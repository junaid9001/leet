func maximumHappinessSum(happiness []int, k int) int64 {
    var ans int64
    n:=len(happiness)
	slices.Sort(happiness)

    for i:=0;i<k;i++{
        current:=int64(happiness[n-1-i]-i)

        if current>0{
            ans+=current
        }else{
            break
        }
    }

    return ans
}