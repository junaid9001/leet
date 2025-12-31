func canAliceWin(nums []int) bool {
    var sd int
    var dd int 
   

    for _,val:=range nums{
        if val<10{
            sd+=val
        }else if val>9&&val<100{
            dd+=val
        }
    }

    if dd!=sd{
        return true
    }
    return  false
}