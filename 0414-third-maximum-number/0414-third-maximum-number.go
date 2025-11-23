func thirdMax(nums []int) int {
    f,s,t:=math.MinInt64,math.MinInt64,math.MinInt64

    for _,v:=range nums{
        if v==f||v==s||v==t{
            continue
        }
        if v>f{
            t=s
            s=f
            f=v
        }else if v>s{
            t=s
            s=v
        }else if v>t{
            t=v
        }
    }

    if t== math.MinInt64{
        return f
    }
    return t
}