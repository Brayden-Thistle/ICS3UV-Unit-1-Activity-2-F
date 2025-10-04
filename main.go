/*
 * @author Brayden Thistle
 * @version 1.0.0
 * @date 2025-10-04
 * @fileoverview this program displays my year of birth and the year im expected to graduate from highschool
 */

package main

import "fmt"

func main () {
	// indicators
	fmt.Println("Year Born     Year Graduated     Result")
	
	// addition
	fmt.Println("2009       +     2027           =   " + fmt.Sprint(2009 + 2027))
	
	// subtraction
	fmt.Println("2009       -     2027           =   " + fmt.Sprint(2009 - 2027))

	// multiplication
	fmt.Println("2009       *     2027           =   " + fmt.Sprint(2009 * 2027))

	// division
	fmt.Println("2009       /     2027           =   " + fmt.Sprint(2009 / 2027))

	// Modulo
	fmt.Println("2009       %     2027           =   " + fmt.Sprint(2009 % 2027))

	fmt.Println("\nDone")
}