def two_fer(name: str = "you") -> str:
    """Gets a name and returns the following string:

    One for X, one for me.

    Where X is the name passed as an argument. If name is empty, X is 'you'
    """

    return f"One for {name}, one for me."
