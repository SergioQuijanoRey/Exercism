def convert(number: int) -> str:
    """Converts a number to a String using the following rules:

        - If number has 3 as a factor, add 'Pling' to the result
        - If number has 5 as a factor, add 'Plang' to the result
        - If number has 7 as a factor, add 'Plong' to the result
        - If number does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number
    """

    result = ""

    if number % 3 == 0:
        result += "Pling"
    if number % 5 == 0:
        result += "Plang"
    if number % 7 == 0:
        result += "Plong"

    # The given number has not the given factors
    if result == "":
        return str(number)

    return result
