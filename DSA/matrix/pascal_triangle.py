def generate(v):
    res = []
    for i in range(v):
        res.append([1]*(i+1))
        if i > 1:
            for j in range(1, i):
                res[i][j] = res[i-1][j-1]+res[i-1][j]

    return res


v = 5
print(generate(v))
