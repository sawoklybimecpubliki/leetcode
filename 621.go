package main

func leastInterval(tasks []byte, n int) int {
	var count int
	//var key byte
	task := make(map[byte]int)
	for _, i := range tasks {
		task[i]++
		/*if task[i] > count {
			count = task[i]
			key = i
		}*/
	}
	count = 0
	lengT := len(task)
	for lengT > 0 {
		c := 0
		for k := range task {
			if task[k]-1 > 0 {
				task[k]--
				count++
				c++
			} else if task[k]-1 == 0 {
				delete(task, k)
				count++
				c++
			}
			if c == n {
				c = 0
				break
			}
		}
		if lengT <= n {
			if len(task) == 0 {
				return count
			}
			count = count + (n + 1 - lengT)
			lengT = len(task)
		} else {
			lengT = len(task)
		}
	}
	return count
}

func main() {
	println(leastInterval([]byte{'A', 'B', 'C', 'A', 'B', 'C'}, 2))
}
