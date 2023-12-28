func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, val := range nums {
		diff := target - val
		if mapIndex, ok := m[diff]; ok {
			return []int{mapIndex, i}
		}
		m[val] = i
	}
	return []int{}
}
