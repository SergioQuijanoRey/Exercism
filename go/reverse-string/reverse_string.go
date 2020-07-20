// Package reverse has the function Reverse which reverses a given string
package reverse

// normalReverse reverses a string using a basic reverse for loop
func normalReverse(word string) string{
    new_word := make([]rune, len(word))
    for index, letter := range word{
        new_word[len(word) - index - 1] = letter
    }

    return string(new_word)
}

func tunedReverse(word string) string{
    new_word := ""
    for index := range word{
        new_word += string(word[len(word) - index - 1])
    }

    return string(new_word)
}

// Reverse reverses a given string
func Reverse(word string) string{
    return tunedReverse(word)
}
