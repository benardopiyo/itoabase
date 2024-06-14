# ItoaBase

* The ItoaBase converts an integer value into its string representation in a specified base. This conversion supports bases ranging from 2 to 16, utilizing the characters '0' through '9' for digits and 'A' through 'F' for values 10 to 15.


## Function Signature

```bash
func ItoaBase(value, base int) string {
}
```

## Parameters

* **value:** The integer value to convert. This can be any integer within the range representable in Go.
* **base:** The base to convert the integer into, ranging from 2 to 16.

## Return Value

* The function returns a string representing the integer value in the specified base. If the value is negative, the resulting string will include a leading '-' sign.
Usage

* The function handles the conversion of integers to their respective representations in the specified base. For bases beyond decimal (base 10), it employs characters A to F for values 10 through 15.
Examples

* Here are some examples demonstrating the usage of the ItoaBase function:

1. Converting 255 to Base 16 (Hexadecimal)

```bash
result := ItoaBase(255, 16)
fmt.Println(result) // Outputs: "FF"
```

2. Converting -31 to Base 2 (Binary)
```bash
result := ItoaBase(-31, 2)
fmt.Println(result) // Outputs: "-11111"
```

3. Converting 42 to Base 8 (Octal)
```bash
result := ItoaBase(42, 8)
fmt.Println(result) // Outputs: "52"
```

4. Converting 10 to Base 4

```bash
result := ItoaBase(10, 4)
fmt.Println(result) // Outputs: "22"
```

## Validity Constraints

* The ItoaBase function assumes valid inputs within the specified range of bases (2 to 16).
* It is designed to handle negative integers by appropriately formatting the result with a minus sign.