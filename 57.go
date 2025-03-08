package main

func main() {
	//{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}
	//{1, 3}, {6, 9}
	intervals := [][]int{{0, 5}, {8, 9}}
	newIntervals := []int{3, 4}
	minMaxInterval := []int{-1, -1}
	var outInterval [][]int
	i := 0
	for newIntervals[0] > intervals[i][0] && i < len(intervals) {
		if newIntervals[0] > intervals[i][1] {
			outInterval = append(outInterval, intervals[i])
			i++
		} else {
			minMaxInterval[0] = intervals[i][0]
			break
		}
	}
	if minMaxInterval[0] == -1 {
		minMaxInterval[0] = newIntervals[0]
	}
	for newIntervals[1] >= intervals[i][0] && i < len(intervals) {
		if newIntervals[1] > intervals[i][1] {
			i++
		} else if newIntervals[1] <= intervals[i][1] {
			minMaxInterval[1] = intervals[i][1]
			i++
			break
		}
	}
	if minMaxInterval[1] == -1 {
		minMaxInterval[1] = newIntervals[1]
	}
	outInterval = append(outInterval, minMaxInterval[:])
	outInterval = append(outInterval, intervals[i:]...)
	println("1:", minMaxInterval[0], "2:", minMaxInterval[1])
	/*	var point [2]int
		if len(intervals) == 0 {
			intervals = append(intervals, newIntervals)
			return
		}
		for i, j := 0, 0; i < len(intervals); i++ {
			if j == 0 {
				if newIntervals[j] >= intervals[i][0] {
					if newIntervals[j] < intervals[i][1] {
						point[j] = i
						j++
						i--
					} else if newIntervals[j] >= intervals[i][1] {
						if newIntervals[j] == intervals[i][1] {
							point[j] = i
							j++
							i--
						}
						continue
					}
				} else {
					intervals[i][0] = newIntervals[j]
					point[j] = i
					j++
					i--
				}
			} else {
				if newIntervals[j] > intervals[i][1] && point[0] != i {
					intervals = append(intervals[:i], intervals[i+1:]...)
					i--
				} else if newIntervals[j] >= intervals[i][0] && point[0] != i {
					intervals[i-1][1] = intervals[i][1]
					if i != len(intervals)-1 {
						intervals = append(intervals[:i], intervals[i+1:]...)
						return
					} else {
						intervals = intervals[:i]
						return
					}
					i--
				} else if newIntervals[j] < intervals[i][0] {
					intervals[i-1][1] = newIntervals[j]
				} else if len(intervals) == 1 {
					if newIntervals[j] > intervals[i][1] {
						intervals[i][1] = newIntervals[j]
						return
					}
				}
			}*/
	return
}
