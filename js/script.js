// Copyright (c) 2023 Jaden Plugowsky All rights reserved
//
// Created by: Jaden Plugowsky
// Created on: April 2023
// This file contains the JS functions for index.html

"use strict"

const bookCost = 5.0 //In dollars

function calculatePressed() {
  let bookDiscount = 0.0
  //This function calculates book sale discounts based on the amount of books purchased
  //Input through Textfields
  const books = parseFloat(document.getElementById("books").value)

  //Process
  const preTotal = books * bookCost //Simply the amount of books selected multiplied by the cost of each book
  if (books >= 3) {
    bookDiscount = bookDiscount + bookCost //This adds the bookCost of 1 book to the discount, making 1 book free (The 3rd book)
  }
  if (books > 3) {
    bookDiscount = bookDiscount + (bookCost / 2) * (books - 3)
    /**
     *This takes the bookCost and divides it by 2 to get the 50% off cost (bookCost / 2)
     *Then, it multiplies the 50% off cost by the amount of books subtract three, which are the original three books (The first two are full price, the third is free)
     *Therefore, every subsequent book after the first three is given a discount of 50% off the bookCost.
     */
  }
  const discountedTotal = preTotal - bookDiscount

  //Output
  document.getElementById("book-cost").innerHTML =
    "Before discount cost of the books is: " + preTotal.toFixed(2)
  document.getElementById("discounted-money").innerHTML =
    "Money discounted is: " + bookDiscount.toFixed(2)
  document.getElementById("discounted-cost").innerHTML =
    "Final Total cost of the books is: " + discountedTotal.toFixed(2)
}
