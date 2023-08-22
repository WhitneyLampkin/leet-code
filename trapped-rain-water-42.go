/*
42. Trapping Rain Water
Hard
27.9K
390
Companies
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.


Example 1:

Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
Example 2:

Input: height = [4,2,0,3,2,5]
Output: 9
 

Constraints:

n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105
*/

func trap(height []int) int {
    // Declare trappedWater placehold var
    trappedWaterCounter := 0
    // Declare var to hold count of trapped water
    trappedWaterTotal := 0
    // Declare highest possible values for next check
    highestCurrentElevation := height[0]
    highestPossibleElevation := 0

    // Iterate, i, from 0 to vals less than length of the slice - 2
        // Check if value is equal to highest possible elevation
            // Lower
                // Inc trapped water counter till equal (add highestElevationPossible - i)
            // Higher or equal
                // Reset highest possible elevation 
                // Add trapped water counter value to total water trapped
                // Reset total water trapped value
                // Reset trapped water counter value

    return trappedWaterTotal
}
