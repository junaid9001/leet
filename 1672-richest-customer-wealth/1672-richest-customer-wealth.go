func maximumWealth(accounts [][]int) int {
    n:=len(accounts)
    max:=0
    for i:=0;i<n;i++{
        current:=accounts[i]
        total:=0
        for _,val:=range current{
            total+=val
        }
        if total>max{
            max=total
        }
    }
    return max
}