def two_fer(name: str) -> str:
    """Gets a name and returns the following string:

    One for X, one for me.

    Where X is the name passed as an argument. If name is empty, X is 'me'
    """

    if name is None or name == "":
        return "One for you, one for me."

    return f"One for {name}, one for me."
