package algorithm

//26. 删除排序数组中的重复项
func removeDuplicates(nums []int) int {
	left := 0
	n := len(nums)
	for right := 0; right < n; right++ {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}

//27. 移除元素
func removeElement(nums []int, val int) int {
	left := 0
	n := len(nums)
	for right := 0; right < n; right++ {
		if nums[right] == val {
			continue
		} else {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}
