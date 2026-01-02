func reverseWords(s string) string {
    words:=strings.Fields(s)

    for i,v:=range words{
        b:=[]byte(v)
        j,k:=0,len(b)-1
        for j<k{
            b[j],b[k]=b[k],b[j]
            j++
            k--
        }
        words[i]=string(b)
    }
    return strings.Join(words," ")
}