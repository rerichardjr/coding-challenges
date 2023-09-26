class Solution:
    def isPalindrome(self, x: int) -> bool:
        rev = 0
        tmp = x

        while tmp > 0 :
            ld = tmp % 10
            rev = rev*10 + ld
            tmp = tmp // 10

        if x == rev:
            return True
        return False
