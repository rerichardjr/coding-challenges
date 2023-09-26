class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        result = {}
        if len(s) != len(t):
            return False
        for _, v in enumerate(s):
            result[v] = 1 + result.get(v, 0)
        for _, v in enumerate(t):
            if v in result:
                result[v] = result.get(v, 0) - 1
        for _ , v in enumerate(s):
            if result[v] != 0:
                return False
        return True
