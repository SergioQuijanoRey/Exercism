use unicode_segmentation::UnicodeSegmentation;

/// Given a string, reverses the string
pub fn reverse(input: &str) -> String {
    input.graphemes(true).rev().collect()
}
