package amazonoa

/**
Given 2 lists a and b. Each element is a pair of integers where the
first integer represents the unique id and the second integer represents
a value. Your task is to find an element from a and an element form b
such that the sum of their values is less or equal to target and as close
to target as possible. Return a list of ids of selected elements. If no
pair is possible, return an empty list.

Input:
a = [[1, 2], [2, 4], [3, 6]]
b = [[1, 2]]
target = 7

Output: [[2, 1]]

Explanation:
There are only three combinations [1, 1], [2, 1], and [3, 1], which have a total sum of 4, 6 and 8, respectively.
Since 6 is the largest sum that does not exceed 7, [2, 1] is the optimal pair.
Example 2:

Input:
a = [[1, 3], [2, 5], [3, 7], [4, 10]]
b = [[1, 2], [2, 3], [3, 4], [4, 5]]
target = 10

Output: [[2, 4], [3, 2]]

Explanation:
There are two pairs possible. Element with id = 2 from the list `a` has a value 5, and element with id = 4 from the list `b` also has a value 5.
Combined, they add up to 10. Similarily, element with id = 3 from `a` has a value 7, and element with id = 2 from `b` has a value 3.
These also add up to 10. Therefore, the optimal pairs are [2, 4] and [3, 2].
Example 3:

Input:
a = [[1, 8], [2, 7], [3, 14]]
b = [[1, 5], [2, 10], [3, 14]]
target = 20

Output: [[3, 1]]
Example 4:

Input:
a = [[1, 8], [2, 15], [3, 9]]
b = [[1, 8], [2, 11], [3, 12]]
target = 20

Output: [[1, 3], [3, 2]]
**/

/**
Input:
a = [[1, 3], [2, 5], [3, 7], [4, 10]]
b = [[1, 2], [2, 3], [3, 4], [4, 5]]
target = 10

Output: [[2, 4], [3, 2]]
if <= target
- if > curMax replace
- if == curMax append
**/

import (
	"sort"
)

func OptimalUtilization(a, b [][]int, target int) [][]int {
	var curMax int
	res := [][]int{}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			sum := a[i][1] + b[j][1]
			if i == j || sum > target || a[i][0] == b[j][0] {
				continue
			}

			curPair := []int{a[i][0], b[j][0]}
			if sum > curMax {
				res = [][]int{curPair}
				curMax = sum
			} else if sum == curMax {
				res = append(res, curPair)
			}
		}
	}

	return res
}

func OptimalUtilization2P(a, b [][]int, target int) [][]int {
	var curMax int
	var res [][]int
	sortPairs(a)
	sortPairs(b)
	left, right := 0, len(b)-1

	for left < len(a) && right >= 0 {
		sum := a[left][1] + b[right][1]
		if sum <= target {
			curPair := []int{a[left][0], b[right][0]}
			if sum > curMax {
				res = [][]int{curPair}
				curMax = sum
			} else if sum == curMax {
				res = append(res, curPair)
			}
		}
		if sum >= target {
			if right >= 0 {
				right--
			}
		} else {
			if left < len(a) {
				left++
			}
		}
	}

	return res
}

func sortPairs(a [][]int) {
	sort.Slice(a, func(i, j int) bool {
		return a[i][1] < a[j][1]
	})
}
