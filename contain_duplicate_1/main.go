package contain_duplicate_1

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int][]int)
	for idx, val := range nums {
		numsMap[val] = append(numsMap[val], idx)
	}

	for _, val := range numsMap {
		if len(val) > 1 {
			return true
		}
	}

	return false
}
