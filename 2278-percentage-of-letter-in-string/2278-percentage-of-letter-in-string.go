func percentageLetter(s string, letter byte) int {
    var c int
    b:=[]byte(s)
    n:=len(s)
    for _,v:= range b{
        if v==letter{
            c++
        }
    }
    ans:= (c*100)/n
    return ans
}