func topKFrequent(nums []int, k int) []int {
  //need to write something better, seems like this could be improved
	result := make(map[int]int)
	count := make(map[int][]int)
	topK := []int{}
	for j := range nums {
		result[nums[j]]++
	}
	for i, j := range result {
		count[j] = append(count[result[i]], i)
	}
 	for i := len(nums); i >= 0; i-- {
		if _, ok := count[i]; ok {
			if len(topK) != k {
					topK = append(topK, count[i]...)
			}
		}
	}
	return topK
}
