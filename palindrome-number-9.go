/*
9. Palindrome Number
Easy
10.4K
2.6K
Companies
Given an integer x, return true if x is a 
palindrome
, and false otherwise.

 

Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
 

Constraints:

-231 <= x <= 231 - 1
 

Follow up: Could you solve it without converting the integer to a string?
*/

func isPalindrome(x int) bool {
    convertedInt := strconv.FormatInt(int64(x), 10)

    // Check if palindromic
        // Yes - continue
        // No - return false
    for i := 0; i < int(math.Floor(float64(len(convertedInt)/2))); i++ {
        j := len(convertedInt) - (i + 1)
        
        if convertedInt[i] != convertedInt[j] {
            return false
        }
    }

    return true
}
