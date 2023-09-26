class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        results = {}
        for k, v in enumerate(nums):
            if target-v in results:
                return [results[target-v], k]
            results[v] = k
        return []
