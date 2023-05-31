package practics

type NumArray struct {
	nums []int
}

func NewNumArray(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

func (na *NumArray) SumRange(left int, right int) int {
	if right > len(na.nums)-1 {
		return 0
	}

	sum := 0
	for i := left; i <= right; i++ {
		sum += na.nums[i]
	}

	return sum
}
