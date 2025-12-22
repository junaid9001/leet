func findCenter(edges [][]int) int {
    var ans int
    set:=make(map[int][]int)
    for _,vals:=range edges{
        u:=vals[0]
        v:=vals[1]
        set[u]=append(set[u],v)
        set[v]=append(set[v],u)
    }

    for key,vals:=range set{
        if len(vals)==len(edges){
            ans=key
            break
        }
    }
    return ans
}