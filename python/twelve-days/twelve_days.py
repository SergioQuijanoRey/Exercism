def recite(start_verse: int, end_verse: int) -> str[]:
    """Recites some verses of the lyrics of 'The Twelve Days of Christmas'

    Parameters:
    ===========
    start_verse: the starting verse
    end_verse: the end verse

    Returns:
    ========
    verses: the expected verses

    Raises:
    =======
    Exception("Bad verse indexing"): if verses index are given badly
    Exception("Starting verse out of index"): if starting verse is out of index

    """
    if start_verse > end_verse:
        raise Exception("Bad verse indexing")

    if start_verse < 1 or start_verse > 12:
        raise Exception("Starting verse out of index")

    # Auxiliar data
    numbers = {
        1: "first",
        2: "second",
        3: "third",
        4: "fourth",
        5: "fifth",
        6: "sixth",
        7: "seventh",
        8: "eighth",
        9: "ninth",
        10: "tenth",
        11: "eleventh",
        12: "twelfth"
    }
    verses = ""


    return numbers[12]
