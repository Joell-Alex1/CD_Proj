package parser

import (
	"fmt"
)

// Grammar and LL(1) Table
type Grammar struct {
	NonTerminals []string
	Terminals    []string
	StartSymbol  string
	Productions  map[string][][]string
}

type LL1Table map[string]map[string][]string

// Compute FIRST sets
func ComputeFirst(g Grammar) map[string]map[string]bool {
	first := make(map[string]map[string]bool)

	// Initialize
	for _, nt := range g.NonTerminals {
		first[nt] = make(map[string]bool)
	}

	changed := true
	for changed {
		changed = false
		for nt, prods := range g.Productions {
			for _, prod := range prods {
				if len(prod) == 0 { // epsilon
					if !first[nt]["ε"] {
						first[nt]["ε"] = true
						changed = true
					}
					continue
				}
				symbol := prod[0]
				if contains(g.Terminals, symbol) {
					if !first[nt][symbol] {
						first[nt][symbol] = true
						changed = true
					}
				} else { // non-terminal
					for k := range first[symbol] {
						if k != "ε" && !first[nt][k] {
							first[nt][k] = true
							changed = true
						}
					}
				}
			}
		}
	}
	return first
}

// Compute FOLLOW sets
func ComputeFollow(g Grammar, first map[string]map[string]bool) map[string]map[string]bool {
	follow := make(map[string]map[string]bool)
	for _, nt := range g.NonTerminals {
		follow[nt] = make(map[string]bool)
	}
	follow[g.StartSymbol]["$"] = true

	changed := true
	for changed {
		changed = false
		for nt, prods := range g.Productions {
			for _, prod := range prods {
				for i, sym := range prod {
					if contains(g.NonTerminals, sym) {
						next := prod[i+1:]
						if len(next) == 0 {
							for k := range follow[nt] {
								if !follow[sym][k] {
									follow[sym][k] = true
									changed = true
								}
							}
						} else {
							firstNext := computeFirstString(next, first, g.Terminals)
							for k := range firstNext {
								if k != "ε" && !follow[sym][k] {
									follow[sym][k] = true
									changed = true
								}
							}
							if firstNext["ε"] {
								for k := range follow[nt] {
									if !follow[sym][k] {
										follow[sym][k] = true
										changed = true
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return follow
}

func computeFirstString(symbols []string, first map[string]map[string]bool, terminals []string) map[string]bool {
	res := make(map[string]bool)
	for _, s := range symbols {
		if contains(terminals, s) {
			res[s] = true
			return res
		} else {
			for k := range first[s] {
				res[k] = true
			}
			if !first[s]["ε"] {
				return res
			}
		}
	}
	res["ε"] = true
	return res
}

// Build LL(1) parsing table
func BuildParsingTable(g Grammar, first, follow map[string]map[string]bool) LL1Table {
	table := make(LL1Table)
	for _, nt := range g.NonTerminals {
		table[nt] = make(map[string][]string)
		for _, prod := range g.Productions[nt] {
			firstProd := computeFirstString(prod, first, g.Terminals)
			for t := range firstProd {
				if t != "ε" {
					table[nt][t] = prod
				} else {
					for f := range follow[nt] {
						table[nt][f] = prod
					}
				}
			}
		}
	}
	return table
}

// Helper to check if slice contains string
func contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

// Push symbols onto stack (in reverse)
func push(stack []string, symbols []string) []string {
	for i := len(symbols) - 1; i >= 0; i-- {
		if symbols[i] != "ε" {
			stack = append(stack, symbols[i])
		}
	}
	return stack
}

// LL(1) parse
func LL1ParseTable(input []string, g Grammar, table LL1Table) {
	stack := []string{"$", g.StartSymbol}
	input = append(input, "$")

	i := 0
	fmt.Printf("\n%-20s %-20s %-20s\n", "Stack", "Input", "Action")
	fmt.Println("------------------------------------------------------------")

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		curr := input[i]

		fmt.Printf("%-20v %-20v ", stack, input[i:])

		if top == curr {
			fmt.Printf("Match %s\n", top)
			stack = stack[:len(stack)-1]
			i++
		} else if prod, ok := table[top][curr]; ok {
			fmt.Printf("%s -> %v\n", top, prod)
			stack = stack[:len(stack)-1]
			stack = push(stack, prod)
		} else {
			fmt.Printf("Error: unexpected %s\n", curr)
			return
		}
	}
	fmt.Println("Parsing successful ✅")
}
