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

func isValid(s string) bool {
    var openBrackets []rune
    hasOpenParenthesis := 0
    hasOpenBracket := 0
    hasOpenCurlyBrace := 0
    result := true

    if len(s) % 2 != 0 {
        return false
    } else if len(s) == 0 {
        return true
    }

    for _, value := range s {
        if value == '(' || value == '[' || value == '{' {
            openBrackets = append(openBrackets, value)

            switch value {
                case '(':
                    hasOpenParenthesis++
                case '[':
                    hasOpenBracket++
                case '{': 
                    hasOpenCurlyBrace++
                default:
                    break
            }
        } else {
            if len(openBrackets) == 0 {
                return false
            }

            lastValue := openBrackets[0]
            lastIdx := len(openBrackets) - 1

            if lastIdx >= 0 {
                lastValue = openBrackets[lastIdx]
            }
            

            switch value {
                case ')':
                    if lastValue != '(' {
                        return false
                    } else {
                        openBrackets = openBrackets[:len(openBrackets)-1]
                        hasOpenParenthesis--
                    }
                case ']':
                    if lastValue != '[' {
                        return false
                    } else {
                        openBrackets = openBrackets[:len(openBrackets)-1]
                        hasOpenBracket--
                    }
                case '}':
                    if lastValue != '{' {
                        return false
                    } else {
                        openBrackets = openBrackets[:len(openBrackets)-1]
                        hasOpenCurlyBrace--
                    }
                default:
                    break
            }
        }
    }

    if hasOpenParenthesis > 0 || hasOpenBracket > 0 || hasOpenCurlyBrace > 0 {
        result = false
    }

    return result
}

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
