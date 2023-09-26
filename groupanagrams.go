func groupAnagrams(strs []string) [][]string {
	group := make(map[[26]int][]string)
	result := [][]string{}

	if len(strs) == 0 {
		return [][]string{{}}
	}
	for _, v := range strs {
		key := [26]int{}
		for i := 0; i < len(v); i++ {
			key[v[i]-'a']++
		}
		group[key] = append(group[key], v)
	}
	for _, v := range group {
		result = append(result, v)
	}
	return result
}
