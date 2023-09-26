func containsDuplicate(nums []int) bool {
    //results is hashmap to hold nums value as key and their index as value
    results := make(map[int]int)
    for k, v := range nums {
        //have we seen this number in nums before?
        if _, ok := results[v]; ok {
            //found a duplicate
            return true
        }
        //update results to show that we've seen this number
        results[v] = k
    }
    //no duplicates found
    return false
}
