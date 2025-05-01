package main

import (
	"fmt"
	"sort"
)

/*
func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	var result int
	slices.Sort(tasks)
	slices.Sort(workers)
	i, j := len(workers)-1, len(tasks)-1
	for j >= 0 {
		if tasks[j] <= workers[i] {
			result++
			j--
			i--
		} else {
			for k := 0; k <= i; k++ {
				if pills > 0 && tasks[j] <= workers[k]+strength {
					pills--
					result++
					i--
					workers = append(workers[:k], workers[k+1:]...)
					break
				}
			}
			j--
		}
		if i < 0 {
			break
		}
	}
	return result
}*/

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	task_requirements := append([]int(nil), tasks...)
	worker_strengths := append([]int(nil), workers...)
	total_pills := pills
	pill_boost := strength
	sort.Ints(task_requirements)
	sort.Ints(worker_strengths)

	can_assign := func(k int) bool {
		avail := append([]int(nil), worker_strengths[len(worker_strengths)-k:]...)
		pills_remain := total_pills
		for i := k - 1; i >= 0; i-- {
			req := task_requirements[i]
			n := len(avail)
			if n > 0 && avail[n-1] >= req {
				avail = avail[:n-1]
			} else {
				if pills_remain <= 0 {
					return false
				}
				threshold := req - pill_boost
				idx := sort.Search(len(avail), func(j int) bool { return avail[j] >= threshold })
				if idx == len(avail) {
					return false
				}
				avail = append(avail[:idx], avail[idx+1:]...)
				pills_remain--
			}
		}
		return true
	}

	low, high, completed := 0, min(len(tasks), len(workers)), 0
	for low <= high {
		mid := (low + high) / 2
		if can_assign(mid) {
			completed = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return completed
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxTaskAssign([]int{7908, 9988, 9571, 5279, 4047, 9760, 3274, 7098, 6367, 4774, 9975, 5544, 8811, 2564, 3835, 6634,
		5648, 9052, 8143, 5502, 1285, 7300, 5630, 7578, 6522, 2243, 7284, 6534, 3766, 1835, 4153, 4333, 6338, 6373, 7517, 6930, 8190, 2834,
		2218, 2945, 6929, 4170, 9254, 9312, 9789, 1324, 3851, 4038, 9497, 6486, 7949, 5781, 7787, 5185, 2726, 8538, 3698, 6929, 2613, 6860,
		6981, 7506, 5294, 6213, 9848, 4539, 1234, 8108, 3832, 5068, 2712, 6301, 8340, 7950, 1320, 9111, 5411, 8075, 9752, 6882, 9679, 9463,
		6580, 9986, 5114, 4483, 7816, 1878, 2204, 4364, 1603, 3131, 3492, 8579, 3026, 2521, 5679, 8630, 8667, 3827, 1555, 3296, 1118, 4644,
		4866, 1312, 7632, 9550, 6914, 9195, 1009, 2145, 5184, 4996, 2913, 6914, 3584, 4866, 8505, 7708, 1309, 2780, 1794, 6103, 6161, 9576,
		9885, 1843, 2180, 9261, 6516, 7411, 8606, 4633, 5653, 8562, 9533, 8516, 2503, 3270, 5264, 1737, 1603, 1376, 1834, 6990, 8234, 8975,
		4194, 2276, 7633, 3401, 3521, 6003, 7685, 2139, 9061, 2786, 8637, 7517, 1411, 3732, 9200, 9772, 8017, 4008, 9189, 4728, 5141, 8600,
		5512, 4239, 5236, 6909, 9942, 4712, 7017, 9755, 2365, 3130, 4668, 9101, 1414, 7787, 4533, 7845, 3755, 2534, 7896, 4710, 7581, 2301,
		5079, 4236, 7276, 8656, 4777, 1857},
		[]int{4612, 4669, 4273, 1929, 1925, 1342, 4565, 2152, 1909, 4897, 2697, 1419, 1998, 4789, 2940, 3388, 4619, 1427, 1798, 917, 2270,
			1639, 1633, 4669, 2674, 2348, 3714, 3909, 4359, 3151, 1639, 4533, 3444, 3531, 887, 528, 2346, 1481, 831, 2345, 3610, 4507, 2890,
			4740, 3959, 648, 1468, 3063, 2416, 4091, 1151, 2196, 999, 2708, 578, 978, 4113, 4229, 2052, 2154, 2149, 583, 1029, 2020, 4498,
			4638, 4932, 2986, 2749, 1047, 2434, 971, 1645, 570, 1395, 3603, 4581, 3607, 3158, 2758, 1046, 1171, 2204, 1421, 3708, 4769, 3399,
			981, 3396, 4837, 3633, 4067, 4141, 2037, 1622, 2257, 1529, 1371, 4090, 3273, 4767, 3543, 1098, 503, 3905, 4427, 3875, 2594, 4878,
			2252, 912, 4838, 509, 2989, 3304, 4268, 1520, 4027, 3691, 993, 600, 3529, 3762, 1854, 592, 1579, 2191, 3051, 1897, 4050, 1345,
			3416, 1108, 4862, 4587, 3065, 3486, 2138, 1272, 4237, 4235, 4107, 1951, 3747, 2774, 4944, 3815, 2560, 3688, 1838, 2821, 3098,
			1361, 1439, 2486, 4599, 2215, 1908, 2432, 1862, 1780, 1143, 4117, 3175, 3839, 3019, 4576, 4343, 4186, 1643, 3472, 1637, 693,
			4128, 4709, 4500, 2639, 4022, 2886, 1708, 2721, 4686, 4870, 772, 557, 1845, 2788, 821, 3752},
		35, 2208))
	/*[]int{10, 15, 30}, []int{0, 10, 10, 10, 10}, 3, 10*/

}
