func isAnagram(s string, t string) bool {
    if len(s) != len(t) { return false }
    sr:=[]rune(s)
    br:=[]rune(t)
    for i:=0;i<len(sr);i++{
        if !slices.Contains(br,sr[i]){
            return false
        }
        indx:=slices.Index(br,sr[i])
        
        br=slices.Delete(br,indx,indx+1)

    }

    return true
}