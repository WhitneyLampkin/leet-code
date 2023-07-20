func lengthOfLongestSubstring(s string) int {
    longestLength := 0
    
    if len(s) <= 1 {
            longestLength = len(s)
    } else {
        for i, _ := range s {
            runeValues := make(map[byte]bool)
            currentLength := 0
            for j := i; j < len(s); j++ {
                if runeValues[s[j]] {
                    break
                }
                
                runeValues[s[j]] = true
                currentLength++
            }

            if currentLength > longestLength {
            longestLength = currentLength
            }
        }
    }

    return longestLength
}
