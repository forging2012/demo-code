package other_code

// 2, 7, 11, 15
// 给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。--逆向

// 64ms
func twoSum(nums []int, target int) []int {
	t := len(nums) - 1
	result := make([]int, 0)

	for i := 0; i <= t-1; i++ {
		outNum := nums[i]
		n := t
		tmp := i
		for n > tmp {
			if nums[n]+outNum == target {
				result = append(result, i)
				result = append(result, n)
				return result
			}
			n--
		}
	}
	return result

}

// 60ms
func toSum2(nums []int, target int) []int {

	result := make([]int, 0)
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j]+nums[i] == target {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}

	return result
}

// 8ms--空间换时间
func toSum3(num []int, target int) []int {
	result := []int{}
	mp := make(map[int]int, 0)
	for index, value := range num {
		mp[value] = index
	}

	for index, value := range num {
		complement := target - value
		value, ok := mp[complement]
		ok2 := value != index
		if ok && ok2 {
			result = append(result, index)
			result = append(result, mp[complement])
			return result
		}
	}
	return result
}

// 时间和之前的都差不多--空间换时间
func toSum4(nums []int, target int) []int {
	result := []int{}
	mp := make(map[int]int, 0)

	for index, value := range nums {
		complement := target - value
		if value, ok := mp[complement]; ok {
			result = append(result, value)
			result = append(result, index)
			return result
		}
		mp[value] = index
	}
	return result
}
