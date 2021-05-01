func replaceSpace(s string) string {
    ans := ""
    for _,v := range s{
        if v == ' '{
            ans = ans + "%20"
        } else {
            ans = ans + string(v)
        }
    }
    return ans
}