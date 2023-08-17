/*
50. Pow(x, n)
Medium

Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

Example 1:

Input: x = 2.00000, n = 10
Output: 1024.00000
Example 2:

Input: x = 2.10000, n = 3
Output: 9.26100
Example 3:

Input: x = 2.00000, n = -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
 
Constraints:

-100.0 < x < 100.0
-231 <= n <= 231-1
n is an integer.
Either x is not zero or n > 0.
-104 <= xn <= 104
*/

// Recursive solution
func myPow(x float64, n int) float64 {
    powerResult := 1.0

    if n == 0 {
        return 1
    } else if n < 0 {
        return 1/myPow(x, -n)
    }
    
    for n != 0 {
        if n&1 == 1 {
            powerResult *= x
        }

        x *= x
        n >>= 1
    }

    return powerResult
}

/*

First Attempt:

func myPow(x float64, n int) float64 {
    powerResult := x

    if n == 0 {
        return float64(1)
    } 
    
    switch x {
        case 1:
            return x
        case -1:
            if int(math.Abs(float64(n)))%2 == 0 {
                return float64(1)
            } else {
                return x
            }
        default: 
            for i := 1; i < int(math.Abs(float64(n))); i++ {
                powerResult *= x
            }
    }

    if n < 0 {
        powerResult = 1/(powerResult)
    }

    return powerResult
}
*/
