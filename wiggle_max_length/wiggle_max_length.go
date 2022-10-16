package wiggle_max_length

type gapStatus int

const (
	gapStatusBig   gapStatus = 1
	gapStatusSmall gapStatus = -1
	gapStatusEqual gapStatus = 0
	gapStatusInit  gapStatus = 100
)

func getGapStatus(gap int) gapStatus {
	switch {
	case gap > 0:
		return gapStatusBig
	case gap == 0:
		return gapStatusEqual
	case gap < 0:
		return gapStatusSmall
	default:
		panic("invalid gapStatus")
	}
}

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 1
	preStatus := gapStatusInit
	for i := 1; i < len(nums); i++ {
		curStatus := getGapStatus(nums[i] - nums[i-1])
		switch curStatus {
		case preStatus, gapStatusEqual:
			continue
		default:
			preStatus = curStatus
			count += 1
		}
	}
	return count
}
