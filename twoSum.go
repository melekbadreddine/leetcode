func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)

	for currentIndex, val := range nums {
		complement := target - val

		if complementIndex, found := indexMap[complement]; found {
			return []int{complementIndex, currentIndex}
		}
		indexMap[val] = currentIndex
	}
	return nil
}
