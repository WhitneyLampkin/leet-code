/*
You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.


Example 1:

Input: flowerbed = [1,0,0,0,1], n = 1
Output: true
Example 2:

Input: flowerbed = [1,0,0,0,1], n = 2
Output: false
 

Constraints:

1 <= flowerbed.length <= 2 * 104
flowerbed[i] is 0 or 1.
There are no two adjacent flowers in flowerbed.
0 <= n <= flowerbed.length
*/

// Better Performance
func canPlaceFlowersOptimized(flowerbed []int, n int) bool {
    planted := 0

    for i, _ := range flowerbed {
        if flowerbed[i] == 0 {
            emptyLeftPlot := (i == 0 ) || (flowerbed[i - 1] == 0)
            emptyRightPlot := (i == len(flowerbed) - 1) || (flowerbed[i + 1] == 0)

            if emptyLeftPlot && emptyRightPlot {
                flowerbed[i] = 1
                planted += 1
            }
        }
    }

    return int(planted) >= n
}

// My original solution
func canPlaceFlowers(flowerbed []int, n int) bool {
    planted := 0

    for i := 0; i < len(flowerbed); i++ {
        if len(flowerbed) == 1 {
            if flowerbed[i] == 0 {
                flowerbed[i] = 1
                planted += 1
            }
        } else {
            if i == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
                flowerbed[i] = 1
                planted += 1
            } else if i == len(flowerbed) - 1 && flowerbed[i] == 0 && flowerbed[i-1] == 0 {
                flowerbed[i] = 1
                planted += 1
            } else if i > 0 && flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
                flowerbed[i] = 1
                planted += 1
            } 
        }
    }

    // Optimizes but negatively affects the overall runtime
    if planted >= n {
        return true
    }

    return int(planted) >= n
}
