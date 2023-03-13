func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	i := 0
	j := 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		if len(s) == i {
			return true
		}
		j++
	}
	return false
}