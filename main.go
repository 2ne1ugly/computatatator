/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: mchi <mchi@student.42.fr>                  +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/04/18 23:23:28 by mchi              #+#    #+#             */
/*   Updated: 2019/04/18 23:23:28 by mchi             ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
)

func printEquation(lhs []Term) {
	for i, term := range lhs {
		if i != 0 {
			if term.coef < 0 {
				fmt.Printf(" - %g", term.coef*-1)
			} else {
				fmt.Printf(" + %g", term.coef)
			}
		} else {
			fmt.Printf("%g", term.coef)
		}
		for name, power := range term.vars {
			if power == 1 {
				fmt.Printf(" * %s", name)
			} else {
				fmt.Printf(" * %s^%d", name, power)
			}
		}
	}
	fmt.Printf(" = 0\n")
}

func main() {
	if len(os.Args) != 2 {
		println("need one equation!")
	}
	terms := SplitByTerm(os.Args[1])
	terms = ParseTerms(terms)
	terms = SimplifyTerm(terms)
	fmt.Printf("Reduced form: ")
	printEquation(terms)
	degree := CountDegree(terms)
	fmt.Printf("Polynomial degree: %d\n", degree)
	switch degree {
	case 0:
		DetermineNonEquation(terms)
	case 1:
		SolveLinear(terms)
	case 2:
		SolveQuadratic(terms)
	default:
		fmt.Printf("The polynomial degree is stricly greater than 2, I can't solve.\n")
	}
}
