package main

import (
	"fmt"
)

type chain struct {
	length int
	parts  []string
}

func preprocess(tokens []string) func(string) bool {
	dict := map[string]bool{}
	for _, token := range tokens {
		dict[token] = true
	}

	return func(input string) bool {
		if input == "" {
			return true
		}
		chains := []chain{{0, nil}}
		length := len(input)
		for i := 1; i <= length; i++ {
			for j := len(chains) - 1; j >= 0; j-- {
				token := input[chains[j].length:i]
				if dict[token] {
					b := append(chains[j].parts, token)
					if i == length {
						return true
					}
					chains = append(chains, chain{i, b})
					break
				}
			}
		}
		return false
	}
}

func main() {
	validator := preprocess([]string{"a", "abcccccc", "abc", "cc", "d", "eee", "ttt"})
	for _, input := range []string{"j", "abcjttt", "abceeeabcccccttt", "tttabcaaaadeee", "tttabcaaaadeeeabccccccccccccccc", "tttabcaaaadeeeabcccccccccccccce", "", "abcccccccccc", "abccccccccc", "tttabcttt"} {
		fmt.Printf("%s: %v\n", input, validator(input))
	}
}
