func isAnagram(s string, t string) bool {
	character := make([]int, 26)

	for _, val := range s {
		i := int(val - 'a')
		character[i]++
	}

	for _, val := range t {
		i := int(val - 'a')
		character[i]--
	}

	for _, val := range character {
		if val != 0 {
			return false
		}
	}
	return true
}