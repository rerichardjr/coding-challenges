class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        results = set()
        for k in nums:
            if k in results:
                return True
            results.add(k)
        return False
