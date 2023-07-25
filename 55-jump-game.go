/*
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.


Example 1:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
 

Constraints:

1 <= nums.length <= 104
0 <= nums[i] <= 105
*/

/*
func canJump(nums []int) bool {
    isPossible := true

    if len(nums) > 1 {
        for i := 0; i <= len(nums) - 2; i++ {
            for j := nums[i]; j >= 0; j-- {
                if i != 0 && nums[i] == 0 && nums[i-1] < 2 {
                    jumpsToPass := 0
                    for k := 0; k < i; k++ {
                        jumpsToPass += nums[k]
                    }
                    if jumpsToPass <= i {
                        isPossible = false
                        break
                    }
                }
            }
        }
    }

    if nums[0] == 0 && len(nums) > 1 {
        isPossible = false
    }

    return isPossible
}
*/
