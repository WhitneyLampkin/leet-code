func sumOfMultiples(n int) int {
    // Declare variables
    total := 0

    // Iterate through each number in the range
        // Check if divisible by 3, 5, or 7
            // Yes? - Add to the overall total
            // No? - Go to next value
    for i := 1; i <= n; i++ {
        if (i % 3) == 0 || (i % 5) == 0 || (i % 7) == 0 {
            total += i
        }
    }

    // return sum of valid numbers
    return total
}
