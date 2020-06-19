// Package scrabble has the function Score that solves the given proglem
package scrabble

import "strings"

var scores map[rune]int

// init initializes the scores map data structure
func init(){
    scores = map[rune]int{
        'a': 1,
        'e': 1,
        'i': 1,
        'o': 1,
        'u': 1,
        'l': 1,
        'n': 1,
        'r': 1,
        's': 1,
        't': 1,
        'd': 2,
        'g': 2,
        'b': 3,
        'c': 3,
        'm': 3,
        'p': 3,
        'f': 4,
        'h': 4,
        'v': 4,
        'w': 4,
        'y': 4,
        'k': 5,
        'j': 8,
        'x': 8,
        'q': 10,
        'z': 10,
    }
}

// Score calculates the score of a given word
func Score(input string) int{
    // Lower letters has the same score as upper letters
    input = strings.ToLower(input)

    score := 0
    for _, letter := range input{
        value, found := scores[letter]
        if found == true{
            score += value
        }
    }

    return score
}
