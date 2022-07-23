package running_sum_1d_array

/*
Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

Return the running sum of nums.



Example 1:

Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].

*/
func runningSum(nums []int) []int {
	var res []int
	if len(nums) == 1 {
		return nums
	}

	res = append(res, nums[0])

	for i := 1; i < len(nums); i++ {
		res = append(res, res[i-1]+nums[i])
	}

	return res
}
