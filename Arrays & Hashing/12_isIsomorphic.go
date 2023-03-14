func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var charMapS, charMapT [128]byte
	for i := 0; i < len(s); i++ {
		charS, charT := s[i], t[i]
		if (charMapS[charS] != 0 && charMapS[charS] != charT) || (charMapT[charT] != 0 && charS != charMapT[charT]) {
			return false
		}
		charMapS[charS], charMapT[charT] = charT, charS
	}
	return true
}