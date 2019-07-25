package main

func main() {
	a := []int{1, 1, 1, 0}
	nextPermutation(a)
	//c := maximalRectangle(a)
	//fmt.Println(c)
}
func nextPermutation(nums []int) {
	i := len(nums) - 2
	j := i + 1
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1)
}
func swap(nums []int, i int, j int) {
	var temp int
	temp = nums[j]
	nums[j] = nums[i]
	nums[i] = temp
}
func reverse(nums []int, begin int) {
	var j int
	i := begin
	j = len(nums) - 1
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}
