func fizzBuzz(n int) []string {
	answer := make([]string, n)
	for i := 1; i <= n; i++ {
    	mod3 := i % 3
			mod5 := i % 5
			if mod3 == 0 && mod5 == 0 {
					answer[i-1] = "FizzBuzz"
					continue
			}
			if mod3 == 0 {
					answer[i-1] = "Fizz"
					continue
			}
			if mod5 == 0 {
					answer[i-1] = "Buzz"
					continue
			}
			answer[i-1] = strconv.Itoa(i)
	}
	return answer
}
