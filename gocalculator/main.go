package main

import (
	"fmt" // For formatting and error messages
	"log" // For logging errors and program status
	"os" // For accessing command-line arguments
	"strconv" // For converting string arguments to numeric types
)

// Calculator provides basic arithmetic operations
type Calculator struct{}

// Add performs addition of two numbers
func (c *Calculator) Add(a, b float64) float64 {
	return a + b
}

// Subtract performs subtraction of two numbers
func (c *Calculator) Subtract(a, b float64) float64 {
	return a - b
}

// Multiply performs multiplication of two numbers
func (c *Calculator) Multiply(a, b float64) float64 {
	return a * b
}

// Divide performs division of two numbers with error handling
// Returns an error if attempting to divide by zero to prevent runtime panic
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")// Check for division by zero - a error condition
	}
	return a / b, nil
}

// PerformCalculation executes the specified arithmetic operation
func PerformCalculation(operation string, num1, num2 float64) (float64, error) {
	calc := &Calculator{}
	
	switch operation {
	case "add":
		return calc.Add(num1, num2), nil
	case "subtract":
		return calc.Subtract(num1, num2), nil
	case "multiply":
		return calc.Multiply(num1, num2), nil
	case "divide":
		return calc.Divide(num1, num2)
	default:
		return 0, fmt.Errorf("Unsupported operation: %s", operation)
	}
}

func main() {
	// Check if sufficient command-line arguments are provided
	// The program expects exactly 4 arguments: Program name, Operation (add/subtract/multiply/divide), First number and Second number
	if len(os.Args) != 4 {
		log.Fatalf("Usage: %s <operation> <num1> <num2>\nSupported operations: add, subtract, multiply, divide", os.Args[0])
	}

	// Parse operation and numbers
	operation := os.Args[1]
	num1, err1 := strconv.ParseFloat(os.Args[2], 64)
	num2, err2 := strconv.ParseFloat(os.Args[3], 64)

	// Error handling for parsing
	if err1 != nil || err2 != nil {
		log.Fatal("Invalid number arguments")
	}

	// Perform calculation
	result, err := PerformCalculation(operation, num1, num2)
	if err != nil {
		log.Fatal(err)
	}

	// Print result
	fmt.Printf("Result of %s: %.2f\n", operation, result)
}