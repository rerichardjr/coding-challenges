func productExceptSelf(nums []int) []int {
  //1rst attempt failed with oom on case 19 of 22
  //2nd attempt uses prefix product from left to right and right to left, then multiplies each element from both lists

	numLen := len(nums)
	productL := make([]int, numLen)
	productR := make([]int, numLen)
	answer := make([]int, numLen)

	productL[0] = 1
	productR[numLen-1] = 1

	for i := 1; i < numLen; i++ {
		j := numLen - i - 1
		productL[i] = nums[i-1] * productL[i-1]
		productR[j] = nums[j+1] * productR[j+1]
	}

	for i := 0; i < numLen; i++ {
		answer[i] = productL[i] * productR[i]
	}

	return answer

}
