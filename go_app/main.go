// Created by: Jaden Plugowsky
// Created on: April 2023
//
// This program calculates book sale discounts based on the amount of books purchased

package main

import "fmt"

var bookCost = 5.00 // In dollars
var bookDiscount = 0.00
var books float64
var preTotal float64
var discountedTotal float64

// Used for adding a dollar sign to the number
var preTotalString string
var bookDiscountString string
var discountedTotalString string

func main() {
	// This function calculates book sale discounts based on the amount of books purchased
	// Input
	fmt.Println("\nThis function calculates book sale discounts based on the amount of books purchased.")
	fmt.Println("\nEnter the amount of books you wish to purchase: ")
	fmt.Scanln(&books)

	// Process
	preTotal = books * bookCost // Simply the amount of books selected multiplied by the cost of each book
	if books >= 3 {
		bookDiscount += bookCost // This adds the bookCost of 1 book to the discount, making 1 book free (The 3rd book)
	}
	if books > 3 {
		bookDiscount += (bookCost / 2) * (books - 3.00)
	}
	discountedTotal = preTotal - bookDiscount

	// Output
	// Formats the numbers into strings and then adds a "$" to them
	preTotalString = "$" + (fmt.Sprintf("%v", preTotal))
	bookDiscountString = "$" + (fmt.Sprintf("%v", bookDiscount))
	discountedTotalString = "$" + (fmt.Sprintf("%v", discountedTotal))

	// Prints the formatted string versions of the numbers
	fmt.Println("\nBefore discount cost of the books is: " + preTotalString)
	fmt.Println("Money discounted is: " + bookDiscountString)
	fmt.Println("The final total cost of the books is: " + discountedTotalString)

	fmt.Println("\nDone.")
}
