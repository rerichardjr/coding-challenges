func isAnagram(s string, t string) bool {
  result := make(map[rune]int)
  // if lengths are different then not an anagram
	if len(s) != len(t) {
		return false
	}
  //count letter frequency in s and save to result
	for _, v := range s {
		result[v]++
	}
  //subtract from letter frequency in result
	for _, v := range t {
    if _, ok := result[v]; ok {
			result[v]--
		}
	}
  //if a non 0 value exists in result, then s and t are not anagrams
	for _, v := range s {
		if result[v] != 0 {
			return false
		}
	}
	return true
}
