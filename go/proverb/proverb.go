// Package proverb has the function Proverb which solves the given problem
package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var proverb []string

	// Edge case
	if len(rhyme) == 0 {
		return proverb
	}

	// Body of the rhyme
	for i := 0; i < len(rhyme)-1; i++ {
		current_sentence := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		proverb = append(proverb, current_sentence)
	}

	// End of the rhyme
	last_sencente := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	proverb = append(proverb, last_sencente)

	return proverb
}
