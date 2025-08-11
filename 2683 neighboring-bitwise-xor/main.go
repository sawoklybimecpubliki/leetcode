package main

func main() {
	derived := []int{1, 0}
	original_last := 0
	for i := 0; i < len(derived)-1; i++ {
		original_last ^= derived[i]
	}
	if 0^original_last != derived[len(derived)-1] {
		original_last ^= 1
		/*for i := 0; i < len(derived)-1; i++ {
			original_last ^= derived[i]
		}*/
		if 1^original_last != derived[len(derived)-1] {
			println("false")
		}
	} else {
		println("true")
	}
}
