class Matrix:
    """Class that represents a matrix, given a string representation of it"""

    def __init__(self, matrix_string):
        """Class initializer"""
        self.matrix = self.__string_to_matrix(matrix_string)

    def row(self, index):
        """Given an index, starting from 1, returns the index-th row of the matrix"""
        return self.matrix[index - 1]

    def column(self, index):
        """Given an index, starting from 1, returns the index-th column of the matrix
        reading column top-to-bottom while moving from left-to-right"""
        col = []
        for col_indx in range(len(self.matrix)):
            col.append(self.matrix[col_indx][index - 1])
        return col

    def __string_to_matrix(self, matrix_string):
        """Given a string representation of the matrix, returns the matrix representation
        using lists"""
        # We get a list representation of the matrix, but numbers are still encoded
        # as strings
        matrix = []
        matrix = matrix_string.split("\n")
        new_matrix = [row.split(" ") for row in matrix]
        matrix = new_matrix

        # We convert the string entries to numbers
        for row in range(len(matrix)):
            for col in range(len(matrix[row])):
                matrix[row][col] = int(matrix[row][col])

        return matrix
