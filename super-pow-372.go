/*
372. Super Pow
Medium

Your task is to calculate ab mod 1337 where a is a positive integer and b is an extremely large positive integer given in the form of an array.


Example 1:

Input: a = 2, b = [3]
Output: 8
Example 2:

Input: a = 2, b = [1,0]
Output: 1024
Example 3:

Input: a = 1, b = [4,3,3,8,5,2]
Output: 1
 

Constraints:

1 <= a <= 231 - 1
1 <= b.length <= 2000
0 <= b[i] <= 9
b does not contain leading zeros.
*/

/*
<TODO: To be continued...totally misread the problem description and don't have time to sort it out today...>
func superPow(a int, b []int) int {
    switch a {
        case 0:
            return 0
        case 1:
            return 1
    }

    result := 1
    var bString string
    var bConverted int

    for i := 0; i < len(b); i++ {
        bString += strconv.FormatInt(int64(b[i]), 10)
    }

    bConverted, _ = strconv.Atoi(bString)

    if bConverted == 0 { 
        return 1 
    } else if bConverted == 1 {
        return a
    }

    for i := 1; i <= bConverted; i++ {
        result *= a
    }

    return result
}
*/
