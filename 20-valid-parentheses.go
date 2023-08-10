/*
20. Valid Parentheses
Easy
21.3K
1.4K
Companies
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
 

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
 

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

// Original solution
/*
hasOpenParenthesis := false
    hasOpenBracket := false
    hasOpenCurlyBrace := false

    // Iterate through s string
    for _, value := range s {
        // Open bracket?
            // Check if appropriate flag is true
                // Yes - return false
                // No - Set appropriate flag to true
        // Closing bracket?
            // Check if appropriate flag is true
                // Yes - Set flag to false and continue
                // No - return false
        switch value {
            case '(':
                if hasOpenParenthesis || hasOpenBracket || hasOpenCurlyBrace {
                    return false
                } else {
                    hasOpenParenthesis = true
                }
            case ')':
                if hasOpenParenthesis && !hasOpenBracket && !hasOpenCurlyBrace {
                    hasOpenParenthesis = false
                } else {
                    return false
                }
            case '[':
                if hasOpenParenthesis || hasOpenBracket || hasOpenCurlyBrace {
                    return false
                } else {
                    hasOpenBracket = true
                }
            case ']':
                if !hasOpenParenthesis && hasOpenBracket && !hasOpenCurlyBrace {
                    hasOpenBracket = false
                } else {
                    return false
                }
            case '{':
                if hasOpenParenthesis || hasOpenBracket || hasOpenCurlyBrace {
                    return false
                } else {
                    hasOpenCurlyBrace = true
                }
            case '}':
                if !hasOpenParenthesis && !hasOpenBracket && hasOpenCurlyBrace {
                    hasOpenCurlyBrace = false
                } else {
                    return false
                }
            default:
                return false
        }
        */
