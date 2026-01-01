func plusOne(digits []int) []int {
    n:=len(digits)

    for i:=n-1;i>=0;i--{
        if digits[i]<9{
            digits[i]+=1
            return digits
        }else{
            digits[i]=0
        }
    }
    ans:=make([]int,n+1)
    ans[0]=1
    return ans
}