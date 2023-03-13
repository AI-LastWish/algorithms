func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false   
    }

    var maps [26]int

    for i := 0; i < len(s); i++ {
        maps[s[i] - 'a']++
        maps[t[i] - 'a']--
    }

    for i := 0; i < len(maps); i++ {
        if maps[i] != 0 {
        	return false   
        }
    }

    return true
}