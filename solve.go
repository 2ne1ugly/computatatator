/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   solve.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: mchi <mchi@student.42.fr>                  +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/04/19 00:43:15 by mchi              #+#    #+#             */
/*   Updated: 2019/04/19 00:43:15 by mchi             ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

//SolveLinear : solves degree 1.
func SolveLinear(terms []Term) {
	var a float32
	var b float32
	for _, term := range terms {
		_, ok := term.vars["X"]
		if ok {
			b = term.coef
		} else {
			a = term.coef
		}
	}
	fmt.Printf("The solution is:\n%g\n", -a/b)
}

//SolveQuadratic : solves degree 2.
func SolveQuadratic(terms []Term) {
	var a float32
	var b float32
	var c float32
	for _, term := range terms {
		i, ok := term.vars["X"]
		if ok && i == 2 {
			a = term.coef
		} else if ok && i == 1 {
			b = term.coef
		} else {
			c = term.coef
		}
	}
	disc := Pow(b, 2) - 4*a*c
	fmt.Printf("This is a Quadratic Equation.\n")
	if disc < 0 {
		fmt.Printf("Disc is negative. Two imaginary solution exists:\n")
		fmt.Printf("%g + i * %g\n%g - i * %g\n", -b/(2*a), Sqrt(-disc)/(2*a), -b/(2*a), Sqrt(-disc)/(2*a))
	} else if disc == 0 {
		fmt.Printf("Disc is 0. Only single solution exists:\n%g\n", -b/2*a)
	} else {
		fmt.Printf("Disc is positive. Two solution exists:\n%g\n%g\n", -b/(2*a)+Sqrt(disc)/(2*a), -b/(2*a)-Sqrt(disc)/(2*a))
	}
}

//DetermineNonEquation : checks if this is always true
func DetermineNonEquation(terms []Term) {
	var a float32
	for _, term := range terms {
		a = term.coef
	}
	if a == 0 {
		fmt.Printf("The solution is:\nALL REAL NUMBERS\n")
	} else {
		fmt.Printf("The solution is:\nNONE\n")
	}
}
