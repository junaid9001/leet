func scoreOfString(s string) int {
    sum:=0
    for i:=0;i<len(s)-1;i++{
        diff:=int(s[i])-int(s[i+1])
        if diff<0{
            sum-=diff
        }else{
            sum+=diff
        }
    }
    return sum
}