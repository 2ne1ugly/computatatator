/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: mchi <mchi@student.42.fr>                  +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/04/18 23:23:28 by mchi              #+#    #+#             */
/*   Updated: 2019/08/07 22:13:09 by mchi             ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"bufio"
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
	var terms []Term
	if len(os.Args) == 1 {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		line = line[:len(line)-1]
		if err != nil {
			panic("error reading line")
		}
		terms = SplitByTerm(line)
	} else {
		terms = SplitByTerm(os.Args[1])
	}
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
