/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   util.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: mchi <mchi@student.42.fr>                  +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/04/19 00:13:53 by mchi              #+#    #+#             */
/*   Updated: 2019/04/19 00:13:53 by mchi             ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"log"
	"reflect"
	"regexp"
	"strings"
)

//Term : elements of equation
type Term struct {
	token string
	coef  float32
	vars  map[string]int
}

//Pow : simple pow.
func Pow(base float32, power int) float32 {
	if power == 1 {
		return base
	} else if power == -1 {
		return 1 / base
	}
	if power%2 == 0 {
		value := Pow(base, power/2)
		return value * value
	}
	return Pow(base, power-1) * base
}

//Sqrt : simple sqrt.
func Sqrt(value float32) float32 {
	if value < 0 {
		log.Fatalln("no negs on sqrt!!")
	}
	var x float32
	x = 1
	for i := 0; i < 10; i++ {
		x = (x + value/x) / 2
	}
	return x
}

//SplitByTerm : splits by + = - and moves all to lhs.
func SplitByTerm(str string) []Term {
	eqTerms := regexp.MustCompile("[\\+\\-\\=]|[^\\=\\+\\-]*")
	str = strings.ReplaceAll(str, " ", "")
	tokens := eqTerms.FindAllString(str, -1)
	isLHS := true
	lhs := make([]Term, len(tokens)/2+1)
	for i := range lhs {
		lhs[i].coef = 1
	}
	i := 0
	expVal := true
	for _, token := range tokens {
		if !expVal {
			if isLHS {
				switch token {
				case "=":
					isLHS = false
					fallthrough
				case "-":
					lhs[i+1].coef = -1
					fallthrough
				case "+":
					i++
				default:
					log.Fatalln("unexpected value", token)
				}
			} else {
				switch token {
				case "=":
					log.Fatalln("multiple =")
				case "+":
					lhs[i+1].coef = -1
					fallthrough
				case "-":
					i++
				default:
					log.Fatalln("unexpected value", token)
				}
			}
		} else {
			lhs[i].token = token
		}
		expVal = !expVal
	}
	if isLHS {
		log.Fatalln("not an equation")
	}
	return lhs
}

//SimplifyTerm : combines and cleans terms
func SimplifyTerm(terms []Term) []Term {
	var result []Term
	for _, term := range terms {
		exist := false
		for i := range result {
			if reflect.DeepEqual(result[i].vars, term.vars) {
				exist = true
				result[i].coef += term.coef
			}
		}
		if !exist {
			result = append(result, term)
		}
	}
	return deleteToNatural(result)
}

//CountDegree : counts degrees
func CountDegree(terms []Term) int {
	var highest int
	highest = 0
	for _, term := range terms {
		for _, power := range term.vars {
			if power > highest {
				highest = power
			}
		}
	}
	return highest
}
