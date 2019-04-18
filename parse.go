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
	"os"
	"regexp"
	"strconv"
)

//saves by terms
type variable struct {
	name string
	exp  int
}

type term struct {
	tokens []string
	coef   float64
	vars   []variable
}

func splitByTerm(str string) []term {
	eqTerms := regexp.MustCompile("\\*|\\+|\\-|=|X\\^[0-9]*|[0-9]*\\.?[0-9]*|[0-9]*|X")
	tokens := eqTerms.FindAllString(str, -1)
	isLHS := true
	var lhs []term
	lhs = append(lhs, term{})
	lhs[len(lhs)-1].coef = 1
	for _, token := range tokens {
		if isLHS {
			switch token {
			case "+":
				lhs = append(lhs, term{})
				lhs[len(lhs)-1].coef = 1
			case "-":
				lhs = append(lhs, term{})
				lhs[len(lhs)-1].coef = -1
			case "=":
				lhs = append(lhs, term{})
				lhs[len(lhs)-1].coef = -1
				isLHS = false
			default:
				lhs[len(lhs)-1].tokens = append(lhs[len(lhs)-1].tokens, token)
			}
		} else {
			switch token {
			case "+":
				lhs = append(lhs, term{})
				lhs[len(lhs)-1].coef = -1
			case "-":
				lhs = append(lhs, term{})
				lhs[len(lhs)-1].coef = 1
			case "=":
				log.Fatalln("multiple =")
			default:
				lhs[len(lhs)-1].tokens = append(lhs[len(lhs)-1].tokens, token)
			}
		}
		println(token)
	}
	if isLHS {
		log.Fatalln("not an equation")
	}
	return lhs
}

var isVar = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func evalTerm(term *term) {
	if len(term.tokens) == 0 {
		log.Fatalln("empty term")
	}
	expVal := true
	isDiv := false
	isPow := false
	var prevOp byte
	var err error
	for _, v := range term.tokens {
		if expVal {
			if isVar(v) {
				term.vars = append(term.vars, variable{})
				term.vars[len(term.vars)-1].name = v
				if isDiv {
					term.vars[len(term.vars)-1].exp = -1
				} else {
					term.vars[len(term.vars)-1].exp = 1
				}
			} else {
				if isPow {

				}
				stash, err = strconv.ParseFloat(v, 64)
				if err != nil {
					log.Fatalln("failed parsing number:", v)
				}
			}
		} else {
			switch v {
			case "*":
				isDiv = false
			case "/":
				isDiv = true
			}
			expVal = true
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		println("need one equation!")
	}
	splitByTerm(os.Args[1])
}
