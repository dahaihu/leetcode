package wiggle_max_length

type status int8

const (
	statusEqual status = 0
	statusBig   status = 1
	statusSmall status = 2
	statusInit  status = 3
)

func compareStatus(val int) status {
	switch {
	case val == 0:
		return statusEqual
	case val > 0:
		return statusBig
	case val < 0:
		return statusSmall
	default:
		panic("invalid switch branch")
	}
}

func wiggleMaxLength(nums []int) int {
	if length := len(nums); length <= 1 {
		return length
	}
	preStatus := statusInit
	count := 1
	for i := 1; i < len(nums); i++ {
		curStatus := compareStatus(nums[i] - nums[i-1])
		if curStatus == statusEqual || curStatus == preStatus {
			continue
		}
		count++
		preStatus = curStatus
	}
	return count
}
