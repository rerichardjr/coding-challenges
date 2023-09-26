func twoSum(nums []int, target int) []int {
	//results is hashmap to hold nums value as key and their index as value
  results := make(map[int]int)

	for k, v := range nums {
		//do we have a key in results that equals target-v? if yes return the value.
		if result, ok := results[target-v]; ok {
			//yes, return the indices
			return []int{result, k}
		}
		//update results with v as key and k as value for tracking 
		results[v] = k
	}
	//return empty int slice if we didn't find two numbers that add up to target
	//shouldn't hit this since we can assume exactly one solution for each case
	return []int{}
}
