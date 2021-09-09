use std::collections::HashMap;
use std::collections::HashSet;
use unicode_segmentation::UnicodeSegmentation;

/// Given a word and a list of possible anagrams of the word, returns the list
/// of the words which are actually anagrams
pub fn anagrams_for<'a>(word: &str, possible_anagrams: &'a [&str]) -> HashSet<&'a str> {
    let mut results: HashSet<&'a str> = HashSet::new();
    for candidate in possible_anagrams {
        if is_anagram(word, candidate) {
            results.insert(candidate);
        }
    }

    return results;
}

/// Checks if two words are anagrams of each other
fn is_anagram(word: &str, candidate: &str) -> bool {
    // Anagram condition is case insensitive
    let word = word.to_lowercase();
    let candidate = candidate.to_lowercase();

    // We do not consider same word as an anagram
    if word == candidate {
        return false;
    }

    // Anagrams need to have same lenght
    if word.len() != candidate.len() {
        return false;
    }

    // Get each grapheme and count the number of repetitions of the grapheme
    let word_count = count_graphemes_on_word(&word);
    let candidate_count = count_graphemes_on_word(&candidate);

    return word_count == candidate_count;
}

/// Returns a hash map with the letters of the word and it's repetitions
fn count_graphemes_on_word(word: &str) -> HashMap<&str, i32> {
    let mut word_count = HashMap::new();
    for grapheme in word.graphemes(true) {
        if word_count.contains_key(grapheme) == false {
            word_count.insert(grapheme, 0);
        }

        let new_count = word_count.get(grapheme).unwrap() + 1;
        word_count.insert(grapheme, new_count);
    }

    return word_count;
}
