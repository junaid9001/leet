func findTheDifference(s string, t string) byte {
    m:=make(map[byte]int)
    b:=[]byte(s)

    for _,v:=range b {
        m[v]++
    }

    for _,v:=range t{
        if m[byte(v)]==0{
            return byte(v)
        }
        m[byte(v)]--
    }
    return byte(0)
}