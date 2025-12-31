func mostWordsFound(sentences []string) int {
    var a int
    for _,val:=range sentences{
        c:=strings.Fields(val)
        if len(c)>a{
            a=len(c)
        }
    }
    return a
}