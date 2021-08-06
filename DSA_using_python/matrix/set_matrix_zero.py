def set_matrix(matrix):
    r = len(matrix)
    c = len(matrix[0])
    row = set()
    col = set()

    for i in range(r):
        for j in range(c):
            if matrix[i][j] == 0:
                row.add(i)
                col.add(j)
    print(row)
    print(col)
    for i in range(r):
        for j in range(c):
            if i in row or j in col:
                matrix[i][j] = 0

    return matrix


arr = [[0, 1, 2, 0], [3, 4, 5, 2], [1, 3, 1, 5]]
print(set_matrix(arr))
