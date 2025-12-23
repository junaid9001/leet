func differenceOfSums(n int, m int) int {
    var div []int
    var nondiv []int
    x,z:=0,0
    for i:=1;i<=n;i++{
        if i%m==0{
            div=append(div,i)
        }else{
            nondiv=append(nondiv,i)
        }
    }
    for _,val:=range div{
        x+=val
    }
    for _,val:=range nondiv{
        z+=val
    }
    return z-x
}