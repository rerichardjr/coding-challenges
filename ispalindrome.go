func isPalindrome(x int) bool {
  rev := 0
	tmp := x

	//loop until tmp is 0
	for tmp > 0 {
		//use mod 10 to get last digit
		ld := tmp % 10
		//reverse the number
		rev = rev*10 + ld
		//divide tmp by 10
		tmp = tmp / 10
	}

	//is x a palindrome number?
	if x == rev {
		return true
	}
	return false
}
