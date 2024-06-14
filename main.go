package main

import (
    "github.com/01-edu/z01"
)

func ItoaBase(value, base int) string {
    if base < 2 || base > 16 {
        return "" // or handle error
    }

    var result []int
    isNegative := value < 0
    if isNegative {
        value = -value
    }

    // Handle the case where value is 0
    if value == 0 {
        result = append(result, 48) // '0' in ASCII
    }

    // Convert value to the specified base
    for value > 0 {
        remainder := value % base
        result = append([]int{remainder}, result...)
        value /= base
    }

    // Handle negative numbers
    if isNegative {
        result = append([]int{'-'}, result...)
    }

    // Convert each int to a rune and print it
    for _, r := range result {
        if r >= 0 && r <= 9 {
            z01.PrintRune(rune('0' + r))
        } else {
            z01.PrintRune(rune('A' + r - 10))
        }
    }
    z01.PrintRune('\n')

    return ""
}

func main() {
    ItoaBase(-40, 8) // Test case for base 16 with negative value
}
