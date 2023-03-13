func sortString(str string) string {
	strArr := strings.Split(str, "")
	sort.Strings(strArr)
	return strings.Join(strArr, "")
}

func groupAnagrams(strs []string) [][]string {
	maps := make(map[string][]string)

	for _, str := range strs {
		key := sortString(str)
		if _, ok := maps[key]; !ok {
			maps[key] = []string{}
		}
		maps[key] = append(maps[key], str)
	}

	res := make([][]string, len(maps))
	i := 0
	for _, anagrams := range maps {
		res[i] = anagrams
		i++
	}

	return res
}