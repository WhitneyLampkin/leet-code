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

/*

///////////////////// Runner's Up a.k.a. Failed Attempts (LOL) /////////////////////

func lengthOfLongestSubstring(s string) int {
    currentLength := 0
    longestLength := 0
    runeValues := make(map[rune]bool)
    
    for i, item := range s {
        if len(s) == 0 {
            break
        } else if _, value := runeValues[item]; len(s) > 1 && value && s[i] != s[i - 1] {
            currentLength = 2
        } else if _, value := runeValues[item]; len(s) == 1 || i == 0 || s[i] == s[i - 1] || value {
            currentLength = 1
        } else {
            currentLength += 1
        }

        if currentLength > longestLength {
            longestLength = currentLength
        }

        if len(s) > 0 {
            runeValues[item] = true
        } 
    }

    return longestLength
}

func ugh(s string) int {
	currentLength := 0
	longestLength := 0
	runeValues := make(map[rune]bool)

	for i, value := range s {
		if len(s) == 0 {
			break
		} else if len(s) == 1 {
			currentLength = 1
		}

		if _, value := runeValues[value]; i == 0 || s[i] == s[i-1] || value {
			fmt.Println(value)
			fmt.Println(s[i], s[i-1])
			currentLength = 1
		} else {
			fmt.Println("here")
			currentLength += 1
		}

		if currentLength > longestLength {
			longestLength = currentLength
		}

		if len(s) > 0 {
			runeValues[value] = true
		}
	}

	fmt.Println(longestLength)

	return longestLength
}

func test(s string) int {
	//curentSubstring := len(s)
	//	longestSubstring := len(s)
	//	index = 0
	count := 0

	for i, current := range s {
		count++
		for j, next := range s {
			fmt.Println(count)
			fmt.Println(i, current)
			fmt.Println(j, next)
		}
	}

	return 0
}

func lengthOfLongestSubstring(s string) int {
	substringMap := make(map[rune]bool)
	currentSubstring := []rune{}
	longestSubstring := []rune{}
	var lastValue rune

	for _, item := range s {
		// range provides both the index and value
		// _ is being used to ignore the unused index value
		// For the if statement below, statements can precede conditionals.
		// The condition is !value.
		if _, value := substringMap[item]; !value && item != lastValue {
			substringMap[item] = true
			currentSubstring = append(currentSubstring, item)
			if len(currentSubstring) >= len(longestSubstring) {
				longestSubstring = currentSubstring
			}
		} else {
			if len(currentSubstring) >= len(longestSubstring) {
				longestSubstring = currentSubstring
				currentSubstring = []rune{}
			}
			substringMap[item] = true
			currentSubstring = append(currentSubstring, item)
		}

		lastValue = item
	}

	fmt.Println(len(longestSubstring))

	return len(longestSubstring)
}
*/
