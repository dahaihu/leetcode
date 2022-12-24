package wiggle_max_length

type status int8

const (
	statusEqual status = 0
	statusBig   status = 1
	statusSmall status = 2
	statusInit  status = 3
)

func getStatus(cur int) status {
	if cur > 0 {
		return statusBig
	} else if cur < 0 {
		return statusSmall
	} else {
		return statusEqual
	}
}

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	preStatus := statusInit
	count := 1
	for i := 1; i < len(nums); i++ {
		switch curStatus := getStatus(nums[i] - nums[i-1]); curStatus {
		case statusEqual, preStatus:
			continue
		default:
			preStatus = curStatus
			count++
		}
	}
	return count
}
