package main

import "math/rand"

func QuickSelectKthSmallest(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	// translate k into a proper index (2nd smallest becomes index 1 in sorted array, 1st -> 0)
	return quickSelect(nums, k-1, 0, len(nums)-1)
}

func QuickSelectKthLargest(nums []int, k int) int {
	if k >= len(nums) {
		return -1
	}
	// translate k into a proper index (2nd largest in len 5 arr becomes idx 3
	return quickSelect(nums, len(nums)-k, 0, len(nums)-1)
}

func quickSelect(nums []int, k, left, right int) int {
	if left == right {
		return nums[left]
	}

	pivotIdx := left + rand.Intn(right-left)
	pivotIdx = lumotoPartition(nums, left, right, pivotIdx)

	if pivotIdx < k {
		return quickSelect(nums, k, pivotIdx+1, right)
	} else if pivotIdx > k {
		return quickSelect(nums, k, left, pivotIdx-1)
	} else {
		return nums[k]
	}
}

func lumotoPartition(nums []int, left, right, pivotIdx int) int {
	pivot := nums[pivotIdx]
	swap(nums, right, pivotIdx)
	storeIdx := left
	for left <= right {
		if nums[left] < pivot {
			swap(nums, left, storeIdx)
			storeIdx++
		}
		left++
	}
	swap(nums, storeIdx, right)
	return storeIdx
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
