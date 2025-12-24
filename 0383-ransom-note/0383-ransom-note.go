func canConstruct(ransomNote string, magazine string) bool {
    dupe:=[]byte(magazine)
    b:=[]byte(ransomNote)
    for i:=0;i<len(b);i++{
        if slices.Contains(dupe,b[i]){
            z:=slices.Index(dupe,b[i])
            dupe=slices.Delete(dupe,z,z+1)
        }else{
            return false
        }
    }
    return true
}