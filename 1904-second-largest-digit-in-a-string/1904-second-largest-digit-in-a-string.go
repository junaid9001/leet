func secondHighest(s string) int {
    var arr []int
    max:=-1
    min:=-1

    for _,val:=range s{
        if val>=48&&val<=57{
            arr=append(arr,int(val-'0'))
        }
    }
    if len(arr)>0{
        for _,val:=range arr{
            if val>max{
                min=max
                max=val
            }else if val<max&&val>min{
              min=val
            }
        }
    }
    return min
}