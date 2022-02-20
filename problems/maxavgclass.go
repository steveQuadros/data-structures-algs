package problems

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	var avgRatio float64
	calcMAR(0, classes, extraStudents, 0.0, &avgRatio)
	return avgRatio
}

func calcMAR(classIdx int, classes [][]int, extraStudents int, curRatio float64, avgRatio *float64) {
	if extraStudents == 0 || classIdx == len(classes) {
		if curRatio/float64(len(classes)) > *avgRatio {
			*avgRatio = curRatio
		}
		return
	}

	for extraStudents > 0 {
		extraStudents--
		classes[classIdx][0]++
		classes[classIdx][1]++
		curRatio += float64(classes[classIdx][0] / classes[classIdx][1])

		calcMAR(classIdx+1, classes, extraStudents, curRatio, avgRatio)

		extraStudents++
		classes[classIdx][0]--
		classes[classIdx][1]--
		curRatio -= float64(classes[classIdx][0] / classes[classIdx][1])
	}
}
