package heapsort

func down(nums []int, idx int, margin int) {
	for idx < margin {
		left := 2*idx + 1
		if left >= margin {
			return
		}
		minChild := left
		right := left + 1
		if right < margin && nums[right] > nums[minChild] {
			minChild = right
		}
		if nums[idx] >= nums[minChild] {
			return
		}
		nums[idx], nums[minChild] = nums[minChild], nums[idx]
		idx = minChild
	}
}

func heapSort(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		down(nums, i, len(nums))
	}
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		down(nums, 0, i)
	}
}
