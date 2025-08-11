package main

func main() {
	nums := [3]int{1, 2, 3}
	nums2 := append(nums[:], nums[:]...)
	println(nums2)
}
