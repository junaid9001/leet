func firstPalindrome(words []string) string {
   
    for _,w:=range words{
        i,j:=0,len(w)-1
        ispal:=true
        for i<j{
            if w[i]!=w[j]{
                ispal=false
            }
            i++
            j--
        }
        if ispal{
            return w
        }
    }
    return ""
}