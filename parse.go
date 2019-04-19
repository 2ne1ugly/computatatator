/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   parse.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: mchi <mchi@student.42.fr>                  +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/04/17 15:45:55 by mchi              #+#    #+#             */
/*   Updated: 2019/04/17 15:45:55 by mchi             ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

//ParsePow : parses values with power
func ParsePow(str string) float32 {
	values := strings.Split(str, "^")
	if len(values) > 2 {
		log.Fatalln("cannot handle nested power")
	}
	base, err := strconv.ParseFloat(values[0], 32)
	if err != nil {
		log.Fatalln("not a float", values[0])
	}
	power, err := strconv.ParseInt(values[1], 10, 32)
	if err != nil {
		log.Fatalln("not a int", values[0])
	}
	return Pow(float32(base), int(power))
}

//ParseVar : parses variable with power
func ParseVar(str string) (string, int) {
	values := strings.Split(str, "^")
	if len(values) > 2 {
		log.Fatalln("cannot handle nested power")
	}
	power, err := strconv.ParseInt(values[1], 10, 32)
	if err != nil {
		log.Fatalln("not a int", values[0])
	}
	return values[0], int(power)
}

//ParseFloat : parses Float
func ParseFloat(str string) float32 {
	value, err := strconv.ParseFloat(str, 32)
	if err != nil {
		log.Fatalln("not a float", str)
	}
	return float32(value)
}

func deleteToNatural(terms []Term) []Term {
	var newTerms []Term
	for _, term := range terms {
		if term.coef != 0 {
			for name, power := range term.vars {
				if power == 0 {
					delete(term.vars, name)
				}
			}
			newTerms = append(newTerms, term)
		}
	}
	return newTerms
}

//ParseTerms : reads sub-tokens of term and translate to variables and coef.
func ParseTerms(terms []Term) []Term {
	for i := range terms {
		opSplit := regexp.MustCompile("[\\*\\/]|[^\\*\\/]*")
		subTokens := opSplit.FindAllString(terms[i].token, -1)
		expVal := true
		isDiv := false
		terms[i].vars = make(map[string]int, 1)
		for _, token := range subTokens {
			if !expVal {
				switch token {
				case "/":
					isDiv = true
				case "*":
					isDiv = false
				default:
					log.Fatalf("unexpected value %s from term %s", token, terms[i].token)
				}
			} else {
				if token[0] == 'X' {
					if strings.Contains(token, "^") {
						variable, power := ParseVar(token)
						if isDiv {
							terms[i].vars[variable] -= power
						} else {
							terms[i].vars[variable] += power
						}
					} else {
						if isDiv {
							terms[i].vars["X"]--
						} else {
							terms[i].vars["X"]++
						}
					}
				} else {
					var value float32
					if strings.Contains(token, "^") {
						value = ParsePow(token)
					} else {
						value = ParseFloat(token)
					}
					if isDiv {
						terms[i].coef /= value
					} else {
						terms[i].coef *= value
					}
				}
			}
			expVal = !expVal
		}
	}
	return terms
}
