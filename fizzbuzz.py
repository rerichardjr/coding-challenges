class Solution:
    def fizzBuzz(self, n: int) -> List[str]:
        answer = []
        for i in range(1,n+1) :
            mod3 = i % 3
            mod5 = i % 5
            if (mod3 == 0) and (mod5 == 0) :
                answer.append("FizzBuzz")
                continue
            if mod3 == 0 :
                answer.append("Fizz")
                continue
            if mod5 == 0 :
                answer.append("Buzz")
                continue
            answer.append(str(i))
        return answer
