package find_median_sorted_arrays

import "math"

func splitArrays(a, b []int) (aIdx, bIdx int) {
	m, n := len(a), len(b)
	left, right := 0, m
	for left < right {
		aMid := (right-left)/2 + left
		bMid := (m+n)/2 - aMid - 1
		if b[bMid] <= a[aMid] {
			right = aMid
		} else {
			left = aMid + 1
		}
	}
	return right, (m+n)/2 - right
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func arrayValue(nums []int, idx int, defaultValue int) int {
	if idx >= 0 && idx < len(nums) {
		return nums[idx]
	}
	return defaultValue
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	nums1Mid, nums2Mid := splitArrays(nums1, nums2)
	rightMargin := min(
		arrayValue(nums1, nums1Mid, math.MaxInt64),
		arrayValue(nums2, nums2Mid, math.MaxInt64),
	)
	if (m+n)%2 == 1 {
		return float64(rightMargin)
	}
	leftMargin := max(
		arrayValue(nums1, nums1Mid-1, math.MinInt64),
		arrayValue(nums2, nums2Mid-1, math.MinInt64),
	)
	return float64(leftMargin+rightMargin) / 2
}
