def is_isogram(string: str) -> bool:
    """Checks if a given string is an Isogram (ie. no repeated chars)
    Parameters:
        string the given string
    Retunrs:
    ========
        True, if string is an isogram
        False, otherwise
    """

    # Controls whether a char has already appeared or not
    char_has_appeared = set()

    # Cleans the input
    cleaned = string
    cleaned = cleaned.replace(" ", "")
    cleaned = cleaned.replace("-", "")
    cleaned = cleaned.lower()

    for char in cleaned:
        # We have found a repeated lchar
        if char in char_has_appeared:
            return False

        # We mark the char as already visited
        char_has_appeared.add(char)

    return True
