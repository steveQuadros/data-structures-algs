package amazonoa

/**
find all unique combinations that sum up to target
ex: n = 5
[][]int{
	{1, 1, 1, 1, 1},
	{1, 1, 1, 2},
	{1, 1, 3},
	{1, 4},
	{2, 2, 1},
	{2, 3},
}
**/
func SumToTarget(n int) [][]int {
	var res [][]int
	// sumToTargetRecursive(1, n, 0, []int{}, &res)
	return res
}

func sumToTargetRecursive(start, end, sum int, path []int, res *[][]int) {
	if sum >= end {
		if sum == end {
			cp := make([]int, len(path))
			copy(cp, path)
			*res = append(*res, cp)
		}
		return
	}

	for i := start; i < end; i++ {
		sum += i
		path = append(path, i)
		sumToTargetRecursive(i, end, sum, path, res)
		sum -= i
		path = path[:len(path)-1]
	}
}
