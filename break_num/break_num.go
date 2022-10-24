package break_num

/*
在你的面前从左到右摆放着n根长短不一的木棍，你每次可以折断一根木棍，
并将折断后得到的两根木棍一左一右放在原来的位置（即若原木棍有左邻居，则两根新木棍必须放在左邻居的右边，若原木棍有右邻居，新木棍必须放在右邻居的左边，所有木棍保持左右排列）。
折断后的两根木棍的长度必须为整数，且它们之和等于折断前的木棍长度。你希望最终从左到右的木棍长度单调不减，那么你需要折断多少次呢？

输入：[3,5,13,9,12]
输出：1

输入：[3,12,13,9,12]
输出：2

输入：[3,13,12,9,12]
输出：3

输入：[3,13,60,7]
输出：10

输入：[3,63,7]
输出：8

输入：[9,1]
输出：8
*/

func breakNum(ticks []int) int {
	preMargin := ticks[len(ticks)-1]
	times := 0
	for i := len(ticks) - 2; i >= 0; i-- {
		tick := ticks[i]
		if tick <= preMargin {
			preMargin = tick
			continue
		}
		curTimes, remainder := tick/preMargin, tick%preMargin
		if remainder == 0 {
			times += curTimes - 1
			continue
		}
		times += curTimes
		preMargin = tick / (curTimes + 1)
	}
	return times
}
